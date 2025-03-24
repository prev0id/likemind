package main

import (
	"context"

	"likemind/cmd/bootstrap"
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

	if err := bootstrap.Server(cfg.App, authService, profileService); err != nil {
		log.Fatal().Err(err).Msg("bootstrap.Server")
	}
}
