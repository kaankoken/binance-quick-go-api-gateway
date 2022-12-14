package helper_test

import (
	"fmt"
	"testing"

	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/fx/fxtest"
)

func TestDebugLogger(t *testing.T) {
	l := helper.SetLoggerFormat()

	t.Parallel()

	t.Run("d-logger-error=no-error", func(t *testing.T) {
		t.Parallel()

		err := l.Error(nil)

		assert.Nil(t, err)
	})

	t.Run("d-logger-error=only-error", func(t *testing.T) {
		t.Parallel()

		msg := "Testing error message"
		err := l.Error(fmt.Errorf(msg))

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, msg)
	})

	t.Run("d-logger-error=no-error&empty-callback", func(t *testing.T) {
		t.Parallel()

		f := func() {}
		err := l.ErrorWithCallback(nil, f)

		assert.Nil(t, err)
	})

	t.Run("d-logger-error=error&empty-callback", func(t *testing.T) {
		t.Parallel()

		msg := "Testing error message"
		f := func() {}

		err := l.ErrorWithCallback(fmt.Errorf(msg), f)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, msg)
	})

	t.Run("d-logger-error=error&callback", func(t *testing.T) {
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

	t.Run("d-logger-info=success", func(t *testing.T) {
		t.Parallel()

		msg := "Testing info message"
		_msg := l.Info(msg)

		assert.Contains(t, _msg, msg)
	})

	t.Run("d-logger-info=success&empty-callback", func(t *testing.T) {
		t.Parallel()

		msg := "Testing info message"
		f := func() {}

		_msg := l.InfoWithCallback(msg, f)

		assert.Contains(t, _msg, msg)
	})

	t.Run("d-logger-info=success&callback", func(t *testing.T) {
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

func TestDebugLoggerWithFx(t *testing.T) {
	t.Parallel()

	t.Run("d-logger=injection-test", func(t *testing.T) {
		var g fx.DotGraph

		app := fxtest.New(
			t,
			fx.Logger(fxtest.NewTestPrinter(t)),
			fx.WithLogger(func() fxevent.Logger { return fxtest.NewTestLogger(t) }),
			helper.DebugModule,
			fx.Populate(&g),
		).RequireStart()

		defer app.RequireStop()

		require.NoError(t, app.Err())
		assert.Contains(t, g, `"fx.DotGraph" [label=<fx.DotGraph>];`)
	})

	t.Run("d-logger=injection-test-with-functions", func(t *testing.T) {
		var g fx.DotGraph
		var l *helper.DLogger

		app := fxtest.New(
			t,
			fx.Logger(fxtest.NewTestPrinter(t)),
			fx.WithLogger(func() fxevent.Logger { return fxtest.NewTestLogger(t) }),
			helper.DebugModule,
			fx.Populate(&g),
			fx.Populate(&l),
		).RequireStart()

		defer app.RequireStop()

		require.NoError(t, app.Err())
		assert.Contains(t, g, `"fx.DotGraph" [label=<fx.DotGraph>];`)

		t.Run("d-logger-error=no-error", func(t *testing.T) {
			t.Parallel()

			err := l.Error(nil)

			assert.Nil(t, err)
		})

		t.Run("d-logger-error=only-error", func(t *testing.T) {
			t.Parallel()

			msg := "Testing error message"
			err := l.Error(fmt.Errorf(msg))

			assert.NotNil(t, err)
			assert.ErrorContains(t, err, msg)
		})

		t.Run("d-logger-error=no-error&empty-callback", func(t *testing.T) {
			t.Parallel()

			f := func() {}
			err := l.ErrorWithCallback(nil, f)

			assert.Nil(t, err)
		})

		t.Run("d-logger-error=error&empty-callback", func(t *testing.T) {
			t.Parallel()

			msg := "Testing error message"
			f := func() {}

			err := l.ErrorWithCallback(fmt.Errorf(msg), f)

			assert.NotNil(t, err)
			assert.ErrorContains(t, err, msg)
		})

		t.Run("d-logger-error=error&callback", func(t *testing.T) {
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

		t.Run("d-logger-info=success", func(t *testing.T) {
			t.Parallel()

			msg := "Testing info message"
			_msg := l.Info(msg)

			assert.Contains(t, _msg, msg)
		})

		t.Run("d-logger-info=success&empty-callback", func(t *testing.T) {
			t.Parallel()

			msg := "Testing info message"
			f := func() {}

			_msg := l.InfoWithCallback(msg, f)

			assert.Contains(t, _msg, msg)
		})

		t.Run("d-logger-info=success&callback", func(t *testing.T) {
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
