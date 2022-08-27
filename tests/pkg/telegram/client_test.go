package telegram_test

import (
	"os"
	"testing"

	"github.com/kaankoken/binance-quick-go-api-gateway/config"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/helper"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/telegram"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/telegram/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/fx/fxtest"
)

func TestClient(t *testing.T) {
	// Needed Module initialization
	// debug logger
	d := helper.SetLoggerFormat()
	// release logger
	logger := helper.InitializeLogger()
	r := helper.InitializeLoggerPtr(logger)
	generateFakeSuccessfulConfig(t)

	t.Parallel()

	t.Run("client-init=success", func(t *testing.T) {
		c, err := config.LoadConfig()

		// log handler initialization
		l := helper.Initialize(c, d, r)

		service := telegram.InitServiceClient(c, l)
		var x pb.TelegramServiceClient

		assert.Nil(t, err)
		assert.NotNil(t, c)
		assert.IsType(t, &service, &x)
	})

	t.Run("client-init=svc-success", func(t *testing.T) {
		var compareType pb.TelegramServiceClient
		var structCompareType telegram.ServiceClient

		c, err := config.LoadConfig()

		// log handler initialization
		l := helper.Initialize(c, d, r)

		service := telegram.InitServiceClient(c, l)
		res := telegram.Initialize(compareType, l)

		assert.Nil(t, err)
		assert.NotNil(t, c)
		assert.IsType(t, &service, &compareType)
		assert.IsType(t, &structCompareType, &res)
	})

	t.Run("client-init=svc-error", func(t *testing.T) {
		var compareType pb.TelegramServiceClient
		var structCompareType telegram.ServiceClient

		c, err := config.LoadConfig()

		// log handler initialization
		l := helper.Initialize(c, d, r)
		c.TelegramSvcURL = ""
		service := telegram.InitServiceClient(c, l)
		res := telegram.Initialize(compareType, l)

		assert.Nil(t, err)
		assert.NotNil(t, c)
		assert.IsType(t, &service, &compareType)
		assert.IsType(t, &structCompareType, &res)
	})
}

func TestClientWithFx(t *testing.T) {
	generateFakeSuccessfulConfig(t)
	t.Parallel()

	t.Run("client-init-injection=success", func(t *testing.T) {
		var g fx.DotGraph
		var h pb.TelegramServiceClient

		app := fxtest.New(
			t,
			fx.Logger(fxtest.NewTestPrinter(t)),
			fx.WithLogger(func() fxevent.Logger { return fxtest.NewTestLogger(t) }),
			config.Module,
			telegram.ClientModule,
			helper.LoggerModule,
			fx.Populate(&h),
			fx.Populate(&g),
		).RequireStart()

		defer app.RequireStop()

		require.NoError(t, app.Err())
		assert.Contains(t, g, `"fx.DotGraph" [label=<fx.DotGraph>];`)
	})

	t.Run("client-init-injection=svc-success", func(t *testing.T) {
		var g fx.DotGraph
		var h telegram.ServiceClient

		app := fxtest.New(
			t,
			fx.Logger(fxtest.NewTestPrinter(t)),
			fx.WithLogger(func() fxevent.Logger { return fxtest.NewTestLogger(t) }),
			config.Module,
			telegram.ClientModule,
			helper.LoggerModule,
			fx.Populate(&h),
			fx.Populate(&g),
		).RequireStart()

		defer app.RequireStop()

		require.NoError(t, app.Err())
		assert.Contains(t, g, `"fx.DotGraph" [label=<fx.DotGraph>];`)
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
