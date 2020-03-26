package cache_test

import (
	"github.com/aibotsoft/micro/cache"
	"github.com/aibotsoft/micro/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestNewCache_Config(t *testing.T) {
	os.Setenv("RISTRETTO_NUMCOUNTERS", "1")
	os.Setenv("RISTRETTO_MAXCOST", "1")
	cfg := config.New()
	assert.True(t, cfg.Ristretto.Metrics)
	assert.Equal(t, int64(1), cfg.Ristretto.NumCounters)
	assert.Equal(t, int64(1), cfg.Ristretto.MaxCost)
}

func TestNewCache(t *testing.T) {
	c := cache.NewCache(config.New())
	assert.NotEmpty(t, c)
	got := c.Set("key", "test", 1)
	assert.True(t, got)
	// Cache needs time to settle
	time.Sleep(time.Millisecond)

	get, b := c.Get("key")
	assert.True(t, b)
	assert.Equal(t, "test", get)
}
