package pkg_test

import (
	"os"
	"testing"

	"github.com/kaankoken/binance-quick-go-api-gateway/config"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/fx/fxtest"
)

func TestGinHandler(t *testing.T) {
	t.Parallel()

	t.Run("gin-handler=with-wrong-mode", func(t *testing.T) {
		generateFakeWrongConfig(t)

		c, err := config.LoadConfig()
		_, err2 := pkg.NewHandler(c)

		assert.Nil(t, err)
		assert.NotNil(t, c)
		assert.NotNil(t, err2)
	})

	t.Run("gin-handler=success", func(t *testing.T) {
		generateFakeSuccessfulConfig(t)
		c, err := config.LoadConfig()

		handler, err2 := pkg.NewHandler(c)

		assert.Nil(t, err)
		assert.NotNil(t, c)
		assert.Nil(t, err2)
		assert.IsType(t, handler, &pkg.Handler{})
	})
}

func TestGinHandlerWithFx(t *testing.T) {
	t.Parallel()

	t.Run("gin-handler-injection=with-wrong-mode", func(t *testing.T) {
		generateFakeWrongConfig(t)

		var g fx.DotGraph

		app := fxtest.New(
			t,
			fx.Logger(fxtest.NewTestPrinter(t)),
			fx.WithLogger(func() fxevent.Logger { return fxtest.NewTestLogger(t) }),
			pkg.EngineModule,
			config.Module,
			fx.Populate(&g),
		).RequireStart()

		defer app.RequireStop()

		require.NoError(t, app.Err())
		assert.Contains(t, g, `"fx.DotGraph" [label=<fx.DotGraph>];`)
	})

	t.Run("gin-handler-injection=success", func(t *testing.T) {
		generateFakeSuccessfulConfig(t)

		var g fx.DotGraph
		var handler *pkg.Handler

		app := fxtest.New(
			t,
			fx.Logger(fxtest.NewTestPrinter(t)),
			fx.WithLogger(func() fxevent.Logger { return fxtest.NewTestLogger(t) }),
			pkg.EngineModule,
			config.Module,
			fx.Populate(&g),
			fx.Populate(&handler),
		).RequireStart()

		defer app.RequireStop()

		require.NoError(t, app.Err())
		assert.Contains(t, g, `"fx.DotGraph" [label=<fx.DotGraph>];`)
	})
}

func generateFakeWrongConfig(tb testing.TB) {
	os.Remove("config.env")

	fakeConfig := `
	FLAVOR=de
	GIN_MODE=releasdase
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
	GIN_MODE=test
	PORT=:6542
	AUTH_SVC_URL=localhost:1
	OBSERVER_SVC_URL=localhost:2
	TELEGRAM_SVC_URL=localhost:3
	`

	err := os.WriteFile("config.env", []byte(fakeConfig), 0644)
	assert.Nil(tb, err)
}
