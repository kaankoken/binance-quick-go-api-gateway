package main

import (
	"context"

	"go.uber.org/fx"

	"github.com/kaankoken/binance-quick-go-api-gateway/config"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/helper"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/telegram"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		helper.Module,
		config.Module,
		telegram.EngineModule,
		telegram.RouteModule,
		fx.Invoke(registerHooks),
	)

	app.Run()
}

func registerHooks(lifecycle fx.Lifecycle, h *telegram.Handler, config *config.Config, logger *helper.Handler) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				logger.Info("Starting application in " + config.Port)
				go h.Gin.Run(config.Port)
				return nil
			},
			OnStop: func(context.Context) error {
				logger.Info("Stopping application")
				return nil
			},
		},
	)
}
