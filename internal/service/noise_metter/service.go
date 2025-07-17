package noise_metter

import (
	"context"
	"noize_metter/internal/config"
	"noize_metter/internal/logger"
)

type Service struct {
	ctx  context.Context
	log  logger.AppLogger
	conf *config.AppConfig
}

func NewService(ctx context.Context, log logger.AppLogger, conf *config.AppConfig) *Service {
	return &Service{
		ctx:  ctx,
		log:  log.With(logger.WithService("noise_metter")),
		conf: conf,
	}
}
