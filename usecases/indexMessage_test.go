package usecases

import (
	"context"
	_ "errors"
	"github.com/rafaeldblima/requests-tracing-poc/interfaces"
	"testing"

	"github.com/rafaeldblima/requests-tracing-poc/domain"
	_ "github.com/rafaeldblima/requests-tracing-poc/interfaces"
	"github.com/stretchr/testify/mock"
)

// Crie um mock do ElasticsearchClientInterface
type MockElasticsearchClient struct {
	mock.Mock
}

func (m *MockElasticsearchClient) Index() interfaces.ElasticsearchIndexServiceInterface {
	args := m.Called()
	return args.Get(0).(interfaces.ElasticsearchIndexServiceInterface)
}

func (m *MockElasticsearchClient) IndexDocument(ctx context.Context, indexName string, documentJSON string) error {
	args := m.Called(ctx, indexName, documentJSON)
	return args.Error(0)
}

func TestMessageIndexer_IndexMessage(t *testing.T) {
	// Crie uma instância do mock
	mockClient := new(MockElasticsearchClient)

	// Crie uma instância do MessageIndexer com o mockClient
	messageIndexer := NewMessageIndexer(mockClient)

	// Crie um contexto de teste
	ctx := context.TODO()

	// Crie um objeto de mensagem de teste
	message := &domain.Message{
		// Preencha os campos do objeto de mensagem conforme necessário
	}

	// Defina o nome do índice de teste
	indexName := "meu_indice"

	// Configure o mockClient para esperar uma chamada com os argumentos corretos e retornar um erro simulado (ou nil se for bem-sucedido)
	mockClient.On("Index").Return(mockClient)
	mockClient.On("IndexDocument", ctx, indexName, mock.AnythingOfType("string")).Return(nil)

	// Chame o método IndexMessage do MessageIndexer
	err := messageIndexer.IndexMessage(ctx, message, indexName)

	// Verifique se não ocorreu um erro durante a chamada
	if err != nil {
		t.Errorf("esperava erro nil, mas obteve %v", err)
	}

	// Verifique se o método IndexDocument do mockClient foi chamado com os argumentos corretos
	mockClient.AssertCalled(t, "IndexDocument", ctx, indexName, mock.AnythingOfType("string"))
}
