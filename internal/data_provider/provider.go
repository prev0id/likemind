package data_provider

import (
	"context"
	"fmt"
	"log"

	"likemind/internal/domain"

	"github.com/georgysavva/scany/v2/pgxscan"
	sql "github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dialect  = "postgres"
	columnID = "id"
)

type provider[T domain.Data] struct {
	conn       *pgxpool.Pool
	table      string
	dataStruct *sql.Struct
}

func New[T domain.Data](conn *pgxpool.Pool, table string) domain.DataProvider[T] {
	return &provider[T]{
		conn:       conn,
		table:      table,
		dataStruct: sql.NewStruct(new(T)).For(sql.PostgreSQL),
	}
}

func (p *provider[T]) List(ctx context.Context) ([]T, error) {
	var result []T

	query, args := p.dataStruct.SelectFrom(p.table).Build()

	if err := pgxscan.Select(ctx, p.conn, &result, query, args); err != nil {
		return nil, fmt.Errorf("pgxscan.Select: %w", err)
	}

	return result, nil
}

func (p *provider[T]) Get(ctx context.Context, field string, value any) (T, error) {
	var result T

	builder := p.dataStruct.SelectFrom(p.table)
	builder.Where(builder.Equal(field, value))
	query, args := builder.Build()

	if err := pgxscan.Get(ctx, p.conn, &result, query, args); err != nil {
		return result, fmt.Errorf("pgxscan.Get: %w", err)
	}

	return result, nil
}

func (p *provider[T]) Update(ctx context.Context, data T) error {
	builder := p.dataStruct.Update(p.table, data)
	builder.Where(builder.Equal(columnID, data.GetID()))
	query, args := builder.Build()

	if _, err := p.conn.Exec(ctx, query, args); err != nil {
		return fmt.Errorf("db.conn.Exec: %w", err)
	}

	return nil
}

func (p *provider[T]) UpdateField(ctx context.Context, id int64, field string, value any) error {
	builder := sql.PostgreSQL.NewUpdateBuilder()
	builder.Update(p.table)
	builder.Set(builder.Assign(field, value))
	builder.Where(builder.Equal(columnID, id))
	query, args := builder.Build()

	if _, err := p.conn.Exec(ctx, query, args); err != nil {
		return fmt.Errorf("db.conn.Exec: %w", err)
	}

	return nil
}

func (p *provider[T]) Insert(ctx context.Context, data T) (int64, error) {
	builder := p.dataStruct.InsertInto(p.table, data)
	builder.SQL("RETURNING " + columnID)
	query, args := builder.Build()

	log.Println(query, args)

	var result int64
	if err := pgxscan.Get(ctx, p.conn, &result, query, args); err != nil {
		return 0, fmt.Errorf("pgxscan.Get: %w", err)
	}

	return result, nil
}

func (p *provider[T]) Delete(ctx context.Context, id int64) error {
	builder := p.dataStruct.DeleteFrom(p.table)
	builder.Where(builder.Equal(columnID, id))
	query, args := builder.Build()

	if _, err := p.conn.Exec(ctx, query, args); err != nil {
		return fmt.Errorf("db.conn.Exec: %w", err)
	}

	return nil
}
