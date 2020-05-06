package util

import (
	"github.com/shopspring/decimal"
	"strconv"
	"time"
)

func UnixMsNow() int64 {
	return time.Now().UnixNano() / 1000000
}
func UnixUsNow() int64 {
	return time.Now().UnixNano() / 1000
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
