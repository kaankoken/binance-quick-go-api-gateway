package helper

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Invoke(SetLoggerFormat),
	fx.Provide(Logger),
)

func SetLoggerFormat() {
	log.SetFormatter(&log.TextFormatter{})
}

type Handler struct {
	Error func(error, *func()) error
	Info  func(msg string) string
}

func Logger() *Handler {
	return &Handler{
		Error: func(err error, f *func()) error {
			if err != nil {
				log.Errorf("Binance-Quick-Go-Api-Gateway -> ", err.Error())
				if f != nil {
					(*f)()
				}

				return fmt.Errorf("Binance-Quick-Go-Api-Gateway -> " + err.Error())
			}

			return nil
		},
		Info: func(msg string) string {
			log.Infoln("Binance-Quick-Go-Api-Gateway -> ", msg)

			return fmt.Sprintf("Binance-Quick-Go-Api-Gateway -> " + msg)
		},
	}
}
