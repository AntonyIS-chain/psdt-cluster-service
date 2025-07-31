package cmd

import "github.com/AntonyIS-chain/psdt-cluster-service/internal/adapters/logging"

func RunService() {
	logrusLogger := logging.NewLogrusLogger()
	kafkaLogger := logging.NewKafkaLogger("localhost:9092", "logs-topic")

	logger := logging.NewCompositeLogger(logrusLogger, kafkaLogger)

	logger.Info("Service started", map[string]interface{}{
		"mdule": "main",
	})
}
