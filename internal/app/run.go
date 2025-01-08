package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"likemind/internal/domain"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

func (a *App) Run(ctx context.Context) error {
	a.registerAPI()

	a.setUpGracefulShoutdown(ctx)

	// file, _ := os.Create("./api.md")
	// reader := strings.NewReader(docgen.MarkdownRoutesDoc(a.router, docgen.MarkdownOpts{
	// 	ProjectPath: "./",
	// 	Intro:       "Welcome to the chi/_examples/rest generated docs.",
	// }))
	// io.Copy(file, reader)

	log.Info().
		Str("address", a.server.Addr).
		Msg("starting app")

	if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	<-ctx.Done()
	return nil
}

func (a *App) registerAPI() {
	a.router.Use(middleware.RequestID)
	a.router.Use(middleware.Logger)
	a.router.Use(middleware.Heartbeat(domain.PathPing))
	a.router.Use(middleware.StripSlashes)
	a.router.Use(middleware.Timeout(a.requestTimeout))
	a.router.Use(middleware.Recoverer) // should be last

	for _, handler := range a.handlers {
		a.router.Mount(handler.Prefix(), handler.Routes())
	}
}

func (a *App) setUpGracefulShoutdown(ctx context.Context) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig // waiting for termination signal
		log.Info().Msg("Recieved termination signal")

		shutdownCtx, shutdownCancel := context.WithTimeout(ctx, a.gracefulPeriod)
		go a.forceExitOnDeadline(shutdownCtx)

		a.stop(shutdownCtx)
		shutdownCancel()
		log.Info().Msg("Gracefully stopped")
	}()
}

func (a *App) forceExitOnDeadline(ctx context.Context) {
	<-ctx.Done()
	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		log.Fatal().Msg("Graceful shutdown timed out. Forcing exit.")
	}
}

func (a *App) stop(ctx context.Context) {
	for _, stopper := range a.stoppers {
		stopper()
	}
	a.server.Shutdown(ctx)
	a.rootCtxCancel()
}
