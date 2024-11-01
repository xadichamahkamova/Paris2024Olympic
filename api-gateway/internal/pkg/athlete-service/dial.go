package athleteservice

import (
	config "api-gateway/internal/pkg/load"
	"fmt"

	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/athletepb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func DialWithAthleteService(cfg config.Config) (*pb.AthleteServiceClient, error) {

	target := fmt.Sprintf("%s:%d", cfg.AthleteService.Host, cfg.AthleteService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	clientService := pb.NewAthleteServiceClient(conn)
	return &clientService, nil
}
