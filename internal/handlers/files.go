package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ciganov-net/backend-media-service/internal/models"
	"github.com/ciganov-net/backend-media-service/internal/repository"
	"github.com/ciganov-net/backend-media-service/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type FileHandler struct {
	repo    *repository.FileRepository
	storage *storage.S3Storage
}

func NewFileHandler(
	repo *repository.FileRepository,
	storage *storage.S3Storage,
) *FileHandler {
	return &FileHandler{repo: repo, storage: storage}
}

func (h *FileHandler) GetFiles(c *gin.Context) {
	files, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, files)
}

func (h *FileHandler) UploadFile(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer file.Close()

	id := uuid.New().String()

	objectKey := fmt.Sprintf("%s_%s", id, fileHeader.Filename)

	err = h.storage.Upload(objectKey, file, fileHeader.Header.Get("Content-Type"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newFile := models.File{
		ID:        id,
		Filename:  fileHeader.Filename,
		ObjectKey: objectKey,
		Size:      fileHeader.Size,
		MimeType:  fileHeader.Header.Get("Content-Type"),
	}

	err = h.repo.Create(newFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newFile)
}

func (h *FileHandler) DeleteFile(c *gin.Context) {
	id := c.Param("id")

	file, err := h.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
		return
	}

	err = h.storage.Delete(file.ObjectKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.repo.Delete(id)
	if err != nil {
		log.Printf(
			"CRITICAL: s3 object deleted but db delete failed for file id=%s object=%s",
			file.ID,
			file.ObjectKey,
		)

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "file deleted"})
}

func (h *FileHandler) GetDownloadURL(c *gin.Context) {
	id := c.Param("id")

	file, err := h.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "file not found"})

		return
	}

	url, err := h.storage.GetPresignedURL(file.ObjectKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"url": url})
}
