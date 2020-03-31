package config_test

import (
	"github.com/aibotsoft/micro/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestLoadEnv(t *testing.T) {
	if !config.IsDev() {
		return
	}
	err := config.LoadEnv()
	assert.NoError(t, err)
	assert.Equal(t, "true", os.Getenv("TEST_LOAD_ENV"))
}

func TestNew(t *testing.T) {
	cfg := config.New()
	assert.Equal(t, true, cfg.Service.TestLoadEnv)
	assert.Equal(t, "dev", cfg.Service.Env)
	err := os.Setenv("ProxyService_Check_Period", "5s")
	assert.NoError(t, err)
	assert.Equal(t, 5*time.Second, cfg.ProxyService.CheckPeriod)
}
