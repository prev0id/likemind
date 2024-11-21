package api

import (
	"log"
	"net/http"

	dev_handlers "likemind/internal/api/handlers/dev_handler"
	profile_handlers "likemind/internal/api/handlers/profile_handler"
	"likemind/internal/api/middlware"
	"likemind/internal/config"
)

func BootstrapServer(cfg config.API) error {
	bootstrapProfileGroup()
	bootstrapStatic()

	log.Printf("starting app at %s", cfg.Addr)

	return http.ListenAndServe(cfg.Addr, nil)
}

func bootstrapProfileGroup() {
	group := middlware.NewGroup(
		"/profile/",
		middlware.Recover,
		middlware.Logging,
		middlware.Auth,
	)

	srv := profile_handlers.New(nil)

	group.Register(http.MethodGet, "{username}", http.HandlerFunc(srv.ProfilePage))
	group.Register(http.MethodPost, "{username}/update_name", http.HandlerFunc(srv.ChangeName))
}

func bootstrapDev() {
	group := middlware.NewGroup(
		"/dev/",
		middlware.Recover,
		middlware.Logging,
		middlware.Auth,
	)

	srv := dev_handlers.New()

	group.Register(http.MethodGet, "/", http.HandlerFunc(srv.Page))
	group.Register(http.MethodGet, "/mock/{widget}", http.HandlerFunc(srv.MockWidget))
	group.Register(http.MethodGet, "/mock/{widget}/list", http.HandlerFunc(srv.ListMocks))
}

func bootstrapStatic() {
	group := middlware.NewGroup(
		"/static/",
		middlware.Recover,
		middlware.Logging,
	)

	group.Register(http.MethodGet, "/", http.StripPrefix("/static/", http.FileServer(http.Dir("website/static"))))
}
