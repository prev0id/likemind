package main

import (
	"log"
	"runtime"

	"likemind/internal/api"
	"likemind/internal/config"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		log.Fatalf("cofnig.Parse: %s", err.Error())
	}

	if !cfg.DB.Insecure {
		runtime.Breakpoint()
	}

	log.Printf("cfg: %+v", cfg)

	if err := api.BootstrapServer(cfg.API); err != nil {
		log.Fatalf("api.BootstrapServer: %s", err.Error())
	}
}
