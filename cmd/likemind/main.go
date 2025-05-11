package main

import (
	"context"

	"github.com/rs/zerolog/log"

	"likemind/cmd/likemind/bootstrap"
	"likemind/internal/config"
	"likemind/internal/database"
	"likemind/internal/database/adapter/interest_adapter"
	"likemind/internal/database/adapter/profile_adapter"
	"likemind/internal/database/adapter/session_adapter"
	"likemind/internal/database/repo/contact_repo"
	"likemind/internal/database/repo/interest_repo"
	profile_picture_repo "likemind/internal/database/repo/picture_repo"
	"likemind/internal/database/repo/session_repo"
	"likemind/internal/database/repo/user_repo"
	s3_image_repo "likemind/internal/s3/image_repo"
	"likemind/internal/service/image"
	"likemind/internal/service/interests"
	"likemind/internal/service/profile"
	"likemind/internal/service/session"
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
	interestsRepo := &interest_repo.Repo{}

	s3, err := s3_image_repo.NewS3Repository(cfg.S3)
	if err != nil {
		log.Fatal().Err(err).Msg("s3_image_repo.NewS3Repository")
	}

	sessionAdapter := session_adapter.NewAdapter(sessionRepo)
	profileAdapter := profile_adapter.NewAdapter(userRepo, contactRepo, profilePictureRepo)
	interestsAdapter := interest_adapter.New(interestsRepo)

	profileService := profile.New(profileAdapter)
	sessionService := session.New(sessionAdapter, cfg.Auth)
	imageService := image.NewImageService(s3, profileAdapter)
	interests := interests.New(interestsAdapter)

	if err := bootstrap.Server(cfg.App, sessionService, profileService, imageService, interests); err != nil {
		log.Fatal().Err(err).Msg("bootstrap.Server")
	}
}
