package database

import (
	"context"
	"log"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/esapi"
)

var ESClient *elasticsearch.Client

const SearchIndex = "bmw"

func ESClientConnection() {
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		panic(err)
	}
	ESClient = client
	log.Println("Elastic Search Client Connected")
}

func ESCreateIndexIfNotExist() {
	_, err := esapi.IndicesExistsRequest{
		Index: []string{SearchIndex},
	}.Do(context.Background(), ESClient)

	if err != nil {
		ESClient.Indices.Create(SearchIndex)
	}
}
