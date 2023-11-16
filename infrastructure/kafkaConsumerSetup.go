package infrastructure

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/linkedin/goavro/v2"
	"os"
)

// NewKafkaConsumerWithAvroCodec cria codec para decodificar o avro da mensagem
func NewKafkaConsumerWithAvroCodec(config *kafka.ConfigMap, schemaPath string) (*kafka.Consumer, *goavro.Codec, error) {
	// Inicializa o Kafka Consumer
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		return nil, nil, err
	}

	// Carrega o esquema Avro do arquivo
	schema, err := os.ReadFile(schemaPath)
	if err != nil {
		return nil, nil, err
	}

	// Cria o codec Avro
	codec, err := goavro.NewCodec(string(schema))
	if err != nil {
		return nil, nil, err
	}

	return consumer, codec, nil
}
