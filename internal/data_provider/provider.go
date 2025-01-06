package data_provider

import (
	"context"
	"fmt"

	"likemind/internal/domain"

	"github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dialect  = "postgres"
	columnID = "id"
)

// type User struct {
// 	ID       int64    `db:"id" goqu:"skipupdate"`
// 	Name     string   `db:"name" goqu:"omitempty"`
// 	Surname  string   `db:"surname" goqu:"omitempty"`
// 	Nickname string   `db:"nickname" goqu:"omitempty"`
// 	About    string   `db:"about" goqu:"omitempty"`
// 	PfpID    string   `db:"pft_id" goqu:"omitempty"`
// 	Contacts []string `db:"contacts" goqu:"omitempty"`
// }

type provider[T domain.Data] struct {
	conn  *pgxpool.Pool
	table string
}

func New[T domain.Data](conn *pgxpool.Pool, table string) domain.DataProvider[T] {
	return &provider[T]{
		conn:  conn,
		table: table,
	}
}

func (p *provider[T]) List(ctx context.Context) ([]T, error) {
	var result []T

	query, _, _ := goqu.Dialect(dialect).
		Select(new(T)).
		From(p.table).
		ToSQL()

	if err := pgxscan.Select(ctx, p.conn, &result, query); err != nil {
		return nil, fmt.Errorf("pgxscan.Select: %w", err)
	}

	return result, nil
}

func (p *provider[T]) Get(ctx context.Context, id int64) (T, error) {
	var result T

	query, _, _ := goqu.Dialect(dialect).
		Select(result).
		From(p.table).
		Where(goqu.C(columnID).Eq(id)).
		ToSQL()

	if err := pgxscan.Get(ctx, p.conn, &result, query); err != nil {
		return result, fmt.Errorf("pgxscan.Get: %w", err)
	}

	return result, nil
}

func (p *provider[T]) Update(ctx context.Context, data T) error {
	query, _, _ := goqu.Dialect(dialect).
		Update(p.table).
		Set(data).
		Where(goqu.C(columnID).Eq(data.GetID())).
		ToSQL()

	if _, err := p.conn.Exec(ctx, query); err != nil {
		return fmt.Errorf("db.conn.Exec: %w", err)
	}

	return nil
}

func (p *provider[T]) UpdateField(ctx context.Context, id int64, field string, value any) error {
	query, _, _ := goqu.Dialect(dialect).
		Update(p.table).
		Set(goqu.Record{
			field: value,
		}).
		Where(goqu.C(columnID).Eq(id)).
		ToSQL()

	if _, err := p.conn.Exec(ctx, query); err != nil {
		return fmt.Errorf("db.conn.Exec: %w", err)
	}

	return nil
}

func (p *provider[T]) Insert(ctx context.Context, data T) (int64, error) {
	query, _, _ := goqu.Dialect(dialect).
		Insert(p.table).
		Rows(data).
		Returning(columnID).
		ToSQL()

	var result int64
	if err := pgxscan.Get(ctx, p.conn, &result, query); err != nil {
		return 0, fmt.Errorf("pgxscan.Get: %w", err)
	}

	return result, nil
}

func (p *provider[T]) Delete(ctx context.Context, id int64) error {
	query, _, _ := goqu.Dialect(dialect).
		Delete(p.table).
		Where(goqu.C(columnID).Eq(id)).
		ToSQL()

	if _, err := p.conn.Exec(ctx, query); err != nil {
		return fmt.Errorf("db.conn.Exec: %w", err)
	}

	return nil
}

// func (p *provider[T]) GetWhere(ctx context.Context, where string, args ...any) (T, error) {
// 	var result T

// 	query, _, _ := goqu.Dialect(dialect).
// 		Select(result).
// 		From(p.table).
// 		Where(goqu.L(where, args...)).
// 		ToSQL()

// 	if err := pgxscan.Get(ctx, p.conn, &result, query); err != nil {
// 		return result, fmt.Errorf("pgxscan.Get: %w", err)
// 	}

// 	return result, nil
// }

// func userFromDomain(user domain.User) User {
// 	return User{
// 		ID:       user.ID,
// 		Name:     user.Name,
// 		Surname:  user.Surname,
// 		Nickname: user.Nickname,
// 		PfpID:    user.PfpID,
// 		Contacts: user.Contacts,
// 		About:    user.About,
// 	}
// }

// func usersFromPosgres(users []User) []domain.User {
// 	return lo.Map(users, func(user User, _ int) domain.User {
// 		return userFromPosgres(user)
// 	})
// }

// func userFromPosgres(user User) domain.User {
// 	return domain.User{
// 		ID:       user.ID,
// 		Name:     user.Name,
// 		Surname:  user.Surname,
// 		Nickname: user.Nickname,
// 		PfpID:    user.PfpID,
// 		Contacts: user.Contacts,
// 		About:    user.About,
// 	}
// }
