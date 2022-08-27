package cmd

import (
	"context"

	"go.uber.org/fx"

	"github.com/kaankoken/binance-quick-go-api-gateway/config"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/helper"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/telegram"
)

// MainApp -> Registering main app to FX
func MainApp() *fx.App {
	app := fx.New(
		config.Module,
		telegram.ClientModule,
		pkg.EngineModule,
		telegram.RouteModule,
		helper.LoggerModule,
		fx.Invoke(RegisterHooks),
	)

	return app
}

// RegisterHooks -> Registering lifecycle of fx & running http server (Gin)
func RegisterHooks(lifecycle fx.Lifecycle, h *pkg.Handler, config *config.Config, logger *helper.LogHandler) {
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
