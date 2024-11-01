package registerservice

import (
	serve "country-service/internal/country/service"
	"fmt"
	"net"

	"google.golang.org/grpc"
	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/countrypb"
	config "country-service/internal/country/pkg/load"
)

type Service struct {
	CountryService *serve.CountryService
}

func NewGrpcService(s *serve.CountryService) *Service {
	return &Service{
		CountryService: s,
	}
}

func (srv *Service) RUN(cfg config.Config) error {

	address := fmt.Sprintf(":%d", cfg.ServerPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCountryServiceServer(grpcServer, srv.CountryService)
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}
	return nil
}
