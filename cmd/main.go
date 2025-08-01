package cmd

import (
	"log"

	"github.com/AntonyIS-chain/psdt-cluster-service/config"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/adapters/db"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/adapters/logging"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/adapters/rest/app/handlers"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/services.go"
)

func RunService() {

	conf, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	logrusLogger := logging.NewLogrusLogger()
	kafkaLogger := logging.NewKafkaLogger("localhost:9092", "logs-topic")

	logger := logging.NewCompositeLogger(logrusLogger, kafkaLogger)

	logger.Info("Service started", map[string]interface{}{
		"mdule": "main",
	})

	// Initialize DB client
	dbClient, err := db.NewDBClient(conf)
	if err != nil {
		log.Fatalf("Failed to initialize DB client: %v", err)
	}

	// Repos
	clusterRepo := db.NewClusterRepository(dbClient.DB)
	tribeRepo := db.NewTribeRepository(dbClient.DB)
	squadRepo := db.NewSquadRepository(dbClient.DB)
	userRepo := db.NewUserRepository(dbClient.DB)

	// Services
	clusterSvc := services.NewClusterManagementService(clusterRepo, logger)
	tribeSvc := services.NewTribeManagementService(tribeRepo, logger)
	squadSvc := services.NewSquadManagementService(squadRepo, logger)
	userSvc := services.NewUserManagementService(userRepo, logger)

	handlers.InitGinRoutes(clusterSvc, tribeSvc, squadSvc, userSvc, conf)

}
