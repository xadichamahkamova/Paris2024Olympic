package main

import (
	"context"
	config "event-service/internal/event/pkg/load"
	pq "event-service/internal/event/pkg/postgres"
	rpc "event-service/internal/event/pkg/register-service"
	eventRepo "event-service/internal/event/repository"
	eventService "event-service/internal/event/service"
	"event-service/logger"
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

	repo := eventRepo.NewPostgresEventRepository(db)
	service := eventService.NewEventService(repo)

	var wg sync.WaitGroup
	wg.Add(1)
	r := rpc.NewGrpcService(service)

	gServer := grpc.NewServer()
	go func() {
		defer wg.Done()
		if err = r.RUN(*cfg); err != nil {
			logger.Fatal("Failed to run gRPC service: ", err)
		}
	}()
	logger.Info("Event service started successfully")
	
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
