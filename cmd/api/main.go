package main

import (
	"log"
	"net"

	pb "github.com/ciganov-net/backend-media-service/gen/lib/proto/contracts"
	"github.com/ciganov-net/backend-media-service/internal/config"
	"github.com/ciganov-net/backend-media-service/internal/database"
	grpcHandler "github.com/ciganov-net/backend-media-service/internal/grpc"
	"github.com/ciganov-net/backend-media-service/internal/repository"
	"github.com/ciganov-net/backend-media-service/internal/router"
	"github.com/ciganov-net/backend-media-service/internal/service"
	"github.com/ciganov-net/backend-media-service/internal/storage"
	"google.golang.org/grpc"
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
	userAvatarRepo := repository.NewUserAvatarRepository(db)

	// fileHandler := handlers.NewFileHandler(fileRepo, s3Storage)

	fileService := service.NewFileService(fileRepo, s3Storage)
	mediaService := service.NewMediaService(fileService, userAvatarRepo)

	grpcServer := grpc.NewServer()
	mediaGrpc := grpcHandler.NewMediaGrpcServer(mediaService)

	pb.RegisterMediaServiceServer(grpcServer, mediaGrpc)

	r := router.SetupRouter()

	log.Printf("server running on :%s", cfg.Port)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	go func() {
		log.Println("gRPC running on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	err = r.Run(":" + cfg.Port)
	if err != nil {
		panic(err)
	}
}
