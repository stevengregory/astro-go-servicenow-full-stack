package config

import (
	"log"

	"github.com/spf13/viper"
)

type ServiceNowConfig struct {
	Username string
	Password string
	Instance string
}

func Init() *ServiceNowConfig {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	viper.SetEnvPrefix("SN")
	viper.AutomaticEnv()

	return &ServiceNowConfig{
		Username: viper.GetString("username"),
		Password: viper.GetString("password"),
		Instance: viper.GetString("instance"),
	}
}
