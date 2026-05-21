package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rin-cast-9/s3_test/backend/internal/handlers"
)

type RouterDependencies struct {
	FileHandler *handlers.FileHandler
}

func SetupRouter(deps RouterDependencies) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},

		AllowMethods: []string{
			"GET",
			"POST",
			"DELETE",
		},

		AllowHeaders: []string{
			"Origin",
			"Content-Type",
		},

		MaxAge: 12 * time.Hour,
	}))

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/files", deps.FileHandler.GetFiles)
	r.POST("/upload", deps.FileHandler.UploadFile)
	r.DELETE("/files/:id", deps.FileHandler.DeleteFile)
	r.GET("/files/:id/download", deps.FileHandler.GetDownloadURL)

	return r
}
