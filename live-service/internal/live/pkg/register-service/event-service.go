package registerservice

import (
	"fmt"
	serve "live-service/internal/live/service"
	"net"

	config "live-service/internal/live/pkg/load"

	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/livepb"
	"google.golang.org/grpc"
)

type Service struct {
	LiveService *serve.LiveService
}

func NewGrpcService(s *serve.LiveService) *Service {
	return &Service{LiveService: s}
}

func (srv *Service) RUN(cfg config.Config) error {

	address := fmt.Sprintf(":%d", cfg.ServerPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterLiveStreamServiceServer(grpcServer, srv.LiveService)
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}
	return nil
}
