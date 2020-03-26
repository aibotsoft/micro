package cache_test

import (
	"github.com/aibotsoft/micro/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCache(t *testing.T) {
	cfg := config.New()
	assert.True(t, cfg.Ristretto.Metrics)
	assert.Equal(t, 1e5, cfg.Ristretto.NumCounters)

	//c := cache.NewCache()
	//assert.NotEmpty(t, c)
	//got := c.Set("key", "test", 1)
	//assert.True(t, got)
	//// Cache needs time to settle
	//time.Sleep(time.Millisecond)
	//
	//get, b := c.Get("key")
	//assert.True(t, b)
	//assert.Equal(t, "test", get)

}
