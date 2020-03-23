package logger

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Fatal checks if err is not nil then log message and calls os.Exit.
func Fatal(err error, log *zap.SugaredLogger, message string) {
	if err == nil {
		return
	}
	log.Fatal(errors.WithMessage(err, message))
}
