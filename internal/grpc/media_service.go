package grpc

import (
	"bytes"
	"context"

	pb "github.com/ciganov-net/backend-media-service/gen/lib/proto/contracts"
	"github.com/ciganov-net/backend-media-service/internal/service"
)

type MediaGrpcServer struct {
	pb.UnimplementedMediaServiceServer
	mediaService *service.MediaService
}

func NewMediaGrpcServer(ms *service.MediaService) *MediaGrpcServer {
	return &MediaGrpcServer{mediaService: ms}
}

func (s *MediaGrpcServer) UploadAvatar(
	ctx context.Context,
	req *pb.UploadAvatarRequest,
) (*pb.UploadAvatarResponse, error) {
	fileID, err := s.mediaService.UploadAvatar(
		req.UserId,
		req.Filename,
		req.ContentType,
		int64(len(req.File)),
		bytes.NewReader(req.File),
	)

	if err != nil {
		return nil, err
	}

	return &pb.UploadAvatarResponse{FileId: fileID}, nil
}
