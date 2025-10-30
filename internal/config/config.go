package config
package config

import (
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/tg/")
	viper.AddConfigPath("$HOME/.tg")
	viper.AutomaticEnv()

	viper.SetDefault("keycloak.url", "http://localhost:8080/auth")
	viper.SetDefault("keycloak.realm", "master")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
		}
	}
}
