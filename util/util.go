package util

import (
	"time"
)

func UnixMsNow() int64 {
	return time.Now().UnixNano() / 1000000
}
func UnixUsNow() int64 {
	return time.Now().UnixNano() / 1000
}
