package service

import (
	"context"
	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/athletepb"
	"athlete-service/internal/athlete/repository"
)

type AthleteService struct {
	pb.UnimplementedAthleteServiceServer
	Repo repository.AthleteRepository
}

func NewAthleteService(repo repository.AthleteRepository) *AthleteService {
	return &AthleteService{
		Repo: repo,
	}
}

func(s *AthleteService) CreateAthlete(ctx context.Context, req *pb.CreateAthleteRequest) (*pb.Athlete, error) {
	return s.Repo.CreateAthlete(req)
}

func(s *AthleteService) GetAthlete(ctx context.Context, req *pb.GetAthleteRequest) (*pb.GetAthleteResponse, error) {
	return s.Repo.GetAthlete(req)
}

func(s *AthleteService) ListOfAthlete(ctx context.Context, req *pb.ListOfAthleteRequest) (*pb.ListOfAthleteResponse, error) {
	return s.Repo.ListAthletes(req)
}

func(s *AthleteService) UpdateAthlete(ctx context.Context, req *pb.UpdateAthleteRequest) (*pb.Athlete, error) {
	return s.Repo.UpdateAthlete(req)
}

func(s *AthleteService) DeleteAthlete(ctx context.Context, req *pb.DeleteAthleteRequest) (*pb.DeleteAthleteResponse, error) {
	return s.Repo.DeleteAthlete(req)
}
