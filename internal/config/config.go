package config

import (
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
)

type Config struct {
	DB  DB  `toml:"db"`
	App App `toml:"app"`
}

type App struct {
	Addr           string        `toml:"addr"`
	GracefulPeriod time.Duration `toml:"graceful_period"`
	RequestTimeout time.Duration `toml:"request_timeout"`
}

type DB struct {
	Addr string `toml:"addr"`
}

func Parse() (Config, error) {
	cfg := Config{}

	_, err := toml.DecodeFile(*configPath, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("toml.DecodeFile: %w", err)
	}

	return cfg, nil
}
