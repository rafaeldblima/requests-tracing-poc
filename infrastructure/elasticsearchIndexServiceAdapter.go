package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"io"
	"strings"
)

// ElasticsearchIndexServiceAdapter adapter para o indexador do elastic search
type ElasticsearchIndexServiceAdapter struct {
	client *elasticsearch.Client
}

// NewElasticsearchIndexServiceAdapter construtor do indexador do elastic search
func NewElasticsearchIndexServiceAdapter(client *elasticsearch.Client) *ElasticsearchIndexServiceAdapter {
	return &ElasticsearchIndexServiceAdapter{client: client}
}

// IndexDocument indexa documento no elastic search
func (adapter *ElasticsearchIndexServiceAdapter) IndexDocument(ctx context.Context, indexName string, documentJSON string) error {
	req := esapi.IndexRequest{
		Index: indexName,
		Body:  strings.NewReader(documentJSON),
	}

	res, err := req.Do(ctx, adapter.client)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			_ = fmt.Errorf("error")
		}
	}(res.Body)

	if res.IsError() {
		return errors.New("erro ao indexar documento")
	}

	return nil
}
