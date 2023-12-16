package config

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

func New() *viper.Viper {
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)
	configFile := path.Join(currentDir, ".." ,".env")

	viper := viper.New()
	viper.SetConfigFile(configFile)
	viper.AutomaticEnv() //use OS environment variable

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error loading configuration : %+v", err))
	}

	return viper
}
