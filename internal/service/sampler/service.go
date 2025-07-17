package sampler

import (
	"context"
	"noize_metter/internal/logger"
	"noize_metter/internal/repository"
)

type Service struct {
	ctx  context.Context
	log  logger.AppLogger
	repo *repository.Repo
}

func InitService(ctx context.Context, log logger.AppLogger, repo *repository.Repo) *Service {
	return &Service{
		ctx:  ctx,
		repo: repo,
		log:  log.With(logger.WithService("sampler")),
	}
}

func (s *Service) Stop() {
	s.log.Info("stopping service")
}
