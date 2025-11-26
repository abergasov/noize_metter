package repository

import (
	"context"
	"noize_metter/internal/config"
	"noize_metter/internal/logger"
	"os"
)

type Repo struct {
	ctx  context.Context
	conf *config.AppConfig
	log  logger.AppLogger
}

func InitRepo(ctx context.Context, log logger.AppLogger, conf *config.AppConfig) *Repo {
	storageFolders := map[string]string{
		"noise storage folder":      conf.StorageNoiseFolder,
		"substation storage folder": conf.StorageSubstationFolder,
		"audio storage folder":      conf.StorageAudioFolder,

		// ces folders
		"ces channels v2": conf.StorageCESChannelsV2Folder,
		"ces channels":    conf.StorageCESChannelsFolder,
		"ces tanks":       conf.StorageCESTanksFolder,
		"ces megaboxes":   conf.StorageCESMegaBoxesFolder,
	}
	for key, folder := range storageFolders {
		if err := os.MkdirAll(folder, 0o755); err != nil {
			log.Fatal("error creating storage "+key+" folder", err)
		}
	}
	return &Repo{
		ctx:  ctx,
		conf: conf,
		log:  log,
	}
}
