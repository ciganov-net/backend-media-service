package repository

import (
	"context"

	"github.com/ciganov-net/backend-media-service/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FileRepository struct {
	db *pgxpool.Pool
}

func NewFileRepository(db *pgxpool.Pool) *FileRepository {
	return &FileRepository{db: db}
}

func (r *FileRepository) GetAll() ([]models.File, error) {
	query := `
		SELECT
			id,
			filename,
			object_key,
			size,
			mime_type,
			created_at
		FROM files
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	files := make([]models.File, 0)

	for rows.Next() {
		var file models.File

		err := rows.Scan(
			&file.ID,
			&file.Filename,
			&file.ObjectKey,
			&file.Size,
			&file.MimeType,
			&file.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		files = append(files, file)
	}

	return files, nil
}

func (r *FileRepository) Create(file models.File) error {
	query := `
	INSERT INTO files (
		id,
		filename,
		object_key,
		size,
		mime_type
	)
	VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(
		context.Background(),
		query,
		file.ID,
		file.Filename,
		file.ObjectKey,
		file.Size,
		file.MimeType,
	)

	return err
}

func (r *FileRepository) GetByID(id string) (*models.File, error) {
	query := `
	SELECT
		id,
		filename,
		object_key,
		size,
		mime_type,
		created_at
	FROM files
	WHERE id = $1
	`

	var file models.File

	err := r.db.QueryRow(context.Background(), query, id).Scan(
		&file.ID,
		&file.Filename,
		&file.ObjectKey,
		&file.Size,
		&file.MimeType,
		&file.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &file, nil
}

func (r *FileRepository) Delete(id string) error {
	query := `
	DELETE FROM files
	WHERE id = $1
	`

	_, err := r.db.Exec(context.Background(), query, id)

	return err
}
