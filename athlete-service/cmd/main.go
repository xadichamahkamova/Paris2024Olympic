package main

import (
	config "athlete-service/internal/athlete/pkg/load"
	pq "athlete-service/internal/athlete/pkg/postgres"
	rpc "athlete-service/internal/athlete/pkg/register-service"
	athleteRepo "athlete-service/internal/athlete/repository"
	athleteService "athlete-service/internal/athlete/service"
	"athlete-service/logger"
	"context"
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

	repo := athleteRepo.NewPostgresAthleteRepository(db)

	service := athleteService.NewAthleteService(repo)
	r := rpc.NewGrpcService(service)

	var wg sync.WaitGroup
	wg.Add(1)

	gServer := grpc.NewServer()
	go func() {
		defer wg.Done()
		if err = r.RUN(*cfg); err != nil {
			logger.Fatal("Failed to run gRPC service: ", err)
		}
	}()
	logger.Info("Athlete service started successfully")

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
