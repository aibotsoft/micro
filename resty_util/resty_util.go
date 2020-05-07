package resty_util

import (
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

func RestyTrace(log *zap.SugaredLogger, resp *resty.Response) {
	ti := resp.Request.TraceInfo()
	log.Infow("Request Trace Info", "TotalTime", ti.TotalTime, "DNSLookup", ti.DNSLookup, "ConnTime", ti.ConnTime, "TLSHandshake", ti.TLSHandshake,
		"ServerTime", ti.ServerTime, "ResponseTime", ti.ResponseTime, "IsConnReused", ti.IsConnReused, "ConnIdleTime", ti.ConnIdleTime)
	//fmt.Println("DNSLookup    :", ti.DNSLookup)
	//fmt.Println("ConnTime     :", ti.ConnTime)
	//fmt.Println("TLSHandshake :", ti.TLSHandshake)
	//fmt.Println("ServerTime   :", ti.ServerTime)
	//fmt.Println("ResponseTime :", ti.ResponseTime)
	//fmt.Println("TotalTime    :", ti.TotalTime)

	//fmt.Println("IsConnReused :", ti.IsConnReused)
	//fmt.Println("IsConnWasIdle:", ti.IsConnWasIdle)
	//fmt.Println("ConnIdleTime :", ti.ConnIdleTime)
}
