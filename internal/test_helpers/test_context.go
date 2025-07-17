package testhelpers

import (
	"context"
	"noize_metter/internal/config"
	"noize_metter/internal/logger"
	"noize_metter/internal/repository"
	samplerService "noize_metter/internal/service/sampler"
	"testing"
	"time"
)

type TestContainer struct {
	Ctx    context.Context
	Cfg    *config.AppConfig
	Logger logger.AppLogger

	Repo *repository.Repo

	ServiceSampler *samplerService.Service
}

func GetClean(t *testing.T) *TestContainer {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	conf := getTestConfig()

	t.Cleanup(cancel)

	appLog := logger.NewAppSLogger()
	// repo init
	repo := repository.InitRepo()

	// service init
	serviceSampler := samplerService.InitService(ctx, appLog, repo)
	return &TestContainer{
		Ctx:            ctx,
		Cfg:            conf,
		Logger:         appLog,
		Repo:           repo,
		ServiceSampler: serviceSampler,
	}
}

func getTestConfig() *config.AppConfig {
	return &config.AppConfig{
		AppPort: 0,
	}
}
