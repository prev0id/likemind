package bootstrap

import (
	"context"
	"log"

	"likemind/internal/app"
	"likemind/internal/config"

	"github.com/jackc/pgx/v5"
)

func DB(ctx context.Context, cfg config.DB) (*pgx.Conn, app.Stopper, error) {
	conn, err := pgx.Connect(ctx, cfg.Addr)
	if err != nil {
		return nil, nil, err
	}

	return conn, wrapPGXConnClose(ctx, conn), nil
}

func wrapPGXConnClose(ctx context.Context, conn *pgx.Conn) app.Stopper {
	return func() {
		if err := conn.Close(ctx); err != nil {
			log.Printf("pgx close err: %s", err.Error())
		}
	}
}
