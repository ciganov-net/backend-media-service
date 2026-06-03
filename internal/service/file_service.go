package service

import (
	"fmt"
	"io"
	"log"

	"github.com/ciganov-net/backend-media-service/internal/models"
	"github.com/ciganov-net/backend-media-service/internal/repository"
	"github.com/ciganov-net/backend-media-service/internal/storage"
	"github.com/google/uuid"
)

type FileService struct {
	repo    *repository.FileRepository
	storage *storage.S3Storage
}

func NewFileService(
	repo *repository.FileRepository,
	storage *storage.S3Storage,
) *FileService {
	return &FileService{repo: repo, storage: storage}
}

func (s *FileService) Upload(
	filename string,
	contentType string,
	size int64,
	file io.Reader,
) (*models.File, error) {
	id := uuid.New().String()

	objectKey := fmt.Sprintf("%s_%s", id, filename)

	if err := s.storage.Upload(objectKey, file, contentType); err != nil {
		return nil, err
	}

	newFile := models.File{
		ID:        id,
		Filename:  filename,
		ObjectKey: objectKey,
		Size:      size,
		MimeType:  contentType,
	}

	if err := s.repo.Create(newFile); err != nil {
		return nil, err
	}

	return &newFile, nil
}

func (s *FileService) Delete(id string) error {
	file, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	err = s.storage.Delete(file.ObjectKey)
	if err != nil {
		return err
	}

	err = s.repo.Delete(id)
	if err != nil {
		log.Printf(
			"CRITICAL: s3 object deleted but db delete failed for file id=%s object=%s",
			file.ID,
			file.ObjectKey,
		)

		return err
	}

	return nil
}

func (s *FileService) Get(id string) (*models.File, error) {
	return s.repo.GetByID(id)
}
