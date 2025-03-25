package bootstrap

import (
	"fmt"
	"net/http"

	"likemind/internal/api"
	"likemind/internal/config"
	"likemind/internal/domain"
	desc "likemind/internal/pkg/api"
	"likemind/internal/service/auth"
	"likemind/internal/service/profile"
	"likemind/website"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

func Server(cfg config.App, authService auth.Service, profileService profile.Service) error {
	server := api.NewServer(authService, profileService)

	router := chi.NewRouter()

	router.Use(
		middleware.RequestID,
		middleware.Logger,
		middleware.Heartbeat(domain.PathAPIPing),
		middleware.StripSlashes,
		middleware.Timeout(cfg.RequestTimeout),
		middleware.Recoverer, // should be last
	)

	app, err := desc.NewServer(server, desc.WithNotFound(server.NotFound))
	if err != nil {
		return fmt.Errorf("desc.NewServer: %w", err)
	}

	router.Mount("/v1/", app)
	router.Handle("/static/*", http.FileServer(http.FS(website.StaticFiles)))

	// TODO: metrics

	log.Warn().Msgf("starting server at %q", cfg.Addr)

	if err := http.ListenAndServe(cfg.Addr, router); err != nil {
		return fmt.Errorf("http.ListenAndServe: %w", err)
	}

	return nil
}
