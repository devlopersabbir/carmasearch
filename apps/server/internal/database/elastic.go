package database

import (
	"log"

	"github.com/carmasearch/carma-server/internal/config"
	"github.com/elastic/go-elasticsearch/v7"
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
	res, err := ESClient.Indices.Create(ESIndexName)

	if err != nil {
		log.Fatalf("failed creating index: %v", err)
	}

	defer res.Body.Close()

	log.Printf("Status code::::::::::%d", res.StatusCode)
	if res.StatusCode == 400 {
		log.Println("Index already created, skip now.")
		return
	}

	log.Printf("✅ Index %s created", ESIndexName)
}
