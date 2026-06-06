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

func (s *MediaGrpcServer) DeleteAvatar(
	ctx context.Context,
	req *pb.DeleteAvatarRequest,
) (*pb.DeleteAvatarResponse, error) {
	err := s.mediaService.DeleteAvatar(req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteAvatarResponse{Ok: true}, nil
}

func (s *MediaGrpcServer) UploadFile(
	ctx context.Context,
	req *pb.UploadFileRequest,
) (*pb.UploadFileResponse, error) {
	fileID, err := s.mediaService.UploadFile(
		req.Filename,
		req.ContentType,
		req.Category,
		int64(len(req.File)),
		bytes.NewReader(req.File),
	)

	if err != nil {
		return nil, err
	}

	return &pb.UploadFileResponse{FileId: fileID}, nil
}

func (s *MediaGrpcServer) GetFile(
	ctx context.Context,
	req *pb.GetFileRequest,
) (*pb.GetFileResponse, error) {
	url, err := s.mediaService.GetFile(req.FileId)
	if err != nil {
		return nil, err
	}

	return &pb.GetFileResponse{Url: url}, nil
}
