package postgres_test

import (
	"github.com/aibotsoft/micro/config"
	"github.com/aibotsoft/micro/postgres"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	cfg := config.New()
	db, err := postgres.New(cfg)
	assert.NoError(t, err)
	assert.NotEmpty(t, db)

}
