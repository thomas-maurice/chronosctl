package config

import (
	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.SetConfigName("chronosctl")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/")

	viper.SetDefault("url", "http://localhost:4400")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}
