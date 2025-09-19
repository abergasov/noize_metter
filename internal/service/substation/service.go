package substation

import (
	"context"
	"fmt"
	"math"
	"noize_metter/internal/config"
	"noize_metter/internal/entities"
	"noize_metter/internal/logger"
	"noize_metter/internal/repository"
	"noize_metter/internal/utils"
	"time"

	"github.com/simonvetter/modbus"
)

type Service struct {
	ctx      context.Context
	log      logger.AppLogger
	conf     *config.AppConfig
	repo     *repository.Repo
	mbClient *modbus.ModbusClient

	items *utils.RWSlice[entities.ModbusRegisters]
}

func NewService(
	ctx context.Context,
	log logger.AppLogger,
	conf *config.AppConfig,
	repo *repository.Repo,
) (*Service, error) {
	client, err := modbus.NewClient(&modbus.ClientConfiguration{
		URL:     conf.CFModbusHost,
		Timeout: 10 * time.Second,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create modbus client: %w", err)
	}
	srv := &Service{
		ctx:      ctx,
		log:      log.With(logger.WithService("cf_modbus")),
		conf:     conf,
		mbClient: client,
		repo:     repo,
		items:    utils.NewRWSlice[entities.ModbusRegisters](),
	}
	go utils.BGPruneOldFiles(ctx, srv.log, srv.conf.StorageSubstationFolder)
	go srv.bgDumpData()
	return srv, nil
}

func (s *Service) Stop() {}

func ParseFloat32(valueA, valueB uint16) float32 {
	uint32Value := uint32(valueA)<<16 | uint32(valueB)
	return math.Float32frombits(uint32Value)
}

func ParseFloat32V(values []uint16, offset int) float32 {
	return ParseFloat32(values[offset], values[offset+1])
}
