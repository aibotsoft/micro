package config_client

import (
	"context"
	"github.com/aibotsoft/micro/config"
	"github.com/aibotsoft/micro/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

var c *ConfClient

func TestMain(m *testing.M) {
	cfg := config.New()
	log := logger.New()
	c = New(cfg, log)
	m.Run()
	c.Close()
}

//func TestConfClient_GetAccount(t *testing.T) {
//	got, err := c.GetAccount(context.Background(), "Sbobet")
//	if assert.NoError(t, err) {
//		assert.NotEmpty(t, got)
//	}
//}
//func TestConfClient_GetGrpcAddr(t *testing.T) {
//	c.cfg.Service.Name = "config-service"
//	got, err := c.GetGrpcAddr(context.Background(), c.cfg.Service.Name)
//	if assert.NoError(t, err) {
//		assert.NotEmpty(t, got)
//	}
//}

func TestConfClient_GetNetStatus(t *testing.T) {
	status := c.GetNetStatus(context.Background())
	assert.True(t, status)
}

func TestConfClient_Ping(t *testing.T) {
	err := c.Ping(context.Background())
	assert.NoError(t, err)
}
