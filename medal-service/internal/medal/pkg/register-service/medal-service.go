package registerservice

import (
	serve "medal-service/internal/medal/service"
	"fmt"
	"net"

	"google.golang.org/grpc"
    pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/medalspb"
	config "medal-service/internal/medal/pkg/load"
)

type Service struct {
	MedalService *serve.MedalService
}

func NewGrpcService(medalService *serve.MedalService) *Service {
	return &Service{
		MedalService: medalService,
	}
}

func (srv *Service) RUN(cfg config.Config) error {

	address := fmt.Sprintf(":%d", cfg.MedalServicePort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterMedalServiceServer(grpcServer, srv.MedalService)
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}
	return nil
}
