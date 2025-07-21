package noise_metter

import (
	"context"
	"fmt"
	"noize_metter/internal/config"
	"noize_metter/internal/entities"
	"noize_metter/internal/logger"
	"noize_metter/internal/utils"
	"sync/atomic"
)

type Service struct {
	ctx  context.Context
	log  logger.AppLogger
	conf *config.AppConfig

	session atomic.Value
	items   *utils.RWSlice[entities.NoiseMeasures]
}

func NewService(ctx context.Context, log logger.AppLogger, conf *config.AppConfig) *Service {
	return &Service{
		ctx:     ctx,
		log:     log.With(logger.WithService("noise_metter")),
		conf:    conf,
		session: atomic.Value{},
		items:   utils.NewRWSlice[entities.NoiseMeasures](),
	}
}

func (s *Service) Run() error {
	s.log.Info("starting Noise Metter service...")
	s.session.Store("")
	if err := s.Auth(); err != nil {
		return fmt.Errorf("authentication failed: %w", err)
	}
	go s.bgSetSession()
	// Placeholder for actual service logic
	select {
	case <-s.ctx.Done():
		s.log.Info("Noise Metter service stopped.")
		return nil
	}
}
