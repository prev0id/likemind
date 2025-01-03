package app

import (
	"context"
	"net/http"
	"time"

	"likemind/internal/config"

	"github.com/go-chi/chi/v5"
)

type Handler interface {
	Routes() chi.Router
	Prefix() string
}

type Stopper func()

type App struct {
	server         *http.Server
	router         chi.Router
	handlers       []Handler
	stoppers       []Stopper
	rootCtxCancel  context.CancelFunc
	gracefulPeriod time.Duration
	requestTimeout time.Duration
}

func InitApp(cfg config.App) (*App, context.Context) {
	router := chi.NewRouter()
	server := &http.Server{
		Handler: router,
		Addr:    cfg.Addr,
	}

	rootCtx, cancel := context.WithCancel(context.Background())
	return &App{
		server:         server,
		router:         router,
		rootCtxCancel:  cancel,
		gracefulPeriod: cfg.GracefulPeriod,
		requestTimeout: cfg.RequestTimeout,
	}, rootCtx
}

func (a *App) WithServer(handlers ...Handler) {
	a.handlers = append(a.handlers, handlers...)
}

func (a *App) WithStoppers(stoppers ...Stopper) {
	a.stoppers = append(a.stoppers, stoppers...)
}
