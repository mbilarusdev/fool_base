package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mbilarusdev/fool_base/log"
)

type DBPool struct {
	InnerPool *pgxpool.Pool
	Limiter   chan struct{}
}

func (pool *DBPool) Acquire(ctx context.Context) (*pgxpool.Conn, error) {
	pool.Limiter <- struct{}{}

	c, err := pool.InnerPool.Acquire(ctx)
	if err != nil {
		log.Err(err, "Postgres connection aquire failed")
		return nil, err
	}

	return c, nil
}

func (pool *DBPool) Release(ctx context.Context, c *pgxpool.Conn) {
	c.Release()
	<-pool.Limiter
}
