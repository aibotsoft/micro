package logger_test

import (
	"errors"
	"github.com/aibotsoft/micro/logger"
	"testing"
)

func TestFatal(t *testing.T) {
	log := logger.New()
	t.Run("err nil", func(t *testing.T) {
		var err error
		logger.Fatal(err, log, "error message")

	})
	t.Run("err not nil", func(t *testing.T) {
		err := errors.New("error")
		logger.Fatal(err, log, "error message")
	})

}
