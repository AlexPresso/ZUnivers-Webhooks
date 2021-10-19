package utils

import (
	"github.com/spf13/viper"
	"log"
)

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}
