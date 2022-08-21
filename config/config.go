package config

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(LoadConfig))

type Config struct {
	Port           string `mapstructure:"PORT"`
	Flavor         string `mapstructure:"FLAVOR"`
	Mode           string `mapstructure:"GIN_MODE"`
	AuthSvcUrl     string `mapstructure:"AUTH_SVC_URL"`
	ObserverSvcUrl string `mapstructure:"OBSERVER_SVC_URL"`
	TelegramSvcUrl string `mapstructure:"TELEGRAM_SVC_URL"`
}

func LoadConfig() (c *Config, err error) {
	viper.SetConfigType("env")

	viper.AddConfigPath("$PWD")
	viper.AddConfigPath("$PWD/config/")
	viper.AddConfigPath("$PWD/config/envs/")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	// Since the Config struct & read version struct is same
	// wont throw error
	_ = viper.Unmarshal(&c)

	return c, nil
}
