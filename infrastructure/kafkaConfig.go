package infrastructure

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
)

// NewKafkaConsumerConfig configuração do consumidor do kafka
func NewKafkaConsumerConfig() (*kafka.ConfigMap, error) {
	return &kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
		"group.id":          "requests-tracing-poc",
		"auto.offset.reset": "earliest",
	}, nil
}
