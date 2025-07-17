package main

import (
	"context"
	"flag"
	"fmt"
	"noize_metter/internal/config"
	"noize_metter/internal/logger"
	"noize_metter/internal/repository"
	"noize_metter/internal/routes"
	samplerService "noize_metter/internal/service/sampler"

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
	repo := repository.InitRepo()

	appLog.Info("init services")
	service := samplerService.InitService(ctx, appLog, repo)

	appLog.Info("init http service")
	appHTTPServer := routes.InitAppRouter(appLog, service, fmt.Sprintf(":%d", appConf.AppPort))
	defer func() {
		if err = appHTTPServer.Stop(); err != nil {
			appLog.Fatal("unable to stop http service", err)
		}
	}()
	go func() {
		if err = appHTTPServer.Run(); err != nil {
			appLog.Fatal("unable to start http service", err)
		}
	}()

	// register app shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c // This blocks the main thread until an interrupt is received
	cancel()
}
