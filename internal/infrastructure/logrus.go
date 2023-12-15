package infrastructure

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewLogger(config *viper.Viper) *logrus.Logger {
	logger := logrus.New()

	logger.SetLevel(logrus.Level(config.GetInt("log.level")))
	logger.SetFormatter(&logrus.JSONFormatter{})

	return logger
}
