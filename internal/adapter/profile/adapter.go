package profile_adapter

import (
	"context"
	"database/sql"
	"fmt"

	"likemind/internal/adapter/model"
	"likemind/internal/config"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Adapter struct {
	db *bun.DB
}

func New(cfg config.DB) (*Adapter, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithNetwork(cfg.Network),
		pgdriver.WithInsecure(cfg.Insecure),
		pgdriver.WithAddr(cfg.Addr),
		pgdriver.WithUser(cfg.User),
		pgdriver.WithPassword(cfg.Password),
		pgdriver.WithDatabase(cfg.Database),
	))

	db := bun.NewDB(sqldb, pgdialect.New())

	return &Adapter{
		db: db,
	}, nil
}

func (a *Adapter) CreateUser(ctx context.Context, user *model.User) (int64, error) {
	result, err := a.db.NewInsert().
		Model(user).
		Exec(ctx)
	if err != nil {
		return 0, fmt.Errorf("inserting query error: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("result.LastInsertId: %w", err)
	}

	return id, nil
}

func (a *Adapter) GetUser(ctx context.Context, id int64) (*model.User, error) {
	result := &model.User{}

	err := a.db.NewSelect().
		Model(result).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("select user error: %w", err)
	}

	return result, nil
}
