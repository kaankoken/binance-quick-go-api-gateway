package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/kaankoken/binance-quick-go-api-gateway/config"
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/helper"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
}

func main() {
	c, err := config.LoadConfig()
	helper.CheckError(err)

	fmt.Println(c.Port)
}
