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
	gotenv.Must(gotenv.Load, "../.env")
	cfg := config.New()
	assert.Equal(t, 5*time.Second, cfg.Postgres.Timeout)
	assert.NotEmpty(t, cfg.Postgres.Timeout, cfg.Postgres.Timeout)
	assert.NotEmpty(t, cfg.Postgres.Port, cfg.Postgres)
	assert.NotEmpty(t, cfg.Postgres.User)
	db, err := postgres.New(cfg)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, db)
	}
}
