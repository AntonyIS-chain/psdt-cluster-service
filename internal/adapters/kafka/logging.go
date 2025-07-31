// internal/adapters/kafka/logging_adapter.go
package kafka

import (
	"context"
	"encoding/json"

	"github.com/AntonyIS-chain/psdt-cluster-service/internal/core/domain"
	"github.com/AntonyIS-chain/psdt-cluster-service/internal/events"
	"github.com/segmentio/kafka-go"
)

type KafkaLoggingAdapter struct {
	writer *kafka.Writer
}

func NewKafkaLoggingAdapter(brokerURL string) *KafkaLoggingAdapter {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(brokerURL),
		Balancer: &kafka.LeastBytes{},
	}

	return &KafkaLoggingAdapter{
		writer: writer,
	}
}

func (k *KafkaLoggingAdapter) LogInfo(ctx context.Context, data domain.LogData) {
	k.publish(ctx, events.TopicLogInfo, data)
}

func (k *KafkaLoggingAdapter) LogError(ctx context.Context, data domain.LogData) {
	k.publish(ctx, events.TopicLogError, data)
}

func (k *KafkaLoggingAdapter) publish(ctx context.Context, topic string, data domain.LogData) {
	msgBytes, _ := json.Marshal(data)
	msg := kafka.Message{
		Topic: topic,
		Value: msgBytes,
	}

	_ = k.writer.WriteMessages(ctx, msg) // TODO: handle error
}
