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
	if err := os.MkdirAll(conf.StorageNoiseFolder, 0o755); err != nil {
		log.Fatal("error creating noise storage folder", err)
	}
	if err := os.MkdirAll(conf.StorageSubstationFolder, 0o755); err != nil {
		log.Fatal("error creating substation storage folder", err)
	}
	return &Repo{
		ctx:  ctx,
		conf: conf,
		log:  log,
	}
}
