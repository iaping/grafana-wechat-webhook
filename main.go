package main

import (
	"github.com/iaping/grafana-wechat-webhook/webhook"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./env.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	var config webhook.Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	if err := webhook.New(config).ListenAndServe(); err != nil {
		panic(err)
	}
}
