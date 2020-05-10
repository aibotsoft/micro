package config_client

import (
	"github.com/aibotsoft/micro/config"
	"github.com/aibotsoft/micro/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConfigClient(t *testing.T) {
	cfg := config.New()
	log := logger.New()
	client := NewConfigClient(cfg, log)
	assert.NotEmpty(t, client)
}
