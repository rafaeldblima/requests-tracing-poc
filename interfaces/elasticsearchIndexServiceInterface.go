package interfaces

import (
	"context"
)

// ElasticsearchIndexServiceInterface define os métodos do serviço de indexação que utilizamos
type ElasticsearchIndexServiceInterface interface {
	IndexDocument(ctx context.Context, indexName string, documentJSON string) error
}
