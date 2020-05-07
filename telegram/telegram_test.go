package telegram

import (
	"github.com/aibotsoft/micro/config"
	"github.com/aibotsoft/micro/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func InitHelper(t *testing.T) *Telegram {
	t.Helper()
	cfg := config.New()
	log := logger.New()
	return New(cfg, log)
}
func TestTelegram_Send(t *testing.T) {
	tel := InitHelper(t)
	tel.Send("test")
}

func TestTelegram_Ping(t *testing.T) {
	tel := InitHelper(t)
	err := tel.Ping()
	assert.NoError(t, err)
}
