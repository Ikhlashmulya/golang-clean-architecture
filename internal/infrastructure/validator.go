package infrastructure

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

func NewValidator(config *viper.Viper) *validator.Validate {
	return validator.New()
}