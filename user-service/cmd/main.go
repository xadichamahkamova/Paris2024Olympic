package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	config "user-service/internal/user/pkg/load"
	pq "user-service/internal/user/pkg/postgres"
	rpc "user-service/internal/user/pkg/register-service"
	userRepo "user-service/internal/user/repository"
	userService "user-service/internal/user/service"
	"user-service/logger"
	"user-service/redis"

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

	rds, err := redis.ConnectRedis(*cfg)
	if err != nil {
		logger.Fatal("Failed to connect to redis: ", err)
	}

	repo := userRepo.NewPostgresUserRepo(db, rds)
	service := userService.NewService(repo, rds)

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
	logger.Info("User service started successfully")

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
