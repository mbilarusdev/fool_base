package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mbilarusdev/fool_base/log"
)

func ConnPGX(serviceName string) *DBPool {
	url := os.Getenv("PGX")

	conf, err := pgxpool.ParseConfig(url)
	if err != nil {
		log.Err(err, "Postgres config parsing failed!")
		panic(err)
	}

	outerPool := new(DBPool)
	outerPool.Limiter = make(chan struct{}, conf.MaxConns)

	ctx := context.Background()

	innerPool, err := pgxpool.NewWithConfig(ctx, conf)
	if err != nil {
		log.Err(err, "Postgres connecting failed!")
		panic(err)
	}

	outerPool.InnerPool = innerPool

	log.Info("Postgres connected!")

	c, err := outerPool.Acquire(ctx)
	if err != nil {
		panic(err)
	}

	defer outerPool.Release(ctx, c)

	_, err = c.Exec(ctx, "SET TIME ZONE 'UTC'")
	if err != nil {
		log.Err(err, "Time zone set failed")
		panic(err)
	}

	return outerPool
}
