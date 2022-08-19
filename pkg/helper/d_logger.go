package helper

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var DebugModule = fx.Options(
	fx.Invoke(SetLoggerFormat),
)

const (
	tag string = "Binance-Quick-Go-Api-Gateway -> "
)

func SetLoggerFormat() {
	log.SetFormatter(&log.TextFormatter{})
}

type dLogger struct {
}

func (logger dLogger) Error(err error) error {
	if err != nil {
		log.Errorf(tag, err.Error())

		return fmt.Errorf(tag + err.Error())
	}

	return nil
}

func (logger dLogger) ErrorWithCallback(err error, f func()) error {
	if err != nil {
		f()
		log.Errorf(tag, err.Error())

		return fmt.Errorf(tag + err.Error())
	}

	return nil
}

func (logger dLogger) Info(msg string) string {
	log.Infoln(tag, msg)

	return fmt.Sprintf(tag + msg)
}

func (logger dLogger) InfoWithCallback(msg string, f func()) string {
	f()
	log.Infoln(tag, msg)

	return fmt.Sprintf(tag + msg)
}
