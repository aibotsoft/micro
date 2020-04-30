package util

import (
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
