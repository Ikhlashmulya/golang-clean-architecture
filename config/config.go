package config

import (
	"fmt"
	"path"

	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/exception"
	"github.com/spf13/viper"
)

// type Config struct {
// }

// func NewConfig(filenames ...string) *Config {
// 	err := godotenv.Load(filenames...)
// 	exception.PanicIfError(err)

// 	return &Config{}
// }

// func (config *Config) Get(key string) string {
// 	return os.Getenv(key)
// }a

func NewConfig(pathS string) *viper.Viper {
	viper := viper.New()
	viper.SetConfigFile("config.json")
	viper.AddConfigPath(pathS)

	fmt.Println(path.Join(pathS, "config.json"))

	err := viper.ReadInConfig()
	exception.PanicIfError(err)

	return viper
}
