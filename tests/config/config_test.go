package config_test

import (
	"os"
	"testing"

	"github.com/kaankoken/binance-quick-go-api-gateway/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/fx/fxtest"
)

func TestConfig(t *testing.T) {
	os.Remove("config.env")

	t.Run("config=no-config-env-file", func(t *testing.T) {
		c, err := config.LoadConfig()

		assert.NotNil(t, err)
		assert.Nil(t, c)
	})

	t.Run("config=viper-unmarshall-error", func(t *testing.T) {
		generateFakeWronglConfig(t)

		c, err := config.LoadConfig()

		assert.NotNil(t, err)
		assert.Nil(t, c)
	})

	t.Run("config=unmarshall-empty", func(t *testing.T) {
		fakeConfig := "test=test11761723"

		err := os.WriteFile("config.env", []byte(fakeConfig), 0644)
		assert.Nil(t, err)

		c, err := config.LoadConfig()

		assert.Nil(t, err)
		assert.NotNil(t, c)

		assert.Empty(t, c.AuthSvcURL)
		assert.Empty(t, c.Flavor)
		assert.Empty(t, c.Mode)
		assert.Empty(t, c.ObserverSvcURL)
		assert.Empty(t, c.Port)
		assert.Empty(t, c.TelegramSvcURL)

		assert.Equal(t, c.AuthSvcURL, "")
		assert.Equal(t, c.Flavor, "")
		assert.Equal(t, c.Mode, "")
		assert.Equal(t, c.ObserverSvcURL, "")
		assert.Equal(t, c.Port, "")
		assert.Equal(t, c.TelegramSvcURL, "")
	})

	t.Run("config=success", func(t *testing.T) {
		generateFakeSuccessfulConfig(t)
		c, err := config.LoadConfig()

		assert.Nil(t, err)
		assert.NotNil(t, c)

		assert.NotEmpty(t, c.AuthSvcURL)
		assert.NotEmpty(t, c.Flavor)
		assert.NotEmpty(t, c.Mode)
		assert.NotEmpty(t, c.ObserverSvcURL)
		assert.NotEmpty(t, c.Port)
		assert.NotEmpty(t, c.TelegramSvcURL)

		assert.Equal(t, c.AuthSvcURL, "localhost:1")
		assert.Equal(t, c.Flavor, "de")
		assert.Equal(t, c.Mode, "uat")
		assert.Equal(t, c.ObserverSvcURL, "localhost:2")
		assert.Equal(t, c.Port, ":6542")
		assert.Equal(t, c.TelegramSvcURL, "localhost:3")
	})
}

func TestConfigWithFx(t *testing.T) {
	t.Run("config=injection-test", func(t *testing.T) {
		var data *config.Config
		var g fx.DotGraph

		app := fxtest.New(
			t,
			fx.Logger(fxtest.NewTestPrinter(t)),
			fx.WithLogger(func() fxevent.Logger { return fxtest.NewTestLogger(t) }),

			fx.Invoke(func() {
				generateFakeSuccessfulConfig(t)
			}),

			config.Module,
			fx.Populate(&g),
			fx.Populate(&data),
		).RequireStart()

		defer app.RequireStop()

		require.NoError(t, app.Err())
		assert.Contains(t, g, `"fx.DotGraph" [label=<fx.DotGraph>];`)

		assert.NotEmpty(t, data.AuthSvcURL)
		assert.NotEmpty(t, data.Flavor)
		assert.NotEmpty(t, data.Mode)
		assert.NotEmpty(t, data.ObserverSvcURL)
		assert.NotEmpty(t, data.Port)
		assert.NotEmpty(t, data.TelegramSvcURL)

		assert.Equal(t, data.AuthSvcURL, "localhost:1")
		assert.Equal(t, data.Flavor, "de")
		assert.Equal(t, data.Mode, "uat")
		assert.Equal(t, data.ObserverSvcURL, "localhost:2")
		assert.Equal(t, data.Port, ":6542")
		assert.Equal(t, data.TelegramSvcURL, "localhost:3")
	})
}

func generateFakeWronglConfig(tb testing.TB) {
	os.Remove("config.env")

	fakeConfig := `
	FLAVOR
	GIN_MODE=uat
	PORT=:6542
	AUTH_SVC_URL=localhost:1
	OBSERVER_SVC_URL=localhost:2
	TELEGRAM_SVC_URL=localhost:3
	`

	err := os.WriteFile("config.env", []byte(fakeConfig), 0644)
	assert.Nil(tb, err)
}

func generateFakeSuccessfulConfig(tb testing.TB) {
	os.Remove("config.env")

	fakeConfig := `
	FLAVOR=de
	GIN_MODE=uat
	PORT=:6542
	AUTH_SVC_URL=localhost:1
	OBSERVER_SVC_URL=localhost:2
	TELEGRAM_SVC_URL=localhost:3
	`

	err := os.WriteFile("config.env", []byte(fakeConfig), 0644)
	assert.Nil(tb, err)
}
