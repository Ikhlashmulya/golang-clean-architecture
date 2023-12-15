package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func New() *viper.Viper {
	viper := viper.New()
	viper.AddConfigPath("../")
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error loading configuration : %+v", err))
	}

	return viper
}
