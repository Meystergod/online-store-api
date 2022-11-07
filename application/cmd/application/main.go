package main

import (
	"context"

	app "github.com/Meystergod/online-store-api/internal/application"
	"github.com/Meystergod/online-store-api/internal/config"
	"github.com/Meystergod/online-store-api/pkg/logging"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.GetConfig(ctx)

	logger := logging.GetLogger(ctx)
	ctx = logging.ContextWithLogger(ctx, logger)

	application, err := app.NewApp(ctx, cfg)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("running api server")
	if application.Run(ctx) != nil {
		logger.Fatal(err)
	}
}
