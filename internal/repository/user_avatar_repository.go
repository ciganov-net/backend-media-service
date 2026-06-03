package repository

import (
	"context"
	"errors"

	"github.com/ciganov-net/backend-media-service/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserAvatarRepository struct {
	db *pgxpool.Pool
}

func NewUserAvatarRepository(db *pgxpool.Pool) *UserAvatarRepository {
	return &UserAvatarRepository{db: db}
}

func (r *UserAvatarRepository) Upsert(userID string, fileID string) error {
	query := `
		INSERT INTO user_avatars (user_id, file_id)
		VALUES ($1, $2)
		ON CONFLICT (user_id)
		DO UPDATE SET file_id = EXCLUDED.file_id
	`

	_, err := r.db.Exec(context.Background(), query, userID, fileID)
	return err
}

func (r *UserAvatarRepository) GetByUserID(userID string) (*models.UserAvatar, error) {
	query := `
		SELECT user_id, file_id
		FROM user_avatars
		WHERE user_id = $1
	`

	var ua models.UserAvatar

	err := r.db.QueryRow(context.Background(), query, userID).Scan(&ua.UserID, &ua.FileID)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &ua, nil
}

func (r *UserAvatarRepository) Delete(userID string) error {
	_, err := r.db.Exec(context.Background(), `DELETE FROM user_avatars WHERE user_id = $1`, userID)

	return err
}
