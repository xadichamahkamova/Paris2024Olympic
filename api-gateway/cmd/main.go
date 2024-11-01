package main

import (
	_ "api-gateway/docs"
	api "api-gateway/internal/http"
	athleteClient "api-gateway/internal/pkg/athlete-service"
	countryClient "api-gateway/internal/pkg/country-service"
	eventClient "api-gateway/internal/pkg/event-service"
	liveClient "api-gateway/internal/pkg/live-service"
	config "api-gateway/internal/pkg/load"
	medalClient "api-gateway/internal/pkg/medal-service"
	userClient "api-gateway/internal/pkg/user-service"
	service "api-gateway/internal/service"
	"api-gateway/logger"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	logger.InitLog()

	cfg, err := config.Load("config/config.yml")
	if err != nil {
		logger.Fatal("Failed to load config: ", err)
	}
	logger.Info("Configuration loaded successfully")

	connUserService, err := userClient.DialWithUserService(*cfg)
	if err != nil {
		logger.Fatal("Failed to connect to user service: ", err)
	}
	logger.Info("Connected to user service successfully")

	connMedalService, err := medalClient.DialWithMedalService(*cfg)
	if err != nil {
		logger.Fatal("Failed to connect to medal service: ", err)
	}
	logger.Info("Connected to medal service successfully")

	connCountryService, err := countryClient.DialWithCountryService(*cfg)
	if err != nil {
		logger.Fatal("Failed to connect to country service: ", err)
	}
	logger.Info("Connected to country service successfully")

	connEventService, err := eventClient.DialWithEventService(*cfg)
	if err != nil {
		logger.Fatal("Failed to connect to event service: ", err)
	}
	logger.Info("Connected to event service successfully")

	connAthleteService, err := athleteClient.DialWithAthleteService(*cfg)
	if err != nil {
		logger.Fatal("Failed to connect to athlete service: ", err)
	}
	logger.Info("Connected to athlete service successfully")

	connLiveService, err := liveClient.DialWithLiveService(*cfg)
	if err != nil {
		logger.Fatal("Failed to connect to liveStream service: ", err)
	}
	logger.Info("Connected to liveStream service successfully")

	s := service.NewServiceRepositoryClient(connUserService, connMedalService, connCountryService, connEventService, connAthleteService, connLiveService)

	r := api.NewGin(s)
	addr := fmt.Sprintf(":%d", cfg.ServerPort)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	go func() {
		logger.Info("Starting API Gateway on: ", addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to run API Gateway: ", err)
		}
	}()
	logger.Info("API Gateway started successfully")

	signalReceived := <-sigChan
	logger.Info("Received signal:", signalReceived)

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Fatal("Server shutdown error: ", err)
	}
	logger.Info("Graceful shutdown complete.")
}
