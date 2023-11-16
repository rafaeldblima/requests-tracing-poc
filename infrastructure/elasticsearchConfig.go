package infrastructure

import (
	"crypto/tls"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

// ElasticsearchConfig objeto para trafegar as configs de conexão do elasticsearch
type ElasticsearchConfig struct {
	URL  string
	user string
	pass string
}

// NewElasticsearchConfig construtor da config
func NewElasticsearchConfig() *ElasticsearchConfig {
	return &ElasticsearchConfig{
		URL:  os.Getenv("ELASTICSEARCH_URL"),
		user: "elastic",
		pass: "elastic_pass",
	}
}

// NewElasticsearchClient construtor do client
func NewElasticsearchClient(config *ElasticsearchConfig) (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{config.URL},
		Username:  config.user,
		Password:  config.pass,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, err := client.Ping()
	if err != nil {
		return nil, err
	}
	log.Printf("Elasticsearch returned with code %d and version %s", info.StatusCode, info.String())

	return client, nil
}
