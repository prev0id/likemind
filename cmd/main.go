package main

import (
	"context"
	"net/http"

	"likemind/cmd/bootstrap"
	"likemind/internal/api"
	"likemind/internal/config"
	"likemind/internal/database"
	"likemind/internal/database/adapter/credentials_adapter"
	"likemind/internal/database/adapter/profile_adapter"
	"likemind/internal/database/repo/contact_repo"
	"likemind/internal/database/repo/credentials_repo"
	profile_picture_repo "likemind/internal/database/repo/picture_repo"
	"likemind/internal/database/repo/user_repo"
	desc "likemind/internal/pkg/api"
	"likemind/internal/service/auth"
	"likemind/internal/service/profile"

	"github.com/rs/zerolog/log"
)

type DB interface{}

type (
	ListOption   interface{}
	GetOption    interface{}
	InsertOption interface{}
	DeleteOption interface{}
)

type DataProvider[Data any] interface {
	Get(ctx context.Context, db DB, opts ...GetOption) (Data, error)
	List(ctx context.Context, db DB, opts ...ListOption) ([]Data, error)
	Insert(ctx context.Context, db DB, opts ...InsertOption) error
	DeleteByField(ctx context.Context, opts ...DeleteOption) error
}

func main() {
	bootstrap.Deps()

	ctx := context.Background()

	cfg, err := config.Parse()
	if err != nil {
		log.Fatal().Err(err).Msg("config.Parse")
	}

	log.Info().Interface("config", cfg).Msgf("successfully parsed")

	database.InitDB(ctx, cfg.DB)

	credsRepo := &credentials_repo.Repo{}
	userRepo := &user_repo.Repo{}
	contactRepo := &contact_repo.Repo{}
	profilePictureRepo := &profile_picture_repo.Repo{}

	credsAdapter := credentials_adapter.NewAdapter(credsRepo)
	profileAdapter := profile_adapter.NewAdapter(userRepo, contactRepo, profilePictureRepo)

	profileService := profile.New(profileAdapter)
	authService, err := auth.New(credsAdapter, cfg.Auth)
	if err != nil {
		log.Fatal().Err(err).Msg("auth.New")
	}

	server := api.NewServer(authService, profileService)

	app, err := desc.NewServer(server, desc.WithNotFound(server.NotFound))
	if err != nil {
		log.Fatal().Err(err).Msg("desc.NewServer")
	}

	if err := http.ListenAndServe(":8080", app); err != nil {
		log.Fatal().Err(err).Msg("http.ListenAndServe")
	}
}
