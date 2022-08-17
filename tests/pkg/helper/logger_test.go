package tests

import (
	"fmt"
	"testing"

	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/helper"
	"github.com/stretchr/testify/assert"
)

func init() {
	helper.SetLoggerFormat()
}

func TestLogger(t *testing.T) {
	t.Run("logger-error=no-error", func(t *testing.T) {
		err := helper.Logger().Error(nil, nil)

		assert.Nil(t, err)
	})

	t.Run("logger-error=only-error", func(t *testing.T) {
		msg := "Testing error message"
		err := helper.Logger().Error(fmt.Errorf(msg), nil)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, msg)
	})

	t.Run("logger-error=error&callback", func(t *testing.T) {
		msg := "Testing error message"
		f := func() {
			fmt.Println("Testing callback")
			assert.Nil(t, nil)
		}

		err := helper.Logger().Error(fmt.Errorf(msg), &f)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, msg)
	})

	t.Run("logger-info=success", func(t *testing.T) {
		msg := "Testing info message"
		_msg := helper.Logger().Info("Testing info message")

		assert.Contains(t, _msg, msg)
	})
}
