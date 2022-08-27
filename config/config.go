package config

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

// Module -> Dependency Injection for Config
var Module = fx.Options(fx.Provide(LoadConfig))

/*
Config -> Data Model for Config

[Port] -> Http server port number
[Flavor] -> It could be dev/uat/prod etc
[Mode] -> It likes debug or release
[AuthSvcUrl] -> Auth server port number
[ObserverSvcUrl] -> Observer server port number
[TelegramSvcUrl] -> Telegram server port number
*/
type Config struct {
	Port           string `mapstructure:"PORT"`
	Flavor         string `mapstructure:"FLAVOR"`
	Mode           string `mapstructure:"GIN_MODE"`
	AuthSvcURL     string `mapstructure:"AUTH_SVC_URL"`
	ObserverSvcURL string `mapstructure:"OBSERVER_SVC_URL"`
	TelegramSvcURL string `mapstructure:"TELEGRAM_SVC_URL"`
}

/*
LoadConfig -> reading config.env using viper
[return] -> returns {Config Data Model} if reads config.env or {error} cannot reads or unmarshal it
*/
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
