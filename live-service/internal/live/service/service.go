package service

import (
	"context"
	"live-service/internal/live/repository"

	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/livepb"
)

type LiveService struct {
	pb.UnimplementedLiveStreamServiceServer
	Repo repository.LiveRepository
}

func NewEventService(repo repository.LiveRepository) *LiveService {
	return &LiveService{
		Repo: repo,
	}
}

func (s *LiveService) CreateLiveStream(ctx context.Context, req *pb.LiveStream) (*pb.ResponseMessage, error) {
	return s.Repo.CreateLiveStream(req)
}

func (s *LiveService) GetLiveStream(ctx context.Context, req *pb.GetStreamRequest) (*pb.LiveStream, error) {
	return s.Repo.GetLiveStream(req)
}
