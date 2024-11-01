package registerservice 

import (
	serve "user-service/internal/user/service"
	"fmt"
	"net"

	"google.golang.org/grpc"
    pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/userpb"
	config "user-service/internal/user/pkg/load"
)

type Service struct {
	UserService *serve.UserService
}

func NewGrpcService(userService *serve.UserService) *Service {
	return &Service{
		UserService: userService,
	}
}

func (srv *Service) RUN(cfg config.Config) error {

	address := fmt.Sprintf(":%d", cfg.UserServicePort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, srv.UserService)
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}
	return nil
}
