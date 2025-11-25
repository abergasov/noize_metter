package ces

import (
	"context"
	"fmt"
	"noize_metter/internal/config"
	"noize_metter/internal/entities"
	"noize_metter/internal/logger"
	"noize_metter/internal/repository"
	"noize_metter/internal/service"
	"sync"
	"time"
)

const (
	ServiceName     = "ces"
	CollectInterval = 5 * time.Minute // how often to collect foreman stat
)

var (
	URLListIMDC1 = []string{
		"http://10.20.10.101",
		"http://10.20.10.107",
	}
	URLListIMDC2 = []string{
		"http://10.20.10.116",
		"http://10.20.10.117",
	}
)

type Service struct {
	ctx  context.Context
	log  logger.AppLogger
	conf *config.AppConfig
	repo *repository.Repo

	token   map[string]string
	tokenMU sync.RWMutex
}

// NewService creates new instance of Service
// Each IMDC has 4 separate Megabox feeders, Megabox 1, 2, 3, 4.
// Each Megabox feeder supplies power to 6 tanks.
// * Megabox 1 to tanks 1-6
// * Megabox 2 to tanks 7-12
// * Megabox 3 to tanks 13-18
// * Megabox 4 to tanks 19-24.
// Servers locations:
//   - IMDC-1
//     http://10.20.10.101/api/
//     http://10.20.10.107/api/
//   - IMDC-2
//     http://10.20.10.116/api/
//     http://10.20.10.117/api/
func NewService(
	ctx context.Context,
	log logger.AppLogger,
	conf *config.AppConfig,
	repoOldDwh *repository.Repo,
) *Service {
	srv := &Service{
		ctx:   ctx,
		log:   log,
		conf:  conf,
		token: make(map[string]string, len(URLListIMDC1)+len(URLListIMDC2)),
		repo:  repoOldDwh,
	}

	hostURL := fmt.Sprintf("%s/api-mapi/v1/private/ces_tank/upload_data", conf.DataHost)
	go service.BGUploadData[entities.CesTank](ctx, log, conf, hostURL, conf.StorageCESTanksFolder)
	go service.BGPruneOldFiles(ctx, srv.log, srv.conf.StorageCESTanksFolder)

	hostURL = fmt.Sprintf("%s/api-mapi/v1/private/ces_megaboxes/upload_data", conf.DataHost)
	go service.BGUploadData[entities.MegaBox](ctx, log, conf, hostURL, conf.StorageCESMegaBoxesFolder)
	go service.BGPruneOldFiles(ctx, srv.log, srv.conf.StorageCESMegaBoxesFolder)

	hostURL = fmt.Sprintf("%s/api-mapi/v1/private/ces_channels/upload_data", conf.DataHost)
	go service.BGUploadData[entities.CesTanksChannels](ctx, log, conf, hostURL, conf.StorageCESChannelsFolder)
	go service.BGPruneOldFiles(ctx, srv.log, srv.conf.StorageCESChannelsFolder)

	hostURL = fmt.Sprintf("%s/api-mapi/v1/private/ces_channels_v2/upload_data", conf.DataHost)
	go service.BGUploadData[entities.CesTanksChannelsV2](ctx, log, conf, hostURL, conf.StorageCESChannelsV2Folder)
	go service.BGPruneOldFiles(ctx, srv.log, srv.conf.StorageCESChannelsV2Folder)

	return srv
}

func (s *Service) Run() {
	s.log.Info("starting CES service...")
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-s.ctx.Done():
			s.log.Info("CES service stopped.")
			return
		case <-ticker.C:
			if err := s.RunIteration(); err != nil {
				s.log.Error("error in CES iteration", err)
			}
		}
	}
}
