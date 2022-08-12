package helper

import (
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Invoke(SetLoggerFormat), fx.Provide(CheckError))

func SetLoggerFormat() {
	log.SetFormatter(&log.TextFormatter{})
}

type Handler struct {
	Logger func(error, *func())
}

func CheckError() *Handler {
	return &Handler{Logger: func(err error, f *func()) {
		if err != nil {
			log.Fatalln("Binance-Quick-Go-Api-Gateway -> ", err.Error())
			if f != nil {
				(*f)()
			}
		}
	}}
}
