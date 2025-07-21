package testhelpers

import (
	"context"
	"noize_metter/internal/config"
	"noize_metter/internal/logger"
	"noize_metter/internal/repository"
	"noize_metter/internal/service/noise_metter"
	"os"
	"path"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type TestContainer struct {
	Ctx    context.Context
	Cfg    *config.AppConfig
	Logger logger.AppLogger

	Repo *repository.Repo

	ServiceNoise *noise_metter.Service
}

func GetClean(t *testing.T) *TestContainer {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	conf := getTestConfig(t)

	t.Cleanup(cancel)

	appLog := logger.NewAppSLogger()
	// repo init
	repo := repository.InitRepo(ctx, appLog, conf)

	// service init
	serviceNoise := noise_metter.NewService(ctx, appLog, conf)
	return &TestContainer{
		Ctx:          ctx,
		Cfg:          conf,
		Logger:       appLog,
		Repo:         repo,
		ServiceNoise: serviceNoise,
	}
}

func getTestConfig(t *testing.T) *config.AppConfig {
	dataPath, err := os.Getwd()
	require.NoError(t, err)
	data := strings.Split(dataPath, "noize_metter")
	appConf, err := config.InitConf(path.Join(data[0], "noize_metter", "configs", "app_conf.yml"))
	require.NoError(t, err)
	return appConf
}
