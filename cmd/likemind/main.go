package main

import (
	"context"

	"likemind/cmd/likemind/bootstrap"
	"likemind/internal/config"
	"likemind/internal/database"
	"likemind/internal/database/adapter/profile_adapter"
	"likemind/internal/database/adapter/session_adapter"
	"likemind/internal/database/repo/contact_repo"
	profile_picture_repo "likemind/internal/database/repo/picture_repo"
	"likemind/internal/database/repo/session_repo"
	"likemind/internal/database/repo/user_repo"
	"likemind/internal/service/profile"
	"likemind/internal/service/session"

	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()

	bootstrap.Libs()

	cfg, err := config.Parse()
	if err != nil {
		log.Fatal().Err(err).Msg("config.Parse")
	}

	log.Info().Interface("config", cfg).Msgf("successfully parsed")

	database.InitDB(ctx, cfg.DB)

	sessionRepo := &session_repo.Repo{}
	userRepo := &user_repo.Repo{}
	contactRepo := &contact_repo.Repo{}
	profilePictureRepo := &profile_picture_repo.Repo{}

	sessionAdapter := session_adapter.NewAdapter(sessionRepo)
	profileAdapter := profile_adapter.NewAdapter(userRepo, contactRepo, profilePictureRepo)

	profileService := profile.New(profileAdapter)
	sessionService := session.New(sessionAdapter, cfg.Auth)

	if err := bootstrap.Server(cfg.App, sessionService, profileService); err != nil {
		log.Fatal().Err(err).Msg("bootstrap.Server")
	}
}
