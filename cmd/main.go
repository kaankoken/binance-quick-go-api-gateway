package main

import (
	"go.uber.org/fx"

	"github.com/kaankoken/binance-quick-go-api-gateway/config"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/helper"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/telegram"
)

func main() {
	app := fx.New(
		helper.Module,
		config.Module,
		telegram.Module,
	)

	app.Run()
}
