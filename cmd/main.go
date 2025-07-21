package main

import (
	"context"
	"flag"
	"noize_metter/internal/config"
	"noize_metter/internal/logger"
	"noize_metter/internal/repository"
	"noize_metter/internal/service/deployer"
	"noize_metter/internal/service/noise_metter"
	"noize_metter/internal/service/notificator"

	"os"
	"os/signal"
	"syscall"
)

var (
	confFile = "configs/app_conf.yml"
)

func main() {
	flag.Parse()
	appLog := logger.NewAppSLogger()

	appLog.Info("app starting", logger.WithString("conf", confFile))
	appConf, err := config.InitConf(confFile)
	if err != nil {
		appLog.Fatal("unable to init config", err, logger.WithString("config", confFile))
	}
	ctx, cancel := context.WithCancel(context.Background())

	appLog.Info("create storage connections")

	appLog.Info("init repositories")
	repo := repository.InitRepo(ctx, appLog, appConf)

	appLog.Info("init services")
	notifier := notificator.NewSlackService(appLog, appConf)
	noise_metter.NewService(ctx, appLog, appConf, repo).Run()

	// register app shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	deployer.NewService(ctx, appConf, appLog, repo, notifier, c).Run()
	<-c // This blocks the main thread until an interrupt is received
	cancel()
}
