package service

import (
	"io"

	"github.com/ciganov-net/backend-media-service/internal/models"
	"github.com/ciganov-net/backend-media-service/internal/repository"
)

type MediaService struct {
	files            *FileService
	avatarRepository *repository.UserAvatarRepository
}

func NewMediaService(
	files *FileService,
	avatarRepository *repository.UserAvatarRepository,
) *MediaService {
	return &MediaService{files: files, avatarRepository: avatarRepository}
}

func (s *MediaService) UploadAvatar(
	userID string,
	filename string,
	contentType string,
	size int64,
	file io.Reader,
) (fileID string, err error) {
	existingAvatar, err := s.avatarRepository.GetByUserID(userID)
	if err != nil {
		return "", err
	}
	if existingAvatar != nil {
		if err := s.files.Delete(existingAvatar.FileID); err != nil {
			return "", err
		}
	}

	newFile, err := s.files.Upload(filename, contentType, size, file)
	if err != nil {
		return "", err
	}

	err = s.avatarRepository.Upsert(userID, newFile.ID)
	if err != nil {
		return "", err
	}

	return newFile.ID, nil
}

func (s *MediaService) DeleteAvatar(userID string) error {
	existingAvatar, err := s.avatarRepository.GetByUserID(userID)
	if err != nil {
		return err
	}
	if existingAvatar != nil {
		return s.files.Delete(existingAvatar.FileID)
	}

	return nil
}

func (s *MediaService) GetAvatar(userID string) (*models.File, error) {
	ua, err := s.avatarRepository.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	file, err := s.files.Get(ua.FileID)
	if err != nil {
		return nil, err
	}

	return file, nil
}
