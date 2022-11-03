package main

import (
	"context"

	app "github.com/Meystergod/online-store-api/internal/application"
	"github.com/Meystergod/online-store-api/internal/config"
	"github.com/Meystergod/online-store-api/pkg/logging"
)

func main() {
	cfg := config.GetConfig()

	logging.Init(cfg.AppConfig.LogLevel)
	logger := logging.GetLogger()
	logger.Info("logger and config initialized")

	application, err := app.NewApp(cfg, logger)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("running api server")
	if application.Run(context.Background()) != nil {
		logger.Fatal(err)
	}
}
