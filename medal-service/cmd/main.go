package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	config "medal-service/internal/medal/pkg/load"
	pq "medal-service/internal/medal/pkg/postgres"
	rpc "medal-service/internal/medal/pkg/register-service"
	medalRepo "medal-service/internal/medal/repository"
	medalService "medal-service/internal/medal/service"
	"medal-service/logger"

	"google.golang.org/grpc"
)

func main() {

	logger.InitLog()

	cfg, err := config.Load("config/config.yml")
	if err != nil {
		logger.Fatal("Failed to load config: ", err)
	}
	logger.Info("Configuration loaded successfully")

	db, err := pq.InitDB(*cfg)
	if err != nil {
		logger.Fatal("Failed to connect to database: ", err)
	}
	logger.Info("Connected to the database successfully")

	repo := medalRepo.NewPostgresMedalRepo(db)
	service := medalService.NewMedalService(repo)

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
	logger.Info("Medal service started successfully")

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
