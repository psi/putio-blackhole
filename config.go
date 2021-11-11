package main

import (
	"github.com/spf13/viper"
)

type Config struct {
	BaseURL    string
	PutIOToken string
	WatchDir   string
}

func LoadConfig() Config {
	config := Config{}

	viper.AutomaticEnv()

	config.BaseURL = viper.Get("base_url").(string)
	config.PutIOToken = viper.Get("putio_token").(string)
	config.WatchDir = viper.Get("watchdir").(string)

	return config
}
