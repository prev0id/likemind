package main

import (
	"likemind/cmd/bootstrap"
	"likemind/internal/app"
	api_handler "likemind/internal/app/handlers/api"
	page_handler "likemind/internal/app/handlers/page"
	static_handler "likemind/internal/app/handlers/static"
	"likemind/internal/config"
	"likemind/internal/data_provider"
	"likemind/internal/domain"
	"likemind/internal/service/auth"
	"likemind/internal/service/profile"

	"github.com/rs/zerolog/log"
)

func main() {
	bootstrap.Deps()

	cfg, err := config.Parse()
	if err != nil {
		log.Fatal().Err(err).Msg("config.Parse")
	}

	log.Info().Interface("config", cfg).Msgf("successfully parsed")

	app, ctx := app.InitApp(cfg.App)

	dbConn, err := bootstrap.DB(ctx, cfg.DB)

	userProvider := data_provider.New[domain.User](dbConn, domain.TableUser)
	credProvider := data_provider.New[domain.Credential](dbConn, domain.TableCredential)
	// groupProvider := data_provider.New[domain.Group](dbConn, domain.TableGroup)
	// interestProvider := data_provider.New[domain.Interest](dbConn, domain.TableInterest)
	// groupInterestProvider := data_provider.New[domain.AppliedInterest](dbConn, domain.TableGroupInterest)
	// userInterestProvider := data_provider.New[domain.AppliedInterest](dbConn, domain.TableUserInterest)

	profileSvc := profile.New(userProvider)

	authSvc, err := auth.New(credProvider, dbConn, cfg.Auth)
	if err != nil {
		log.Fatal().Err(err)
	}

	app.WithHandlers(
		page_handler.New(authSvc),
		api_handler.New(authSvc, profileSvc),
		static_handler.New(),
	)

	app.WithStoppers(
		authSvc.Close,
		dbConn.Close,
	)

	if err := app.Run(ctx); err != nil {
		log.Fatal().Err(err)
	}
}
