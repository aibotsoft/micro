package config

import (
	"time"
)

// Config provides the system configuration.
type Config struct {
	Service Service
	Mssql   Mssql
	//Broker       Broker
	//Pg     Pg
	ProxyService ProxyService
	Ristretto    Ristretto
	//Logging  Logging
	Migrate        Migrate
	SurebetService SurebetService
	PinService     PinService
	SboService     SboService
	FortedService  FortedService
}
type Mssql struct {
	Host        string `envconfig:"default=localhost"`
	Port        string `envconfig:"default=1433"`
	User        string `envconfig:"default=sa"`
	Password    string `envconfig:"default=Password"`
	Database    string `envconfig:"default=#temp"`
	AppName     string `envconfig:"default=no name"`
	DialTimeout string `envconfig:"default=10"`
	KeepAlive   string `envconfig:"default=1440"`
	PacketSize  string `envconfig:"default=4096"`
	Log         string `envconfig:"default=0"`

	ConnTimeout time.Duration `envconfig:"default=10s"`
}
type Service struct {
	Name    string `envconfig:"default=no name"`
	Env     string `envconfig:"default=dev"`
	TestEnv bool   `envconfig:"default=false"`
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

type SurebetService struct {
	GrpcTimeout time.Duration `envconfig:"default=1s"`
	GrpcPort    string        `envconfig:"default=50051"`
}

type FortedService struct {
	GrpcPort string `envconfig:"default=50051"`
}

type PinService struct {
	GrpcPort string `envconfig:"default=50051"`
	User     string `envconfig:"optional"`
	Pass     string `envconfig:"optional"`
	Debug    bool   `envconfig:"default=false"`
}
type SboService struct {
	GrpcPort string `envconfig:"default=50051"`
	User     string `envconfig:"optional"`
	Pass     string `envconfig:"optional"`
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

//type (
//
//	//// Web provides api server configuration.
//	//Web struct {
//	//	APIHost      string
//	//	ReadTimeout  time.Duration
//	//	WriteTimeout time.Duration
//	//}
//	//// Logging provides the logging configuration.
//	//Logging struct {
//	//	Level  string `envconfig:"LOG_LEVEL"`
//	//	Trace  bool   `envconfig:"DRONE_LOGS_TRACE"`
//	//	Color  bool   `envconfig:"DRONE_LOGS_COLOR"`
//	//	Pretty bool   `envconfig:"DRONE_LOGS_PRETTY"`
//	//	Text   bool   `envconfig:"DRONE_LOGS_TEXT"`
//	//}
//	//// Database provides the database configuration.
//)
//type (
//	// Config provides the system configuration.
//	Config struct {
//		Database Database
//		Logging  Logging
//		Web      Web
//	}
//	// Web provides api server configuration.
//	Web struct {
//		APIHost      string
//		ReadTimeout  time.Duration
//		WriteTimeout time.Duration
//	}
//	// Logging provides the logging configuration.
//	Logging struct {
//		Level  string `envconfig:"LOG_LEVEL"`
//		Trace  bool   `envconfig:"DRONE_LOGS_TRACE"`
//		Color  bool   `envconfig:"DRONE_LOGS_COLOR"`
//		Pretty bool   `envconfig:"DRONE_LOGS_PRETTY"`
//		Text   bool   `envconfig:"DRONE_LOGS_TEXT"`
//	}
//	// Database provides the database configuration.

//)
