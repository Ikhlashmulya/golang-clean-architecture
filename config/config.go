package config

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/exception"
	"github.com/spf13/viper"
)

func NewConfig() *viper.Viper {
	_, fileName, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(fileName)
	pathToFileConfig := filepath.Join(currentDir, "..")
	fmt.Println(pathToFileConfig)

	viper := viper.New()
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(pathToFileConfig)

	err := viper.ReadInConfig()
	exception.PanicIfError(err)

	return viper
}
