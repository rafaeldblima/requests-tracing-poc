package main

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/linkedin/goavro/v2"
	"github.com/rafaeldblima/requests-tracing-poc/infrastructure"
	"github.com/rafaeldblima/requests-tracing-poc/usecases"
	"log"
)

func main() {
	messages := make(chan *kafka.Message, 100)
	// Configuração do Kafka Consumer
	kafkaConfig, err := infrastructure.NewKafkaConsumerConfig()
	if err != nil {
		log.Fatalf("Erro na configuração do Kafka: %v", err)
	}

	// Cria o Kafka Consumer com o Codec Avro
	consumer, codec, err := infrastructure.NewKafkaConsumerWithAvroCodec(kafkaConfig, "avro/message.avsc")
	if err != nil {
		log.Fatalf("Erro na inicialização do Kafka Consumer: %v", err)
	}
	defer func(consumer *kafka.Consumer) {
		err := consumer.Close()
		if err != nil {
			log.Fatalf("Erro no fechamento do consumidor.")
		}
	}(consumer)

	// Inicia o processo de consumo
	ctx := context.Background()
	topic := "tracing-user"
	messageConsumer := usecases.NewMessageConsumer(consumer)

	config := infrastructure.NewElasticsearchConfig()
	client, err := infrastructure.NewElasticsearchClient(config)
	if err != nil {
		log.Fatalf("Failed to create Elasticsearch client: %v", err)
	}
	clientAdapter := infrastructure.NewElasticsearchClientAdapter(client)

	err = clientAdapter.EnsureIndice("tracing")
	if err != nil {
		log.Fatalf("Erro ao verificar/criar índice: %v", err)
	}
	indexer := usecases.NewMessageIndexer(clientAdapter)

	for i := 0; i < 5; i++ { // Você pode ajustar o número de goroutines conforme necessário
		go processMessages(messages, codec, indexer)
	}
	for {
		// Consume uma mensagem
		msg, err := messageConsumer.Consume(ctx, topic)
		if err != nil {
			log.Printf("Erro ao consumir mensagem: %v", err)
			continue
		}

		// Coloque a mensagem no canal para ser processada pelas goroutines
		messages <- msg
	}
}

func processMessages(messages <-chan *kafka.Message, codec *goavro.Codec, indexer *usecases.MessageIndexer) {
	for msg := range messages {
		// Converte a mensagem Avro para o objeto Message
		message, err := usecases.ConvertAvroToMessage(msg.Value, codec)
		if err != nil {
			log.Printf("Erro ao converter mensagem Avro: %v", err)
			continue
		}
		log.Printf(message.AccountID)
		err = indexer.IndexMessage(context.Background(), message, "tracing")
		if err != nil {
			log.Printf("Failed to index message: %v", err)
		}
	}
}
