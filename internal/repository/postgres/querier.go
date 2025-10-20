package postgres

import (
	"context"
	"database/sql"
)

type SQLQuerier struct {
	db *sql.DB
}

func NewSQLQuerier(db *sql.DB) Querier {
	return &SQLQuerier{db: db}
}

func (q *SQLQuerier) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return q.db.QueryRowContext(ctx, query, args...)
}

func (q *SQLQuerier) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return q.db.QueryContext(ctx, query, args...)
}

func (q *SQLQuerier) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return q.db.ExecContext(ctx, query, args...)
}

type TxQuerier struct {
	tx *sql.Tx
}

func NewTxQuerier(tx *sql.Tx) Querier {
	return &TxQuerier{tx: tx}
}

func (q *TxQuerier) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return q.tx.QueryRowContext(ctx, query, args...)
}

func (q *TxQuerier) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return q.tx.QueryContext(ctx, query, args...)
}

func (q *TxQuerier) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return q.tx.ExecContext(ctx, query, args...)
}
