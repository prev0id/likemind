package bootstrap

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"

	"likemind/internal/api"
	"likemind/internal/config"
	"likemind/internal/domain"
	desc "likemind/internal/pkg/api"
	"likemind/internal/service/image"
	"likemind/internal/service/interests"
	"likemind/internal/service/profile"
	"likemind/internal/service/session"
	"likemind/website"
)

func Server(
	cfg config.App,
	sessionService session.Service,
	profileService profile.Service,
	imageService image.Service,
	interests interests.Service,
) error {
	server := api.NewServer(sessionService, profileService, imageService, interests)
	security := api.NewSecurityHandler(sessionService)

	router := chi.NewRouter()

	router.Use(
		middleware.RequestID,
		middleware.Logger,
		middleware.Heartbeat(domain.PathAPIPing),
		middleware.StripSlashes,
		middleware.Timeout(cfg.RequestTimeout),
		middleware.Recoverer, // should be last
	)

	app, err := desc.NewServer(
		server,
		security,
		desc.WithNotFound(server.NotFound),
	)
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
