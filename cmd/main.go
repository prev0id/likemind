package main

import (
	"log"

	"likemind/internal/api"
)

func main() {
	if err := api.BootstrapServer(); err != nil {
		log.Fatalf("http.ListenAndServe: %v", err)
	}
}
