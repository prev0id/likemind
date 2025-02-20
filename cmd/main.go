package main

import (
	"context"

	"likemind/cmd/bootstrap"
	"likemind/internal/app"
	api_handler "likemind/internal/app/handlers/api"
	page_handler "likemind/internal/app/handlers/page"
	static_handler "likemind/internal/app/handlers/static"
	"likemind/internal/app/middleware"
	"likemind/internal/config"
	"likemind/internal/database"
	"likemind/internal/database/adapter/credentials_adapter"
	"likemind/internal/database/adapter/profile_adapter"
	"likemind/internal/database/repo/contact_repo"
	"likemind/internal/database/repo/credentials_repo"
	profile_picture_repo "likemind/internal/database/repo/picture_repo"
	"likemind/internal/database/repo/user_repo"
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

	cfg, err := config.Parse()
	if err != nil {
		log.Fatal().Err(err).Msg("config.Parse")
	}

	log.Info().Interface("config", cfg).Msgf("successfully parsed")

	app, ctx := app.InitApp(cfg.App)

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

	authMiddleware := middleware.NewAuthMiddleware(authService)

	app.WithHandlers(
		page_handler.New(authMiddleware),
		api_handler.New(profileService, authService, authMiddleware),
		static_handler.New(),
	)

	app.WithStoppers(
		authService.Close,
		database.DB.Close,
	)

	if err := app.Run(ctx); err != nil {
		log.Fatal().Err(err)
	}
}
