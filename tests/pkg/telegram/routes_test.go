package telegram_test

import (
	"os"
	"testing"

	"github.com/kaankoken/binance-quick-go-api-gateway/config"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/helper"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/telegram"
	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	// Needed Module initialization
	// debug logger
	d := helper.SetLoggerFormat()
	// release logger
	logger := helper.InitializeLogger()
	r := helper.InitializeLoggerPtr(logger)

	generateFakeConfig(t)
	t.Parallel()

	t.Run("routes-test=registeration-success", func(t *testing.T) {
		c, err := config.LoadConfig()

		// log handler initialization
		l := helper.Initialize(c, d, r)
		handler, err2 := pkg.NewHandler(c)

		service := telegram.InitServiceClient(c, l)
		res := telegram.Initialize(service, l)

		telegram.RegisterRoutes(res, handler)

		assert.Nil(t, err)
		assert.Nil(t, err2)
	})
}

func generateFakeConfig(tb testing.TB) {
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
