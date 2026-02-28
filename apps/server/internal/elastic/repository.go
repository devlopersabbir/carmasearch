package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/carmasearch/carma-server/api/vehicle/core"
	es "github.com/carmasearch/carma-server/internal/database"
	esCore "github.com/carmasearch/carma-server/internal/elastic/core"
)

func IndexVehicle(s *core.Vehicle) {
	jsonBody, err := json.Marshal(s)
	if err != nil {
		log.Println("marshal error:", err)
		return
	}

	res, err := es.ESClient.Index(
		es.ESIndexName,
		bytes.NewReader(jsonBody),
		es.ESClient.Index.WithDocumentID(strconv.Itoa(int(s.ID))), // VERY IMPORTANT
		es.ESClient.Index.WithRefresh("true"),                     // force refresh so you can see immediately
	)
	if err != nil {
		log.Println("index request failed:", err)
		return
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Println("indexing error:", res.String())
		return
	}

	log.Println("indexed todo successfully:", s.ID)
}

func Search(ctx context.Context, req *esCore.VehicleSearchQuery) ([]uint64, int64, error) {
	if es.ESClient == nil {
		return nil, 0, nil
	}

	mustArr := []interface{}{}
	// all searching, maching, comparing will be go here

	var query interface{}
	if len(mustArr) == 0 {
		query = map[string]interface{}{"match_all": map[string]interface{}{}}
	} else {
		query = map[string]interface{}{
			"bool": map[string]interface{}{
				"must": mustArr,
			},
		}
	}
	queryMap := map[string]interface{}{
		"query": query,
		"from":  (req.Page - 1) * req.PageSize,
		"size":  req.PageSize,
	}
	queryBytes, _ := json.Marshal(queryMap)
	res, err := es.ESClient.Search(
		es.ESClient.Search.WithContext(ctx),
		es.ESClient.Search.WithIndex(es.ESIndexName),
		es.ESClient.Search.WithBody(bytes.NewReader(queryBytes)),
		es.ESClient.Search.WithTrackTotalHits(true),
	)

	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Println("search error:", res.String())
		return nil, 0, nil
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, 0, err
	}

	hitsMap, ok := r["hits"].(map[string]interface{})
	if !ok {
		return nil, 0, nil
	}

	hits, ok := hitsMap["hits"].([]interface{})
	if !ok {
		return nil, 0, nil
	}

	var total int64
	if totalMap, ok := hitsMap["total"].(map[string]interface{}); ok {
		if val, ok := totalMap["value"].(float64); ok {
			total = int64(val)
		}
	}

	var ids []uint64
	for _, hit := range hits {
		if hitMap, ok := hit.(map[string]interface{}); ok {
			if source, ok := hitMap["_source"].(map[string]interface{}); ok {
				if idFloat, ok := source["id"].(float64); ok {
					ids = append(ids, uint64(idFloat))
				}
			}
		}
	}

	return ids, total, nil
}
