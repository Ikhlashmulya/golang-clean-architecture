package config

import (
	"os"

	"github.com/Ikhlashmulya/golang-clean-architecture-project-structure/internal/exception"
	"github.com/joho/godotenv"
)

type Config struct {
}

func NewConfig(filenames ...string) *Config {
	err := godotenv.Load(filenames...)
	exception.PanicIfError(err)

	return &Config{}
}

func (config *Config) Get(key string) string {
	return os.Getenv(key)
}
