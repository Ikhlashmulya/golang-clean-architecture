package config

import (
	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/exception"
	"github.com/spf13/viper"
)

func NewConfig(pathString string) *viper.Viper {
	viper := viper.New()
	viper.SetConfigFile("config.json")
	viper.AddConfigPath(pathString)

	err := viper.ReadInConfig()
	exception.PanicIfError(err)

	return viper
}
