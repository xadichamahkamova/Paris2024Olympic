package main

import (
	"context"
	config "live-service/internal/live/pkg/load"
	mongosh "live-service/internal/live/pkg/mongosh"
	rpc "live-service/internal/live/pkg/register-service"
	liveRepo "live-service/internal/live/repository"
	liveService "live-service/internal/live/service"
	"live-service/logger"
	"log"
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

	db, err := mongosh.NewConnection(cfg)
	if err != nil {
		logger.Fatal("Failed to connect to database: ", err)
	}
	logger.Info("Connected to the database successfully")

	repo := liveRepo.NewMongoshLiveRepository(*db)
	service := liveService.NewEventService(repo)

	var wg sync.WaitGroup
	wg.Add(1)
	r := rpc.NewGrpcService(service)

	gServer := grpc.NewServer()
	go func() {
		defer wg.Done()
		log.Printf("Live Service running on :%d port", cfg.ServerPort)
		if err = r.RUN(*cfg); err != nil {
			logger.Fatal("Failed to run gRPC service: ", err)
		}
	}()
	logger.Info("Live service started successfully")

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
