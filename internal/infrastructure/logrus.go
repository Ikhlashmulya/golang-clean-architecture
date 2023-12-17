package infrastructure

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewLogger(config *viper.Viper) *logrus.Logger {
	logger := logrus.New()

	logger.SetLevel(logrus.Level(config.GetInt("LOG_LEVEL")))
	logger.SetFormatter(&logrus.JSONFormatter{})

	return logger
}
