package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rin-cast-9/s3_test/backend/internal/config"
)

func NewPostgres(cfg config.Config) *pgxpool.Pool {
	pool, err := pgxpool.New(
		context.Background(),
		cfg.PostgresURL,
	)

	if err != nil {
		panic(err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	log.Println("connected to postgres")

	return pool
}
