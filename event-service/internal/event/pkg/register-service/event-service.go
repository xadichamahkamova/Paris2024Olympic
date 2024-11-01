package registerservice

import (
	serve "event-service/internal/event/service"
	"fmt"
	"net"

	"google.golang.org/grpc"
	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/eventpb"
	config "event-service/internal/event/pkg/load"
)

type Service struct {
	EventService *serve.EventService
}

func NewGrpcService(s *serve.EventService) *Service {
	return &Service{EventService: s}
}

func (srv *Service) RUN(cfg config.Config) error {

	address := fmt.Sprintf(":%d", cfg.ServerPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterEventServiceServer(grpcServer, srv.EventService)
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}
	return nil
}
