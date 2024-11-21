package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	DB  DB
	API API
}

type API struct {
	Addr string
}

type DB struct {
	Network  string
	Insecure bool
	Addr     string
	User     string
	Password string
	Database string
}

func Parse() (Config, error) {
	cfg := Config{}

	_, err := toml.DecodeFile(*configPath, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("toml.DecodeFile: %w", err)
	}

	return cfg, nil
}
