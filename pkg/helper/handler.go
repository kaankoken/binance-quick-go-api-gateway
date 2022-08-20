package helper

import (
	"fmt"

	"github.com/kaankoken/binance-quick-go-api-gateway/config"
	"go.uber.org/fx"
)

const (
	dev string = "dev"
)

var LoggerModule = fx.Options(
	DebugModule,
	ReleaseModule,
	fx.Provide(initialize),
)

type logHandler interface {
	Error(err error) error
	ErrorWithCallback(err error, f func()) error
	Info(msg string) string
	InfoWithCallback(msg string, f func()) string
}

type LogHandler struct {
	debug, release *logHandler
	config         *config.Config
}

func initialize(c *config.Config) *LogHandler {
	var d logHandler
	var r logHandler

	d = DLogger{}
	r = RLogger{}

	return &LogHandler{debug: &d, release: &r, config: c}
}

func (l *LogHandler) Error(err error) error {
	if l.config.Flavor == dev {
		(*l.debug).Error(err)
	}

	(*l.release).Error(err)

	return nil
}

func (l *LogHandler) ErrorWithCallback(err error, f func()) error {
	if l.config.Flavor == dev {
		(*l.debug).ErrorWithCallback(err, f)
	}

	(*l.release).ErrorWithCallback(err, f)

	return nil
}

func (l *LogHandler) Info(msg string) string {
	if l.config.Flavor == dev {
		(*l.debug).Info(msg)
	}

	(*l.release).Info(msg)

	return fmt.Sprintf(tag + msg)
}

func (l *LogHandler) InfoWithCallback(msg string, f func()) string {
	if l.config.Flavor == dev {
		(*l.debug).InfoWithCallback(msg, f)
	}

	(*l.release).InfoWithCallback(msg, f)

	return fmt.Sprintf(tag + msg)
}
