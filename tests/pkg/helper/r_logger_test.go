package helper_test

import (
	"fmt"
	"testing"

	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/helper"
	"github.com/stretchr/testify/assert"
)

func init() {
	helper.InitializeLogger()
}

func TestReleaseLogger(t *testing.T) {
	l := helper.RLogger{}

	t.Run("logger-error=no-error", func(t *testing.T) {
		err := l.Error(nil)

		assert.Nil(t, err)
	})

	t.Run("logger-error=only-error", func(t *testing.T) {
		msg := "Testing error message"
		err := l.Error(fmt.Errorf(msg))

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, msg)
	})

	t.Run("logger-error=no-error&empty-callback", func(t *testing.T) {
		f := func() {}

		err := l.ErrorWithCallback(nil, f)

		assert.Nil(t, err)
	})

	t.Run("logger-error=error&empty-callback", func(t *testing.T) {
		msg := "Testing error message"
		f := func() {}

		err := l.ErrorWithCallback(fmt.Errorf(msg), f)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, msg)
	})

	t.Run("logger-error=error&callback", func(t *testing.T) {
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

	t.Run("logger-info=success", func(t *testing.T) {
		msg := "Testing info message"
		_msg := l.Info(msg)

		assert.Contains(t, _msg, msg)
	})

	t.Run("logger-info=success&empty-callback", func(t *testing.T) {
		msg := "Testing info message"
		f := func() {}

		_msg := l.InfoWithCallback(msg, f)

		assert.Contains(t, _msg, msg)
	})

	t.Run("logger-info=success&callback", func(t *testing.T) {
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
