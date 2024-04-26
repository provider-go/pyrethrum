package config

import "github.com/spf13/viper"

func Start() {
	viper.AddConfigPath(".")
	viper.SetConfigFile("yaml/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
