package database

import (
	"context"
	"log"

	"github.com/carmasearch/carma-server/internal/config"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

var ESClient *elasticsearch.Client

const ESIndexName = "vehicles"

func ESClientConnection(config *config.ElasticConfig) {
	esConfig := elasticsearch.Config{
		Addresses: config.ESAddresses,
		Username:  config.ESUsername,
		Password:  config.ESPassword,
	}
	es, err := elasticsearch.NewClient(esConfig)
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

func ESCreateIndexIfNotExist() {
	_, err := esapi.IndicesExistsRequest{
		Index: []string{ESIndexName},
	}.Do(context.Background(), ESClient)

	if err != nil {
		ESClient.Indices.Create(ESIndexName)
		log.Printf("✅ Index %s created", ESIndexName)
	}
	log.Printf("✅ Index %s already exists", ESIndexName)
}
