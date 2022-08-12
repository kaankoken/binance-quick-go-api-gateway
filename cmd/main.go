package main

import (
	"go.uber.org/fx"

	"github.com/kaankoken/binance-quick-go-api-gateway/config"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/helper"
)

func main() {
	app := fx.New(helper.Module, config.Module)

	app.Run()
}
