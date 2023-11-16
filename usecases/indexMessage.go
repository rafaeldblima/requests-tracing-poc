package usecases

import (
	"context"
	"encoding/json"
	"github.com/rafaeldblima/requests-tracing-poc/domain"
	"github.com/rafaeldblima/requests-tracing-poc/interfaces"
)

// MessageIndexer adapter para indexador
type MessageIndexer struct {
	ElasticsearchClient interfaces.ElasticsearchClientInterface
}

// NewMessageIndexer construtor do adapter para indexador
func NewMessageIndexer(client interfaces.ElasticsearchClientInterface) *MessageIndexer {
	return &MessageIndexer{ElasticsearchClient: client}
}

// IndexMessage função de abstração para indexar mensagens
func (mi *MessageIndexer) IndexMessage(ctx context.Context, message *domain.Message, indexName string) error {
	documentJSON, err := json.Marshal(message)
	if err != nil {
		// Lida com o erro aqui
	}
	err = mi.ElasticsearchClient.Index().
		IndexDocument(ctx, indexName, string(documentJSON))
	if err != nil {
		return err
	}
	return nil
}
