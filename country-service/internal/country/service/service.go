package service

import (
	"context"
	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/countrypb"
	"country-service/internal/country/repository"
)

type CountryService struct {
	pb.UnimplementedCountryServiceServer
	Repo repository.CountryRepository
}

func NewCountryService(repo repository.CountryRepository) *CountryService {
	return &CountryService{
		Repo: repo,
	}
}

func (s *CountryService) CreateCountry(ctx context.Context, req *pb.CreateCountryRequest) (*pb.Country, error) {
	return s.Repo.CreateCountry(req)
}

func (s *CountryService) GetCountry(ctx context.Context, req *pb.GetCountryRequest) (*pb.Country, error) {
	return s.Repo.GetCountry(req)
}

func (s *CountryService) ListOfCountry(ctx context.Context, req *pb.ListOfCountryRequest) (*pb.ListOfCountryResponse, error) {
	return s.Repo.ListOfCountry(req)
}

func (s *CountryService) UpdateCountry(ctx context.Context, req *pb.UpdateCountryRequest) (*pb.Country, error) {
	return s.Repo.UpdateCountry(req)
}

func (s *CountryService) DeleteCountry(ctx context.Context, req *pb.DeleteCountryRequest) (*pb.DeleteCountryResponse, error) {
	return s.Repo.DeleteCountry(req)
}
