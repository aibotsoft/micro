package sqlserver

import (
	"github.com/aibotsoft/micro/config"
	"github.com/subosito/gotenv"
	"testing"
)

func TestBuildConnString(t *testing.T) {
	gotenv.Must(gotenv.Load, "./../.env")
	cfg := config.New()
	t.Logf("%+v", cfg)
	got := BuildConnString(cfg)
	t.Log(got)
}
