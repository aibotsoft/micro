package config

import (
	"time"
)

// Config provides the system configuration.
type Config struct {
	Service       Service
	Mssql         Mssql
	ProxyService  ProxyService
	Ristretto     Ristretto
	Migrate       Migrate
	PinService    PinService
	SboService    SboService
	FortedService FortedService
	Telegram      Telegram
}
type Telegram struct {
	Token  string `envconfig:"optional"`
	ChatId string `envconfig:"optional"`
	Host   string `envconfig:"default=https://api.telegram.org"`
	Debug  bool   `envconfig:"default=false"`
}
type Mssql struct {
	Host        string        `envconfig:"default=localhost"`
	Port        string        `envconfig:"default=1433"`
	User        string        `envconfig:"default=sa"`
	Password    string        `envconfig:"default=Password"`
	Database    string        `envconfig:"default=#temp"`
	AppName     string        `envconfig:"default=no name"`
	DialTimeout string        `envconfig:"default=10"`
	KeepAlive   string        `envconfig:"default=1440"`
	PacketSize  string        `envconfig:"default=4096"`
	Log         string        `envconfig:"default=0"`
	ConnTimeout time.Duration `envconfig:"default=10s"`
}
type Service struct {
	Name        string        `envconfig:"default=no name"`
	Env         string        `envconfig:"default=dev"`
	TestEnv     bool          `envconfig:"default=false"`
	GrpcPort    string        `envconfig:"default=50051"`
	ConfigPort  string        `envconfig:"default=50055"`
	Debug       bool          `envconfig:"default=false"`
	GrpcTimeout time.Duration `envconfig:"default=1s"`
}

type Migrate struct {
	User  string `envconfig:"optional"`
	Token string `envconfig:"optional"`
}

type ProxyService struct {
	CollectPeriod        time.Duration `envconfig:"default=60s"`
	DeleteBadProxyPeriod time.Duration `envconfig:"default=300s"`
	DeleteOldStatPeriod  time.Duration `envconfig:"default=300s"`

	CollectUrl string `envconfig:"default=https://www.sslproxies.org/"`

	CollectHttpTimeout time.Duration `envconfig:"default=5s"`
	GrpcTimeout        time.Duration `envconfig:"default=1s"`
	GrpcPort           int           `envconfig:"default=50051"`
	CheckTimeout       time.Duration `envconfig:"default=10s"`
	CheckPeriod        time.Duration `envconfig:"default=10s"`
}

type FortedService struct {
	GrpcPort string `envconfig:"default=50051"`
}

type PinService struct {
	GrpcPort string `envconfig:"default=50051"`
	Debug    bool   `envconfig:"default=false"`
}
type SboService struct {
	GrpcPort string `envconfig:"default=50051"`
	Debug    bool   `envconfig:"default=false"`
}

type Broker struct {
	Url            string        `envconfig:"NATS_URL"`
	AllowReconnect bool          `envconfig:"NATS_ALLOW_RECONNECT"`
	MaxReconnect   int           `envconfig:"NATS_MAX_RECONNECT"`
	ReconnectWait  time.Duration `envconfig:"NATS_RECONNECT_WAIT"`
	Timeout        time.Duration `envconfig:"NATS_TIMEOUT"`
}
type Pg struct {
	Host       string        `envconfig:"PGHOST"`
	Port       string        `envconfig:"PGPORT"`
	User       string        `envconfig:"PGUSER"`
	Password   string        `envconfig:"PGPASSWORD"`
	Database   string        `envconfig:"PGDATABASE"`
	AppName    string        `envconfig:"PGAPPNAME"`
	DisableTLS bool          `envconfig:"PGDISABLETLS"`
	Timeout    time.Duration `envconfig:"PGTIMEOUT" default:"5s"`
}
