package util

import (
	"fmt"
	"github.com/agnivade/levenshtein"
	"github.com/mozillazg/go-unidecode"
	"github.com/shopspring/decimal"
	"strconv"
	"strings"
	"time"
)

const ISOFormat = "2006-01-02 15:04:05-07:00"

func UnixMsNow() int64 {
	return time.Now().UnixNano() / 1000000
}
func UnixUsNow() int64 {
	return time.Now().UnixNano() / 1000
}
func TimeId(seed int8) int64 {
	now := time.Now().UnixNano() / 100
	if seed < 1 {
		return now
	}
	strId := fmt.Sprintf("%d%d", seed, now)
	value, err := strconv.ParseInt(strId, 10, 64)
	if err != nil {
		return now
	}
	return value
}

func ToUSD(value float64, currency float64) float64 {
	currencyD := decimal.NewFromFloat(currency)
	valueD := decimal.NewFromFloat(value)
	res := valueD.Div(currencyD)
	inUSD, _ := res.Float64()
	return inUSD
}
func ToUSDInt(value float64, currency float64) int64 {
	currencyD := decimal.NewFromFloat(currency)
	valueD := decimal.NewFromFloat(value)
	return valueD.Div(currencyD).IntPart()
}

func AdaptStake(stake float64, currency float64, roundValue float64) float64 {
	currencyD := decimal.NewFromFloat(currency)
	stakeD := decimal.NewFromFloat(stake)
	roundValueD := decimal.NewFromFloat(roundValue)

	res := stakeD.Mul(currencyD).Div(roundValueD).Floor().Mul(roundValueD)
	f, _ := res.Float64()
	return f
}

func FloatToStr(value float64) string {
	return strconv.FormatFloat(value, 'g', -1, 64)
}

func RoundDown(value float64, roundValue float64) float64 {
	rv := decimal.NewFromFloat(roundValue)
	res, _ := decimal.NewFromFloat(value).Div(rv).Floor().Mul(rv).Float64()
	return res
}

// Truncate truncates off digits from the number, without rounding.
//
// NOTE: precision is the last digit that will not be truncated (must be >= 0).
//
// Example:
//
//     decimal.NewFromString("123.456").Truncate(2).String() // "123.45"
//
func TruncateFloat(value float64, precision int32) float64 {
	res, _ := decimal.NewFromFloat(value).Truncate(precision).Float64()
	return res
}
func MaxFloat(values ...float64) (max float64) {
	for i := range values {
		if values[i] > max {
			max = values[i]
		}
	}
	return
}

func PtrFloat64(v float64) *float64 { return &v }
func PtrString(v string) *string    { return &v }

func Normalize(s string) string {
	return strings.TrimSpace(strings.ToLower(unidecode.Unidecode(s)))
}
func Compare(a string, b string) float64 {
	distance := levenshtein.ComputeDistance(Normalize(a), Normalize(b))
	return TruncateFloat(1-float64(distance)/float64(len(a+b)), 3)
}
func StringInList(needle string, haystack []string) bool {
	for i := range haystack {
		if needle == haystack[i] {
			return true
		}
	}
	return false
}

func Int64InList(needle int64, haystack []int64) bool {
	for i := range haystack {
		if needle == haystack[i] {
			return true
		}
	}
	return false
}
