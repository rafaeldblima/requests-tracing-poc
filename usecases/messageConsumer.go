package usecases

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rafaeldblima/requests-tracing-poc/interfaces"
)

// MessageConsumer adapter para consumidor
type MessageConsumer struct {
	Consumer interfaces.KafkaConsumerInterface
}

// NewMessageConsumer construtor do adapter para consumidor
func NewMessageConsumer(consumer interfaces.KafkaConsumerInterface) *MessageConsumer {
	return &MessageConsumer{Consumer: consumer}
}

// Consume consome mensagens de tópico
func (mc *MessageConsumer) Consume(ctx context.Context, topic string) (*kafka.Message, error) {
	err := mc.Consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return nil, err
	}

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			msg, err := mc.Consumer.ReadMessage(-1)
			if err != nil {
				return nil, err
			}
			return msg, nil
		}
	}
}
