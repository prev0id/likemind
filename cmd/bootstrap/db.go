package bootstrap

import (
	"context"

	"likemind/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DB(ctx context.Context, cfg config.DB) (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(ctx, cfg.Addr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
