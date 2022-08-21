package helper_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/kaankoken/binance-quick-go-api-gateway/config"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/fx/fxtest"
)

func TestHandlerLogger(t *testing.T) {
	// Needed Module initialization
	// debug logger
	d := helper.SetLoggerFormat()
	// release logger
	logger := helper.InitializeLogger()
	r := helper.InitializeLoggerPtr(logger)

	//config
	generateFakeSuccessfulConfig(t, "dev")
	c, err := config.LoadConfig()
	assert.Nil(t, err)

	// log handler initialization
	l := helper.Initialize(c, d, r)

	t.Parallel()

	t.Run("handler-logger-error=no-error", func(t *testing.T) {
		t.Parallel()

		err := l.Error(nil)

		assert.Nil(t, err)
	})

	t.Run("handler-logger-error=only-error", func(t *testing.T) {
		t.Parallel()

		msg := "Testing error message"
		err := l.Error(fmt.Errorf(msg))

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, msg)
	})

	t.Run("handler-logger-error=no-error&empty-callback", func(t *testing.T) {
		t.Parallel()

		f := func() {}
		err := l.ErrorWithCallback(nil, f)

		assert.Nil(t, err)
	})

	t.Run("handler-logger-error=error&empty-callback", func(t *testing.T) {
		t.Parallel()

		msg := "Testing error message"
		f := func() {}

		err := l.ErrorWithCallback(fmt.Errorf(msg), f)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, msg)
	})

	t.Run("handler-logger-error=error&callback", func(t *testing.T) {
		t.Parallel()

		msg := "Testing error message"
		f := func() {
			msg := fmt.Sprintf("%s %d", "Testing callback", 1)

			assert.Nil(t, nil)
			assert.NotNil(t, msg)
		}

		err := l.ErrorWithCallback(fmt.Errorf(msg), f)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, msg)
	})

	t.Run("handler-logger-info=success", func(t *testing.T) {
		t.Parallel()

		msg := "Testing info message"
		_msg := l.Info(msg)

		assert.Contains(t, _msg, msg)
	})

	t.Run("handler-logger-info=success&empty-callback", func(t *testing.T) {
		t.Parallel()

		msg := "Testing info message"
		f := func() {}

		_msg := l.InfoWithCallback(msg, f)

		assert.Contains(t, _msg, msg)
	})

	t.Run("handler-logger-info=success&callback", func(t *testing.T) {
		t.Parallel()

		msg := "Testing info message"
		f := func() {
			msg := fmt.Sprintf("%s %d", "Testing callback", 1)
			assert.Nil(t, nil)
			assert.NotNil(t, msg)
		}

		_msg := l.InfoWithCallback(msg, f)

		assert.Contains(t, _msg, msg)
	})
}

func TestHandlerLoggerWithFx(t *testing.T) {
	t.Parallel()

	generateFakeSuccessfulConfig(t, "release")

	t.Run("handler-logger=injection-test", func(t *testing.T) {
		var g fx.DotGraph

		app := fxtest.New(
			t,
			fx.Logger(fxtest.NewTestPrinter(t)),
			fx.WithLogger(func() fxevent.Logger { return fxtest.NewTestLogger(t) }),
			config.Module,
			helper.LoggerModule,
			fx.Populate(&g),
		).RequireStart()

		defer app.RequireStop()

		require.NoError(t, app.Err())
		assert.Contains(t, g, `"fx.DotGraph" [label=<fx.DotGraph>];`)
	})

	t.Run("handler-logger=injection-test-with-functions", func(t *testing.T) {
		var g fx.DotGraph
		var l *helper.LogHandler

		app := fxtest.New(
			t,
			fx.Logger(fxtest.NewTestPrinter(t)),
			fx.WithLogger(func() fxevent.Logger { return fxtest.NewTestLogger(t) }),
			config.Module,
			helper.LoggerModule,
			fx.Populate(&g),
			fx.Populate(&l),
		).RequireStart()

		defer app.RequireStop()

		require.NoError(t, app.Err())
		assert.Contains(t, g, `"fx.DotGraph" [label=<fx.DotGraph>];`)

		t.Run("handler-logger-error=no-error", func(t *testing.T) {
			t.Parallel()
			err := l.Error(nil)

			assert.Nil(t, err)
		})

		t.Run("handler-logger-error=only-error", func(t *testing.T) {
			t.Parallel()

			msg := "Testing error message"
			err := l.Error(fmt.Errorf(msg))

			assert.NotNil(t, err)
			assert.ErrorContains(t, err, msg)
		})

		t.Run("handler-logger-error=no-error&empty-callback", func(t *testing.T) {
			t.Parallel()

			f := func() {}
			err := l.ErrorWithCallback(nil, f)

			assert.Nil(t, err)
		})

		t.Run("handler-logger-error=error&empty-callback", func(t *testing.T) {
			t.Parallel()

			msg := "Testing error message"
			f := func() {}

			err := l.ErrorWithCallback(fmt.Errorf(msg), f)

			assert.NotNil(t, err)
			assert.ErrorContains(t, err, msg)
		})

		t.Run("handler-logger-error=error&callback", func(t *testing.T) {
			t.Parallel()

			msg := "Testing error message"
			f := func() {
				msg := fmt.Sprintf("%s %d", "Testing callback", 1)

				assert.Nil(t, nil)
				assert.NotNil(t, msg)
			}

			err := l.ErrorWithCallback(fmt.Errorf(msg), f)

			assert.NotNil(t, err)
			assert.ErrorContains(t, err, msg)
		})

		t.Run("handler-logger-info=success", func(t *testing.T) {
			t.Parallel()

			msg := "Testing info message"
			_msg := l.Info(msg)

			assert.Contains(t, _msg, msg)
		})

		t.Run("handler-logger-info=success&empty-callback", func(t *testing.T) {
			t.Parallel()

			msg := "Testing info message"
			f := func() {}

			_msg := l.InfoWithCallback(msg, f)

			assert.Contains(t, _msg, msg)
		})

		t.Run("handler-logger-info=success&callback", func(t *testing.T) {
			t.Parallel()

			msg := "Testing info message"
			f := func() {
				msg := fmt.Sprintf("%s %d", "Testing callback", 1)
				assert.Nil(t, nil)
				assert.NotNil(t, msg)
			}

			_msg := l.InfoWithCallback(msg, f)

			assert.Contains(t, _msg, msg)
		})
	})
}

func generateFakeSuccessfulConfig(tb testing.TB, flavor string) {
	os.Remove("config.env")

	fakeConfig := `
	FLAVOR=%s
	GIN_MODE=uat
	PORT=:6542
	AUTH_SVC_URL=localhost:1
	OBSERVER_SVC_URL=localhost:2
	TELEGRAM_SVC_URL=localhost:3
	`

	fakeConfig = fmt.Sprintf(fakeConfig, flavor)

	err := os.WriteFile("config.env", []byte(fakeConfig), 0644)
	assert.Nil(tb, err)
}
