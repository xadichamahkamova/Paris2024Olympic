package userservice

import (
	config "api-gateway/internal/pkg/load"
	"fmt"

	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithUserService(cfg config.Config) (*pb.UserServiceClient, error) {

	target := fmt.Sprintf("%s:%d", cfg.UserService.Host, cfg.UserService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	clientService := pb.NewUserServiceClient(conn)
	return &clientService, nil
}
