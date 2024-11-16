package config

import (
	"flag"
)

var configPath = flag.String("config", "./config.toml", "path to application config")

func init() {
	flag.Parse()
}
