package logger_test

import (
	"github.com/aibotsoft/micro/logger"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewLogger(t *testing.T) {
	t.Run("new dev logger", func(t *testing.T) {
		err := os.Setenv("SERVICE_ENV", "dev")
		assert.Nil(t, err)
		log := logger.New()
		log.Info("dev logger")
	})
	t.Run("new prod logger", func(t *testing.T) {
		err := os.Setenv("SERVICE_ENV", "prod")
		assert.Nil(t, err)
		log := logger.New()
		log.Info("prod logger")
	})
	t.Run("new prod logger", func(t *testing.T) {
		err := os.Setenv("SERVICE_ENV", "")
		assert.Nil(t, err)
		log := logger.New()
		log.Info("unset SERVICE_ENV logger")
	})

}
