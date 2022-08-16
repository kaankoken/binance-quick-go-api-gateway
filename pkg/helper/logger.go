package helper

import (
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
	Error func(error, *func())
	Info  func(msg string)
}

func Logger() *Handler {
	return &Handler{
		Error: func(err error, f *func()) {
			if err != nil {
				log.Fatalln("Binance-Quick-Go-Api-Gateway -> ", err.Error())
				if f != nil {
					(*f)()
				}
			}
		},
		Info: func(msg string) {
			log.Infoln("Binance-Quick-Go-Api-Gateway -> ", msg)
		},
	}
}
