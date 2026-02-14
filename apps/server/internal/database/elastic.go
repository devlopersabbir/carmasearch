package database

import (
	"context"
	"log"

	"github.com/carmasearch/carma-server/internal/config"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

var ESClient *elasticsearch.Client

func ESClientConnection(config *config.ElasticConfig) {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			config.Addr,
		},
		MaxRetries:    3,
		RetryOnStatus: []int{502, 503, 504},
	})
	if err != nil {
		log.Fatalf("failed to create elastic client: %v", err)
	}

	// Ping cluster
	res, err := es.Info()
	if err != nil {
		log.Fatalf("elastic ping failed: %v", err)
	}
	defer res.Body.Close()

	ESClient = es
	log.Println("✅ Elasticsearch connected")
}

func ESCreateIndexIfNotExist(indexName string) {
	_, err := esapi.IndicesExistsRequest{
		Index: []string{indexName},
	}.Do(context.Background(), ESClient)

	if err != nil {
		ESClient.Indices.Create(indexName)
	}
}
