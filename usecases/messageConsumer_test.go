package usecases

import (
	"context"
	"errors"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

// MockKafkaConsumer simula o Kafka Consumer
type MockKafkaConsumer struct {
	mock.Mock
}

func (m *MockKafkaConsumer) ReadMessage(timeout time.Duration) (*kafka.Message, error) {
	args := m.Called(timeout)
	return args.Get(0).(*kafka.Message), args.Error(1)
}

func (m *MockKafkaConsumer) SubscribeTopics(topics []string, rebalanceCb kafka.RebalanceCb) error {
	args := m.Called(topics, rebalanceCb)
	return args.Error(0)
}

func TestMessageConsumer_Consume(t *testing.T) {
	mockConsumer := new(MockKafkaConsumer)
	messageConsumer := NewMessageConsumer(mockConsumer)
	// Configurando a expectativa para SubscribeTopics
	mockConsumer.On("SubscribeTopics", []string{"testTopic"}, mock.Anything).Return(nil).Times(3)

	// Simulação de uma mensagem Kafka válida
	mockMessage := &kafka.Message{Value: []byte("mensagem de teste")}
	mockConsumer.On("ReadMessage", mock.Anything).Return(mockMessage, nil).Once()

	// Teste de consumo bem-sucedido
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	msg, err := messageConsumer.Consume(ctx, "testTopic")
	assert.NoError(t, err)
	assert.Equal(t, mockMessage, msg)

	// Simulação de erro ao ler a mensagem
	mockConsumer.On("ReadMessage", mock.Anything).Return((*kafka.Message)(nil), errors.New("erro de leitura")).Once()

	// Teste de tratamento de erro
	_, err = messageConsumer.Consume(ctx, "testTopic")
	assert.Error(t, err)

	// Simulação do contexto sendo cancelado
	cancelCtx, cancelFunc := context.WithCancel(context.Background())
	cancelFunc() // Cancela o contexto imediatamente

	_, err = messageConsumer.Consume(cancelCtx, "testTopic")
	assert.Error(t, err)
}
