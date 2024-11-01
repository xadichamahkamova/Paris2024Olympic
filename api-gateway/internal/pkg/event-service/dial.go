package eventservice

import (
	config "api-gateway/internal/pkg/load"
	"fmt"

	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/eventpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithEventService(cfg config.Config) (*pb.EventServiceClient, error) {

	target := fmt.Sprintf("%s:%d", cfg.EventService.Host, cfg.EventService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	clientService := pb.NewEventServiceClient(conn)
	return &clientService, nil
}
