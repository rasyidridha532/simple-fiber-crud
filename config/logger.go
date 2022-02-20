package config

import (
	"go.uber.org/zap"
	"os"
)

func Logger() *zap.Logger {
	goEnv := os.Getenv("GO_ENV")

	if goEnv == "production" {
		logger, _ := zap.NewProduction()
		defer logger.Sync()

		return logger
	} else {
		logger, _ := zap.NewDevelopment()
		defer logger.Sync()

		return logger
	}
}
