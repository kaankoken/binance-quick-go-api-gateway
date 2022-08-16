package config

import (
	"github.com/kaankoken/binance-quick-go-api-gateway/pkg/helper"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(LoadConfig))

type Config struct {
	Port           string `mapstructure:"PORT"`
	AuthSvcUrl     string `mapstructure:"AUTH_SVC_URL"`
	ObserverSvcUrl string `mapstructure:"OBSERVER_SVC_URL"`
	TelegramSvcUrl string `mapstructure:"TELEGRAM_SVC_URL"`
}

func LoadConfig(handler *helper.Handler) (c *Config, err error) {
	viper.SetConfigType("env")

	viper.AddConfigPath("$PWD")
	viper.AddConfigPath("$PWD/config/")
	viper.AddConfigPath("$PWD/config/envs/")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	handler.Error(err, nil)

	err = viper.Unmarshal(&c)
	handler.Error(err, nil)

	return
}
