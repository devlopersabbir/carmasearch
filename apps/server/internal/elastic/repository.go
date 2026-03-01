package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/carmasearch/carma-server/api/vehicle/core"
	es "github.com/carmasearch/carma-server/internal/database"
	esCore "github.com/carmasearch/carma-server/internal/elastic/core"
)

func Search2(c context.Context, body *core.Vehicle) ([]uint64, int64, error) {
	return nil, 0, nil
}

func Search(ctx context.Context, req *esCore.VehicleSearchQuery) ([]uint64, int64, error) {
	if es.ESClient == nil {
		return nil, 0, nil
	}

	mustArr := []interface{}{}

	if req.Query != "" {
		mustArr = append(mustArr, map[string]interface{}{
			"query_string": map[string]interface{}{
				"query": "*" + req.Query + "*",
			},
		})
	}

	v := reflect.ValueOf(req)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := t.Field(i)

		jsonTag := strings.Split(fieldType.Tag.Get("json"), ",")[0]
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		if fieldValue.Kind() == reflect.Ptr && fieldValue.IsNil() {
			continue
		}

		if fieldValue.Kind() == reflect.Slice && fieldValue.Len() == 0 {
			continue
		}

		val := fieldValue.Interface()
		if fieldValue.Kind() == reflect.Ptr {
			val = fieldValue.Elem().Interface()
		}

		if strings.HasSuffix(jsonTag, "_from") {
			field := strings.TrimSuffix(jsonTag, "_from")
			mustArr = append(mustArr, map[string]interface{}{
				"range": map[string]interface{}{
					field: map[string]interface{}{
						"gte": val,
					},
				},
			})
		} else if strings.HasSuffix(jsonTag, "_to") {
			field := strings.TrimSuffix(jsonTag, "_to")
			mustArr = append(mustArr, map[string]interface{}{
				"range": map[string]interface{}{
					field: map[string]interface{}{
						"lte": val,
					},
				},
			})
		} else if fieldValue.Kind() == reflect.Slice {
			mustArr = append(mustArr, map[string]interface{}{
				"terms": map[string]interface{}{
					jsonTag: val,
				},
			})
		} else {
			mustArr = append(mustArr, map[string]interface{}{
				"match": map[string]interface{}{
					jsonTag: val,
				},
			})
		}
	}

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

	if req.SortBy != "" {
		order := "desc"
		if req.SortOrder != "" {
			order = req.SortOrder
		}

		sortBy := req.SortBy
		// Map `created_at` to `CreatedAt` because `core.Vehicle` misses the json tag for it.
		if sortBy == "created_at" {
			sortBy = "CreatedAt"
		} else if sortBy == "updated_at" {
			sortBy = "UpdatedAt"
		}

		queryMap["sort"] = []map[string]interface{}{
			{
				sortBy: map[string]interface{}{
					"order": order,
					"unmapped_type": "long",
				},
			},
		}
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

func IndexVehicle(s *core.Vehicle) {
	jsonBody, err := json.Marshal(s)
	if err != nil {
		log.Println("marshal error:", err)
		return
	}

	res, err := es.ESClient.Index(
		es.ESIndexName,
		bytes.NewReader(jsonBody),
		es.ESClient.Index.WithDocumentID(strconv.Itoa(int(s.ID))),
		es.ESClient.Index.WithRefresh("true"),
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
