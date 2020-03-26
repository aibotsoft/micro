package postgres_test

import (
	"github.com/aibotsoft/micro/config"
	"github.com/aibotsoft/micro/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/subosito/gotenv"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	cfg := config.New()
	if cfg.Service.Env != "dev" {
		return
	}
	t.Log(cfg.Service.Env)
	gotenv.Must(gotenv.Load, "../.env")
	assert.Equal(t, 5*time.Second, cfg.Postgres.Timeout)
	assert.NotEmpty(t, cfg.Postgres.Port)
	db, err := postgres.New(cfg)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, db)
	}

}
