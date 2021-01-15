package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnixUsNow(t *testing.T) {
	got := UnixUsNow()
	t.Log(got)

	//9223372036854775807
	//1595760433707325
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

func TestTruncateFloat(t *testing.T) {
	got := TruncateFloat(1.0059, 3)
	assert.Equal(t, 1.005, got)
}

func TestMaxFloat(t *testing.T) {
	got := MaxFloat(0, 1, 2, 3)
	assert.Equal(t, float64(3), got)
}

func TestTimeId(t *testing.T) {
	got := TimeId(5)
	t.Log(got)
	//9223372036854775807
	//115957610403641252
	//15957605681168675

}
