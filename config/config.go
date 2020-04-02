package config

import (
	"github.com/vrischmann/envconfig"
	"log"
	"os"
)

const envFileName = ".env"
const devEnv = "dev"

// New returns the settings from the environment.
func New() *Config {
	cfg := &Config{}
	err := envconfig.Init(cfg)
	if err != nil {
		log.Print(err)
	}
	return cfg
}

func IsDev() bool {
	return os.Getenv("SERVICE_ENV") == devEnv
}
