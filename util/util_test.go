package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnixUsNow(t *testing.T) {
	got := UnixUsNow()
	t.Log(got)
}

func TestUnixMsNow(t *testing.T) {
	got := UnixMsNow()
	t.Log(got)
}

func TestAdaptStake(t *testing.T) {
	got := AdaptStake(1, 1, 1)
	assert.Equal(t, float64(1), got)
	got = AdaptStake(2.05, 1, 1)
	assert.Equal(t, float64(2), got)
	t.Log(got)
	got = AdaptStake(1.99, 1, 0.1)
	assert.Equal(t, float64(1.9), got)
	t.Log(got)

	got = AdaptStake(2.99, 1, 3)
	assert.Equal(t, float64(2.9), got)
	t.Log(got)
}

func TestFloatToStr(t *testing.T) {
	got := FloatToStr(1.49)
	t.Log(got)
	got = FloatToStr(1)
	t.Log(got)

}
