package logging

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

type KafkaLogger struct {
	writer *kafka.Writer
}

type LogMessage struct {
	Level     string                 `json:"level"`
	Message   string                 `json:"message"`
	Fields    map[string]interface{} `json:"fields"`
	Timestamp time.Time              `json:"timestamp"`
	LogID     string                 `json:"log_id"`
}

func NewKafkaLogger(broker, topic string) *KafkaLogger {
	return &KafkaLogger{
		writer: &kafka.Writer{
			Addr:         kafka.TCP(broker),
			Topic:        topic,
			RequiredAcks: kafka.RequireAll,
			Balancer:     &kafka.LeastBytes{},
		},
	}
}

func (k *KafkaLogger) log(level string, message string, fields map[string]interface{}) {
	logEntry := LogMessage{
		Level:     level,
		Message:   message,
		Fields:    fields,
		Timestamp: time.Now(),
		LogID:     uuid.NewString(),
	}

	msgBytes, err := json.Marshal(logEntry)
	if err != nil {
		return // optionally log to stdout as fallback
	}

	ctx := context.Background() // 👈 Use a valid context
	err = k.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(level),
		Value: msgBytes,
	})
	if err != nil {
		// Optional: log to stdout or fallback logger
	}
}

func (k *KafkaLogger) Info(msg string, fields map[string]interface{}) {
	k.log("INFO", msg, fields)
}

func (k *KafkaLogger) Error(msg string, fields map[string]interface{}) {
	k.log("ERROR", msg, fields)
}

func (k *KafkaLogger) Debug(msg string, fields map[string]interface{}) {
	k.log("DEBUG", msg, fields)
}

func (k *KafkaLogger) Warn(msg string, fields map[string]interface{}) {
	k.log("WARN", msg, fields)
}

func (k *KafkaLogger) Close() error {
	return k.writer.Close()
}
