package cmd_test

import (
	"context"
	"os"
	"testing"

	"github.com/kaankoken/binance-quick-go-api-gateway/cmd"
	"github.com/kaankoken/binance-quick-go-api-gateway/config"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/helper"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/telegram"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func TestMain(t *testing.T) {
	generateFakeSuccessfulConfig(t)

	t.Run("test-main-app-func", func(t *testing.T) {
		app := fxtest.New(
			t,
			config.Module,
			telegram.ClientModule,
			pkg.EngineModule,
			telegram.RouteModule,
			helper.LoggerModule,
			fx.Invoke(cmd.RegisterHooks),
		)

		app.RequireStart()
		defer app.RequireStop()
	})

	t.Run("test-main-app", func(t *testing.T) {
		app := cmd.MainApp()

		app.Start(context.Background())

		defer app.Stop(context.Background())
	})
}

func generateFakeSuccessfulConfig(tb testing.TB) {
	os.Remove("config.env")

	fakeConfig := `
	FLAVOR=debug
	GIN_MODE=test
	PORT=:6542
	AUTH_SVC_URL=localhost:1
	OBSERVER_SVC_URL=localhost:2
	TELEGRAM_SVC_URL=localhost:3
	`

	err := os.WriteFile("config.env", []byte(fakeConfig), 0644)
	assert.Nil(tb, err)
}
