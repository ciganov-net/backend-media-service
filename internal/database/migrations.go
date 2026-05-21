package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func RunMigrations(db *pgxpool.Pool) error {
	query := `
	CREATE TABLE IF NOT EXISTS files (
		id UUID PRIMARY KEY,
		filename TEXT NOT NUlL,
		object_key TEXT NOT NULL,
		size BIGINT NOT NULL,
		mime_type TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW()
	);
	`

	_, err := db.Exec(context.Background(), query)

	return err
}
