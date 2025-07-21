package deployer

import (
	"context"
	"noize_metter/internal/config"
	"noize_metter/internal/logger"
	"noize_metter/internal/repository"
	"noize_metter/internal/service/notificator"
	"os"
)

type Service struct {
	ctx    context.Context
	log    logger.AppLogger
	cfg    *config.AppConfig
	repo   *repository.Repo
	notify notificator.Notificator

	rst chan os.Signal
}

func NewService(ctx context.Context, cfg *config.AppConfig, log logger.AppLogger, repo *repository.Repo, notify notificator.Notificator, rst chan os.Signal) *Service {
	return &Service{
		ctx:    ctx,
		cfg:    cfg,
		log:    log.With(logger.WithService("deployer")),
		repo:   repo,
		notify: notify,
		rst:    rst,
	}
}

func (s *Service) Run() {
	s.log.Info("deployer service started")
	go s.fetchLatestVersionLoop()
}
