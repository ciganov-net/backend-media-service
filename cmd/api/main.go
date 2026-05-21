package main

import (
	"log"

	"github.com/rin-cast-9/s3_test/backend/internal/config"
	"github.com/rin-cast-9/s3_test/backend/internal/database"
	"github.com/rin-cast-9/s3_test/backend/internal/handlers"
	"github.com/rin-cast-9/s3_test/backend/internal/repository"
	"github.com/rin-cast-9/s3_test/backend/internal/router"
	"github.com/rin-cast-9/s3_test/backend/internal/storage"
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
