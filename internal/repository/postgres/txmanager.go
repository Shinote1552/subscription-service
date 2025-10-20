package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type keyTxType int

const keyTxValue keyTxType = iota

var ErrNoTransaction = errors.New("no transaction in context")

// SQLTxManager реализует TxManager интерфейс
type SQLTxManager struct {
	db *sql.DB
}

// NewSQLTxManager возвращает TxManager (интерфейс), а не конкретную реализацию
func NewSQLTxManager(db *sql.DB) TxManager {
	return &SQLTxManager{
		db: db,
	}
}

func (tm *SQLTxManager) WithTx(ctx context.Context, opts *sql.TxOptions, fn func(ctx context.Context) error) error {
	if opts == nil {
		opts = &sql.TxOptions{
			Isolation: sql.LevelSerializable,
			ReadOnly:  false,
		}
	}

	tx, err := tm.db.BeginTx(ctx, opts)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	ctx = context.WithValue(ctx, keyTxValue, tx)

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	if err := fn(ctx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (tm *SQLTxManager) GetQuerier(ctx context.Context) (Querier, error) {
	tx, err := tm.getTx(ctx)
	if err != nil {
		if errors.Is(err, ErrNoTransaction) {
			return NewSQLQuerier(tm.db), nil
		}
		return nil, fmt.Errorf("get transaction: %w", err)
	}
	return NewTxQuerier(tx), nil
}

func (tm *SQLTxManager) getTx(ctx context.Context) (*sql.Tx, error) {
	tx, ok := ctx.Value(keyTxValue).(*sql.Tx)
	if !ok {
		return nil, ErrNoTransaction
	}
	return tx, nil
}
