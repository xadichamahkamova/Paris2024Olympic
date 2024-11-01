package countryservice

import (
	config "api-gateway/internal/pkg/load"
	"fmt"

	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/countrypb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialWithCountryService(cfg config.Config) (*pb.CountryServiceClient, error) {

	target := fmt.Sprintf("%s:%d", cfg.CountryService.Host, cfg.CountryService.Port)
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	clientService := pb.NewCountryServiceClient(conn)
	return &clientService, nil
}
