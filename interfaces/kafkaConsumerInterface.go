package interfaces

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

// KafkaConsumerInterface interface do consumidor kafka
type KafkaConsumerInterface interface {
	ReadMessage(timeout time.Duration) (*kafka.Message, error)
	SubscribeTopics(topics []string, rebalanceCb kafka.RebalanceCb) error
}
