package config

import (
	"go.uber.org/zap"
	"os"
)

func Logger() *zap.Logger {
	goEnv := os.Getenv("GO_ENV")

	if goEnv == "production" {
		logger, _ := zap.NewProduction()

		return logger
	} else {
		logger, _ := zap.NewDevelopment()

		return logger
	}
}
