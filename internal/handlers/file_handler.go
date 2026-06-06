package handlers

import (
	"github.com/ciganov-net/backend-media-service/internal/service"
)

type FileHandler struct {
	service *service.FileService
}

func NewFileHandler(
	service *service.FileService,
) *FileHandler {
	return &FileHandler{service: service}
}

// was used to test in the browser [unrefactored]
// func (h *FileHandler) GetFiles(c *gin.Context) {
// 	files, err := h.repo.GetAll()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, files)
// }

// func (h *FileHandler) UploadFile(c *gin.Context) {
// 	fileHeader, err := c.FormFile("file")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
// 		return
// 	}

// 	file, err := fileHeader.Open()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	defer file.Close()

// 	uploadedFile, err := h.service.Upload(
// 		fileHeader.Filename,
// 		fileHeader.Header.Get("Content-Type"),
// 		fileHeader.Size,
// 		file,
// 	)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, uploadedFile)
// }

// func (h *FileHandler) DeleteFile(c *gin.Context) {
// 	id := c.Param("id")

// 	err := h.service.Delete(id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "file deleted"})
// }

// was used to test in the browser [unrefactored]
// func (h *FileHandler) GetDownloadURL(c *gin.Context) {
// 	id := c.Param("id")

// 	file, err := h.repo.GetByID(id)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "file not found"})

// 		return
// 	}

// 	url, err := h.storage.GetPresignedURL(file.ObjectKey)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"url": url})
// }
