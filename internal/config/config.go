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
	S3   S3   `toml:"s3"`
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

type S3 struct {
	Endpoint        string `toml:"endpoint"`
	AccessKeyID     string `toml:"access_key_id"`
	SecretAccessKey string `toml:"secret_access_key"`
	BucketName      string `toml:"bucket_name"`
	Location        string `toml:"location"`
	UseSSL          bool   `toml:"use_ssl"`
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
