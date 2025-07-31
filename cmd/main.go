package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/AntonyIS-chain/psdt-cluster-service/config"
	"github.com/AntonyIS-chain/psdt-cluster-service/github.com/AntonyIS-chain/proto/loggingpb"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/adapters/app/handlers"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/adapters/repository"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/credentials/insecure"
)

func RunService() {
	// Load config
	conf, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize DB client
	dbClient, err := repository.NewDBClient(conf)
	if err != nil {
		log.Fatalf("Failed to initialize DB client: %v", err)
	}

	// Setup gRPC connection to Logging Service
	fmt.Println("LOGGING_SERVICE_URL>>>", conf.LOGGING_SERVICE_URL)
	conn, err := grpc.Dial(
		conf.LOGGING_SERVICE_URL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithConnectParams(grpc.ConnectParams{
			Backoff:           backoff.DefaultConfig,
			MinConnectTimeout: 5 * time.Second,
		}),
	)
	if err != nil {
		log.Fatalf("Failed to connect to logging service: %v", err)
	}
	defer conn.Close()

	// Initialize Logging Service Client
	loggingClient := loggingpb.NewLoggingServiceClient(conn)
	loggerService := logging.NewLoggingClient(loggingClient) // <-- uses your `LoggingClient` struct

	// Repositories with logging
	clusterRepo := repository.NewClusterRepository(dbClient.DB, loggerService)
	tribeRepo := repository.NewTribeRepository(dbClient.DB, loggerService)
	squadRepo := repository.NewSquadRepository(dbClient.DB, loggerService)
	userRepo := repository.NewUserRepository(dbClient.DB, loggerService)

	// Services
	clusterService := services.NewClusterManagementService(clusterRepo)
	tribeService := services.NewTribeManagementService(tribeRepo)
	userService := services.NewUserManagementService(userRepo)
	squadService := services.NewSquadManagementService(squadRepo)

	// Initialize HTTP routes
	handlers.InitGinRoutes(clusterService, tribeService, squadService, userService, conf)
}
