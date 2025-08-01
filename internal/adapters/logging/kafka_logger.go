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
	LogID     string                 `json:"log_id"`    // Unique ID for log entry
	Timestamp time.Time              `json:"timestamp"` // Time the log was created
	Level     string                 `json:"level"`     // INFO, ERROR, DEBUG, etc.
	Message   string                 `json:"message"`   // Human-readable message
	Fields    map[string]interface{} `json:"fields"`    // Additional structured metadata

	// Contextual Metadata
	ServiceName    string   `json:"service_name"`      // Which service generated it
	ServiceVersion string   `json:"service_version"`   // Deployed version or git hash
	Environment    string   `json:"environment"`       // e.g., production, staging
	Hostname       string   `json:"hostname"`          // Host or pod name
	InstanceID     string   `json:"instance_id"`       // Instance or container ID
	TraceID        string   `json:"trace_id"`          // For tracing logs
	SpanID         string   `json:"span_id"`           // For tracing spans
	CorrelationID  string   `json:"correlation_id"`    // To tie logs across services
	UserID         string   `json:"user_id,omitempty"` // If action is user-specific
	Tags           []string `json:"tags,omitempty"`    // Optional tags for filtering
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
        ServiceName: "psdt-cluster-service",
        ServiceVersion: "V1.0.0",
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
