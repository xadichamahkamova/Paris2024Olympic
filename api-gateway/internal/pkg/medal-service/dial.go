package medalservice

import (
	config "api-gateway/internal/pkg/load"
	"fmt"

	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/medalspb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithMedalService(cfg config.Config) (*pb.MedalServiceClient, error) {

	target := fmt.Sprintf("%s:%d", cfg.MedalService.Host, cfg.MedalService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	clientService := pb.NewMedalServiceClient(conn)
	return &clientService, nil
}
