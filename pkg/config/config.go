package config

import (
	"log"

	"github.com/spf13/viper" ///lib to read local database
)

func GetConfig() {
	viper.SetConfigName("App")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configurations")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("config error : ", err.Error())
	}
}
