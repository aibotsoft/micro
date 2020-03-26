package config

import (
	"time"
)

// Config provides the system configuration.
type Config struct {
	Service        Service
	Broker         Broker
	Postgres       Postgres
	CollectService ProxyService
	//Logging  Logging
	//Web      Web
}

type Service struct {
	Name        string
	Env         string
	TestLoadEnv bool `envconfig:"TEST_LOAD_ENV" default:"true"`
	//Logging  Logging
	//Web      Web
}

type ProxyService struct {
	CollectPeriod time.Duration `default:"60s"`
	CollectUrl    string        `default:"https://www.sslproxies.org/"`

	HttpTimeout time.Duration `default:"5s"`
	GRPCTimeout time.Duration `default:"1s"`
	GRPCPort    int           `default:"50051"`
}

type Broker struct {
	Url            string        `envconfig:"NATS_URL"`
	AllowReconnect bool          `envconfig:"NATS_ALLOW_RECONNECT"`
	MaxReconnect   int           `envconfig:"NATS_MAX_RECONNECT"`
	ReconnectWait  time.Duration `envconfig:"NATS_RECONNECT_WAIT"`
	Timeout        time.Duration `envconfig:"NATS_TIMEOUT"`
}
type Postgres struct {
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
