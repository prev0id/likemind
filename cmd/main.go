package main

import (
	"context"
	"fmt"
	"log"

	profile_adapter "likemind/internal/adapter/profile"
	"likemind/internal/config"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		log.Fatalf("cofnig.Parse: %s", err.Error())
	}

	log.Printf("cfg: %+v", cfg)

	adapter, err := profile_adapter.New(cfg.DB)
	if err != nil {
		log.Fatalf("profile_adapter.New: %s", err.Error())
	}

	user, err := adapter.GetUser(context.Background(), 1)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	fmt.Println(user)
}
