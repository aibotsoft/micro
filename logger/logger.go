package logger

import (
	"go.uber.org/zap"
	"log"
	"os"
)

const devEnv = "dev"

func newDevLogger() *zap.SugaredLogger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("newDevLogger error: %v", err)
	}
	return logger.Sugar()
}
func newProdLogger() *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("newProdLogger error: %v", err)
	}
	return logger.Sugar()
}

func New() *zap.SugaredLogger {
	if os.Getenv("SERVICE_ENV") == devEnv {
		return newDevLogger()
	}
	return newProdLogger()
}
