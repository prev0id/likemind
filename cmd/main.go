package main

import (
	"log"

	"likemind/internal/app"
	page_handler "likemind/internal/app/handlers/page"
	static_handler "likemind/internal/app/handlers/static"
	"likemind/internal/config"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		log.Fatalf("cofnig.Parse: %s", err.Error())
	}

	log.Printf("Config: %+v", cfg)

	app, ctx := app.InitApp(cfg.App)

	// dbConn, dbStopper, err := bootstrap.DB(ctx, cfg.DB)

	pageHandler := page_handler.New()
	staticHandler := static_handler.New()

	app.WithServer(
		pageHandler,
		staticHandler,
	)

	app.WithStoppers(
	// dbStopper,
	)

	if err := app.Run(ctx); err != nil {
		log.Fatalf("app.Run: %s", err.Error())
	}
}
