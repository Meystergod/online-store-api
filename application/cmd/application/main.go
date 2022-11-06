package main

import (
	"context"

	app "github.com/Meystergod/online-store-api/application/internal/application"
	"github.com/Meystergod/online-store-api/application/internal/config"
	"github.com/Meystergod/online-store-api/application/pkg/logging"
)

func main() {
	cfg := config.GetConfig()

	logger := logging.GetLogger(cfg.AppConfig.LogLevel)
	logger.Info("logger and config initialized")

	application, err := app.NewApp(cfg, &logger)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("running api server")
	if application.Run(context.Background()) != nil {
		logger.Fatal(err)
	}
}
