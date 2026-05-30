package main

import (
	"log"

	"github.com/ciganov-net/backend-media-service/internal/config"
	"github.com/ciganov-net/backend-media-service/internal/database"
	"github.com/ciganov-net/backend-media-service/internal/handlers"
	"github.com/ciganov-net/backend-media-service/internal/repository"
	"github.com/ciganov-net/backend-media-service/internal/router"
	"github.com/ciganov-net/backend-media-service/internal/storage"
)

func main() {
	cfg := config.Load()

	db := database.NewPostgres(cfg)
	defer db.Close()

	err := database.RunMigrations(db)
	if err != nil {
		panic(err)
	}

	s3Storage := storage.NewS3Storage(cfg)

	fileRepo := repository.NewFileRepository(db)

	fileHandler := handlers.NewFileHandler(fileRepo, s3Storage)

	r := router.SetupRouter(router.RouterDependencies{FileHandler: fileHandler})

	log.Printf("server running on :%s", cfg.Port)

	err = r.Run(":" + cfg.Port)
	if err != nil {
		panic(err)
	}
}
