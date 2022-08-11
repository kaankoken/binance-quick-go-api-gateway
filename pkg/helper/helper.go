package helper

import log "github.com/sirupsen/logrus"

func CheckError(err error) {
	if err != nil {
		log.Fatalln("Binance-Quick-Go-Api-Gateway -> ", err.Error())
	}
}
