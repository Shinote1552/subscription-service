package storage

import (
	"context"

	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	pool   *pgxpool.Pool
	trm    trm.Manager
	getter *trmpgx.CtxGetter
}

func NewStorage(pool *pgxpool.Pool, trm trm.Manager, getter *trmpgx.CtxGetter) *Storage {
	return &Storage{
		pool:   pool,
		trm:    trm,
		getter: getter,
	}
}

func (s *Storage) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return s.trm.Do(ctx, fn)
}

func (s *Storage) GetQueryable(ctx context.Context) interface{} {
	return s.getter.DefaultTrOrDB(ctx, s.pool)
}
