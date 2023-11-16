package infrastructure

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/rafaeldblima/requests-tracing-poc/interfaces"
)

// ElasticsearchClientAdapter Adapter para o elasticseach,Client
type ElasticsearchClientAdapter struct {
	client *elasticsearch.Client
}

// NewElasticsearchClientAdapter construtor do adapter
func NewElasticsearchClientAdapter(client *elasticsearch.Client) *ElasticsearchClientAdapter {
	return &ElasticsearchClientAdapter{client: client}
}

// Index Ajusta adapater para ser utilizado
func (adapter *ElasticsearchClientAdapter) Index() interfaces.ElasticsearchIndexServiceInterface {
	return NewElasticsearchIndexServiceAdapter(adapter.client)
}

// EnsureIndice garante que o indice do elastic search já foi criado
func (adapter *ElasticsearchClientAdapter) EnsureIndice(indexName string) error {
	// Verifica se o índice existe
	exists, err := adapter.client.Indices.Exists([]string{indexName})
	if err != nil {
		return err
	}
	if !exists.IsError() {
		// Cria o índice se ele não existir
		createIndex, err := adapter.client.Indices.Create(indexName)
		if err != nil {
			return err
		}
		if createIndex.IsError() {
			fmt.Println(createIndex.String())
		}
	}

	return nil
}
