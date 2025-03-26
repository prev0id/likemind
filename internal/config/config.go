package config

import (
	"flag"
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
)

type Config struct {
	DB   DB   `toml:"db"`
	App  App  `toml:"app"`
	Auth Auth `toml:"auth"`
}

type App struct {
	Addr           string        `toml:"addr"`
	RequestTimeout time.Duration `toml:"request_timeout"`
}

type DB struct {
	Addr string `toml:"addr"`
}

type Auth struct {
	Expiration time.Duration `toml:"expiration"`
}

var configPath = flag.String("config", "./config.toml", "path to application config")

func Parse() (Config, error) {
	cfg := Config{}

	_, err := toml.DecodeFile(*configPath, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("toml.DecodeFile: %w", err)
	}

	return cfg, nil
}
