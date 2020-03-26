package cache

import (
	"github.com/aibotsoft/micro/config"
	"github.com/dgraph-io/ristretto"
	"github.com/pkg/errors"
)

func NewCache(cfg *config.Config) *ristretto.Cache {
	c, err := ristretto.NewCache(&ristretto.Config{
		// number of keys to track frequency of (10M).
		NumCounters: 1e6,
		// maximum cost of cache (1GB).
		MaxCost:     1 << 30,
		BufferItems: 64, // number of keys per Get buffer.
		Metrics:     true,
	})
	if err != nil {
		panic(errors.Wrap(err, "error get new ristretto cache"))
	}
	return c

}
