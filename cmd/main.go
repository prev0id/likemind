package main

import (
	"log"

	"likemind/cmd/bootstrap"
	"likemind/internal/app"
	page_handler "likemind/internal/app/handlers/page"
	static_handler "likemind/internal/app/handlers/static"
	"likemind/internal/config"
	"likemind/internal/data_provider"
	"likemind/internal/domain"
	"likemind/internal/service/auth"
	"likemind/internal/service/profile"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		log.Fatalf("cofnig.Parse: %s", err.Error())
	}

	log.Printf("Config: %+v", cfg)

	app, ctx := app.InitApp(cfg.App)

	dbConn, err := bootstrap.DB(ctx, cfg.DB)

	userProvider := data_provider.New[domain.User](dbConn, domain.TableUser)
	credProvider := data_provider.New[domain.Credential](dbConn, domain.TableCredential)
	// groupProvider := data_provider.New[domain.Group](dbConn, domain.TableGroup)
	// interestProvider := data_provider.New[domain.Interest](dbConn, domain.TableInterest)
	// groupInterestProvider := data_provider.New[domain.AppliedInterest](dbConn, domain.TableGroupInterest)
	// userInterestProvider := data_provider.New[domain.AppliedInterest](dbConn, domain.TableUserInterest)

	profile.New(userProvider)

	authSvc, err := auth.New(credProvider, dbConn, cfg.Auth)
	if err != nil {
		log.Fatalf("auth.New: %s", err.Error())
	}

	app.WithHandlers(
		page_handler.New(),
		static_handler.New(),
	)

	app.WithStoppers(
		authSvc.Close,
		dbConn.Close,
	)

	if err := app.Run(ctx); err != nil {
		log.Fatalf("app.Run: %s", err.Error())
	}
}
