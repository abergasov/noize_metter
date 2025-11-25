package testhelpers

import (
	"context"
	"noize_metter/internal/config"
	"noize_metter/internal/logger"
	"noize_metter/internal/repository"
	"noize_metter/internal/service/ces"
	"noize_metter/internal/service/noise_metter"
	"noize_metter/internal/service/notificator"
	"noize_metter/internal/service/substation"
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

	ServiceNoise            *noise_metter.Service
	ServiceSubstation       *substation.Service
	ServiceSlackNotificator *notificator.SlackService
	ServiceCes              *ces.Service
}

func GetClean(t *testing.T) *TestContainer {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	conf := getTestConfig(t)

	t.Cleanup(cancel)

	appLog := logger.NewAppSLogger()
	// repo init
	repo := repository.InitRepo(ctx, appLog, conf)

	// service init
	serviceNoise := noise_metter.NewService(ctx, appLog, conf, repo)
	serviceSubstation, err := substation.NewService(ctx, appLog, conf, repo)
	require.NoError(t, err)
	serviceCes := ces.NewService(ctx, appLog, conf, repo)
	return &TestContainer{
		Ctx:                     ctx,
		Cfg:                     conf,
		Logger:                  appLog,
		Repo:                    repo,
		ServiceNoise:            serviceNoise,
		ServiceSubstation:       serviceSubstation,
		ServiceSlackNotificator: notificator.NewSlackService(appLog, conf),
		ServiceCes:              serviceCes,
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
