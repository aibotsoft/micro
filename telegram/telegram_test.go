package telegram

import (
	"github.com/aibotsoft/micro/config"
	"github.com/aibotsoft/micro/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tel *Telegram

func TestMain(m *testing.M) {
	cfg := config.New()
	log := logger.New()
	tel = New(cfg, log)
	m.Run()
}

func TestTelegram_Send(t *testing.T) {
	tel.Send("test")
}

func TestTelegram_Ping(t *testing.T) {
	err := tel.Ping()
	assert.NoError(t, err)
}

func TestTelegram_Sendf(t *testing.T) {
	tel.Sendf("asdf %v, %v, %s", 1, "b", "c")
	t.Log("i am")
}
