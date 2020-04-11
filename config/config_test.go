package config_test

import (
	"github.com/aibotsoft/micro/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestMssqlConfig(t *testing.T) {
	//gotenv.Must(gotenv.Load, "./../.env")
	t.Run("service config", func(t *testing.T) {
		_ = os.Setenv("SERVICE_NAME", "test_name")
		_ = os.Setenv("SERVICE_ENV", "test_env")
		_ = os.Setenv("SERVICE_TEST_ENV", "true")
		cfg := config.New()
		assert.Equal(t, "test_name", cfg.Service.Name)
		assert.Equal(t, "test_env", cfg.Service.Env)
		assert.Equal(t, true, cfg.Service.TestEnv)
	})
	t.Run("mssql config", func(t *testing.T) {
		_ = os.Setenv("MSSQL_HOST", "0.0.0.0")
		_ = os.Setenv("MSSQL_PORT", "999")
		_ = os.Setenv("MSSQL_USER", "test_user")
		_ = os.Setenv("MSSQL_PASSWORD", "test_pass")
		_ = os.Setenv("MSSQL_DATABASE", "test_db")
		_ = os.Setenv("MSSQL_APP_NAME", "test_appName")
		_ = os.Setenv("MSSQL_dial_Timeout", "66")
		_ = os.Setenv("MSSQL_Keep_Alive", "66")
		_ = os.Setenv("MSSQL_Packet_Size", "66")
		_ = os.Setenv("MSSQL_Log", "66")
		_ = os.Setenv("MSSQL_CONN_Timeout", "66s")
		cfg := config.New()
		assert.Equal(t, "0.0.0.0", cfg.Mssql.Host)
		assert.Equal(t, "999", cfg.Mssql.Port)
		assert.Equal(t, "test_user", cfg.Mssql.User)
		assert.Equal(t, "test_pass", cfg.Mssql.Password)
		assert.Equal(t, "test_db", cfg.Mssql.Database)
		assert.Equal(t, "test_appName", cfg.Mssql.AppName)
		assert.Equal(t, "66", cfg.Mssql.DialTimeout)
		assert.Equal(t, "66", cfg.Mssql.KeepAlive)
		assert.Equal(t, "66", cfg.Mssql.PacketSize)
		assert.Equal(t, "66", cfg.Mssql.Log)
		assert.Equal(t, 66*time.Second, cfg.Mssql.ConnTimeout)
	})
	t.Run("proxy service config", func(t *testing.T) {
		_ = os.Setenv("PROXY_SERVICE_COLLECT_PERIOD", "66s")
		_ = os.Setenv("PROXY_SERVICE_COLLECT_URL", "https://test.test")
		_ = os.Setenv("PROXY_SERVICE_collect_http_timeout", "66s")
		_ = os.Setenv("PROXY_SERVICE_grpc_timeout", "66s")
		_ = os.Setenv("PROXY_SERVICE_GRPC_PORT", "66")
		_ = os.Setenv("PROXY_SERVICE_CHECK_TIMEOUT", "66s")
		_ = os.Setenv("PROXY_SERVICE_check_period", "66s")
		_ = os.Setenv("PROXY_SERVICE_delete_bad_proxy_period", "66s")
		_ = os.Setenv("PROXY_SERVICE_delete_old_stat_period", "66s")
		cfg := config.New()
		assert.Equal(t, time.Duration(66000000000), cfg.ProxyService.CollectPeriod)
		assert.Equal(t, "https://test.test", cfg.ProxyService.CollectUrl)
		assert.Equal(t, time.Duration(66000000000), cfg.ProxyService.CollectHttpTimeout)
		assert.Equal(t, time.Duration(66000000000), cfg.ProxyService.CollectHttpTimeout)
		assert.Equal(t, time.Duration(66000000000), cfg.ProxyService.GrpcTimeout)
		assert.Equal(t, 66, cfg.ProxyService.GrpcPort)
		assert.Equal(t, time.Duration(66000000000), cfg.ProxyService.CheckTimeout)
		assert.Equal(t, time.Duration(66000000000), cfg.ProxyService.CheckPeriod)
		assert.Equal(t, time.Duration(66000000000), cfg.ProxyService.DeleteBadProxyPeriod)
		assert.Equal(t, time.Duration(66000000000), cfg.ProxyService.DeleteOldStatPeriod)
	})
	t.Run("ristretto config", func(t *testing.T) {
		_ = os.Setenv("Ristretto_Num_Counters", "66")
		_ = os.Setenv("Ristretto_Max_Cost", "66")
		_ = os.Setenv("Ristretto_Metrics", "true")
		cfg := config.New()
		assert.Equal(t, int64(66), cfg.Ristretto.NumCounters)
		assert.Equal(t, int64(66), cfg.Ristretto.MaxCost)
		assert.Equal(t, true, cfg.Ristretto.Metrics)

	})
	t.Run("migrate config", func(t *testing.T) {
		cfg := config.New()
		assert.Equal(t, "", cfg.Migrate.User)
		assert.Equal(t, "", cfg.Migrate.Token)

		_ = os.Setenv("migrate_user", "test_user")
		_ = os.Setenv("migrate_token", "test_token")
		cfg = config.New()

		assert.Equal(t, "test_user", cfg.Migrate.User)
		assert.Equal(t, "test_token", cfg.Migrate.Token)

	})
}
