package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func New() *viper.Viper {
	viper := viper.New()
	viper.AddConfigPath("../")
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv() //use OS environment variable if exists

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error loading configuration : %+v", err))
	}

	return viper
}
