package helper

import (
	"github.com/kaankoken/binance-quick-go-api-gateway/config"
	"go.uber.org/fx"
)

const (
	dev string = "dev"
)

var LoggerModule = fx.Options(
	DebugModule,
	ReleaseModule,
	fx.Provide(Initialize),
)

type ILogHandler interface {
	Error(err error) error
	ErrorWithCallback(err error, f func()) error
	Info(msg string) string
	InfoWithCallback(msg string, f func()) string
}

type LogHandler struct {
	debug, release *ILogHandler
	config         *config.Config
}

func Initialize(c *config.Config, D *DLogger, R *RLogger) *LogHandler {
	var d, r ILogHandler

	d = D
	r = R

	return &LogHandler{debug: &d, release: &r, config: c}
}

func (l LogHandler) Error(err error) error {
	if l.config.Flavor == dev {
		return (*l.debug).Error(err)
	}

	return (*l.release).Error(err)
}

func (l LogHandler) ErrorWithCallback(err error, f func()) error {
	if l.config.Flavor == dev {
		return (*l.debug).ErrorWithCallback(err, f)
	}

	return (*l.release).ErrorWithCallback(err, f)
}

func (l LogHandler) Info(msg string) string {
	if l.config.Flavor == dev {
		return (*l.debug).Info(msg)
	}

	return (*l.release).Info(msg)
}

func (l LogHandler) InfoWithCallback(msg string, f func()) string {
	if l.config.Flavor == dev {
		return (*l.debug).InfoWithCallback(msg, f)
	}

	return (*l.release).InfoWithCallback(msg, f)
}
