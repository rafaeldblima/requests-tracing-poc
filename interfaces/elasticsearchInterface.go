package interfaces

// ElasticsearchClientInterface define os métodos do cliente Elasticsearch que utilizamos
type ElasticsearchClientInterface interface {
	Index() ElasticsearchIndexServiceInterface
}
