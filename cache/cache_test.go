package cache_test

import (
	"github.com/aibotsoft/micro/cache"
	"github.com/aibotsoft/micro/config"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

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
