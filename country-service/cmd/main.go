package main

import (
	"context"
	config "country-service/internal/country/pkg/load"
	pq "country-service/internal/country/pkg/postgres"
	rpc "country-service/internal/country/pkg/register-service"
	countryRepo "country-service/internal/country/repository"
	countryService "country-service/internal/country/service"
	"country-service/logger"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"google.golang.org/grpc"
)

func main() {

	logger.InitLog()

	cfg, err := config.Load("config/config.yml")
	if err != nil {
		logger.Fatal("Failed to load config: ", err)
	}
	logger.Info("Configuration loaded successfully")

	db, err := pq.ConnectDB(*cfg)
	if err != nil {
		logger.Fatal("Failed to connect to database: ", err)
	}
	logger.Info("Connected to the database successfully")

	repo := countryRepo.NewPostgresCountryRepository(db)
	service := countryService.NewCountryService(repo)

	var wg sync.WaitGroup
	wg.Add(1)

	r := rpc.NewGrpcService(service)

	gServer := grpc.NewServer()
	go func() {
		defer wg.Done()
		if err := r.RUN(*cfg); err != nil {
			logger.Fatal("Failed to run gRPC service: ", err)
		}
	}()
	logger.Info("Country service started successfully")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	logger.Info("Received signal:", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	gServer.GracefulStop()

	<-ctx.Done()
	logger.Info("Graceful shutdown complete.")
}
