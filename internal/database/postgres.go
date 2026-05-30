package database

import (
	"context"
	"log"

	"github.com/ciganov-net/backend-media-service/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
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
