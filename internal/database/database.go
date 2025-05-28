package database

import (
	"context"

	"likemind/internal/config"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type postgresConn interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
}

var DB *pgxpool.Pool

func InitDB(ctx context.Context, cfg config.DB) error {
	pool, err := pgxpool.New(ctx, cfg.Addr)
	if err != nil {
		return err
	}

	DB = pool
	return nil
}

type SQL interface {
	Build() (string, []any)
}

func RawSQL(query string, args ...any) SQL {
	return rawSQL{
		query: query,
		args:  args,
	}
}

type rawSQL struct {
	query string
	args  []any
}

func (r rawSQL) Build() (string, []any) {
	return r.query, r.args
}

func Exec(ctx context.Context, sql SQL) (int64, error) {
	query, args := sql.Build()

	cmd, err := txOrPool(ctx).Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return cmd.RowsAffected(), nil
}

func Get[T any](ctx context.Context, sql SQL) (T, error) {
	var result T
	query, args := sql.Build()

	err := pgxscan.Get(ctx, txOrPool(ctx), &result, query, args...)

	return result, err
}

func Select[T any](ctx context.Context, sql SQL) ([]T, error) {
	var result []T
	query, args := sql.Build()

	err := pgxscan.Select(ctx, txOrPool(ctx), &result, query, args...)

	return result, err
}

func InTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := DB.Begin(ctx)
	if err != nil {
		log.Err(err).Msg("starting db failed")
	}

	ctx = setTxIntoCtx(ctx, tx)

	if err := fn(ctx); err != nil {
		handleRollback(ctx, tx)
		return err
	}

	return handleCommit(ctx, tx)
}

func handleRollback(ctx context.Context, tx pgx.Tx) {
	if err := tx.Rollback(ctx); err != nil {
		log.Err(err).Msg("transaction rollback failed")
	}
}

func handleCommit(ctx context.Context, tx pgx.Tx) error {
	if err := tx.Commit(ctx); err != nil {
		log.Err(err).Msg("transaction commit failed")
		return err
	}

	return nil
}

type txKeyType struct{}

var txKey = txKeyType{}

func txOrPool(ctx context.Context) postgresConn {
	tx, ok := ctx.Value(txKey).(pgx.Tx)
	if !ok {
		return DB
	}

	return tx
}

func setTxIntoCtx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, txKey, tx)
}
