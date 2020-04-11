package mig

import (
	"github.com/aibotsoft/micro/config"
	"github.com/aibotsoft/micro/logger"
	"github.com/aibotsoft/micro/sqlserver"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUp(t *testing.T) {
	cfg := config.New()
	log := logger.New()
	db := sqlserver.MustConnect(cfg)
	err := MigrateUp(cfg, log, db)
	assert.NoError(t, err)
}

//
//func TestUpTo(t *testing.T) {
//	targetVersion := 1
//	db, err := storage.NewStorage()
//	assert.Nil(t, err)
//	err = UpTo(db, targetVersion)
//	assert.Nil(t, err)
//}
