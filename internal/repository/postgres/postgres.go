package postgres

import (
	"context"
	"database/sql"
)

type PostgresStorage struct {
	db  *sql.DB
	txm TxManager
}

type TxManager interface {
	WithTx(ctx context.Context, opts *sql.TxOptions, fn func(ctx context.Context) error) error
	GetQuerier(ctx context.Context) (Querier, error)
}

type Querier interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

func NewPostgresStorage(db *sql.DB, txm TxManager) *PostgresStorage {
	return &PostgresStorage{
		db:  db,
		txm: txm,
	}
}
