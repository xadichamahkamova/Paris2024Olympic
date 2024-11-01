package eventservice

import (
	config "api-gateway/internal/pkg/load"
	"fmt"

	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/livepb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithLiveService(cfg config.Config) (*pb.LiveStreamServiceClient, error) {

	target := fmt.Sprintf("%s:%d", cfg.LiveService.Host, cfg.LiveService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	clientService := pb.NewLiveStreamServiceClient(conn)
	return &clientService, nil
}
