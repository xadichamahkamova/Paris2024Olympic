package registerservice

import (
	serve "athlete-service/internal/athlete/service"
	"fmt"
	"net"

	"google.golang.org/grpc"
	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/athletepb"
	config "athlete-service/internal/athlete/pkg/load"
)

type Service struct {
	AthleteService *serve.AthleteService
}

func NewGrpcService(s *serve.AthleteService) *Service {
	return &Service{AthleteService: s}
}

func (srv *Service) RUN(cfg config.Config) error {

	address := fmt.Sprintf(":%d", cfg.ServerPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAthleteServiceServer(grpcServer, srv.AthleteService)
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}
	return nil
}
