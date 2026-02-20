package elastic

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/carmasearch/carma-server/api/vehicle/core"
	esCore "github.com/carmasearch/carma-server/internal/elastic/core"
)

func (s *vehicleService) SearchSimilarVehicles(input *esCore.VehicleSearchQuery) ([]uint, error) {

	query := map[string]interface{}{
		"size": 20,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": []interface{}{
					map[string]interface{}{
						"term": map[string]interface{}{
							"listing_status": "active",
						},
					},
				},
				"must": []interface{}{
					map[string]interface{}{
						"match": map[string]interface{}{
							"make": *input.Make,
						},
					},
					map[string]interface{}{
						"match": map[string]interface{}{
							"model": *input.Model,
						},
					},
				},
				"should": []interface{}{
					map[string]interface{}{
						"range": map[string]interface{}{
							"year": map[string]interface{}{
								"gte":   *input.RegistrationFrom - 2,
								"lte":   *input.RegistrationTo + 2,
								"boost": 2,
							},
						},
					},
					map[string]interface{}{
						"range": map[string]interface{}{
							"price": map[string]interface{}{
								"gte":   *input.PriceFrom * 0.85,
								"lte":   *input.PriceTo * 1.15,
								"boost": 3,
							},
						},
					},
				},
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}

	res, err := s.esClient.Search(
		s.esClient.Search.WithContext(context.Background()),
		s.esClient.Search.WithIndex(s.indexName),
		s.esClient.Search.WithBody(&buf),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	hits := r["hits"].(map[string]interface{})["hits"].([]interface{})

	var ids []uint
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		idFloat := source["id"].(float64)
		ids = append(ids, uint(idFloat))
	}

	return ids, nil
}

func (r *vehicleService) GetVehiclesByIDs(ids []uint) ([]core.Vehicle, error) {

	var vehicles []core.Vehicle

	if err := r.db.Where("id IN ?", ids).Find(&vehicles).Error; err != nil {
		return nil, err
	}

	// Preserve order
	vehicleMap := make(map[uint]core.Vehicle)
	for _, v := range vehicles {
		vehicleMap[v.ID] = v
	}

	ordered := make([]core.Vehicle, 0, len(ids))
	for _, id := range ids {
		if v, ok := vehicleMap[id]; ok {
			ordered = append(ordered, v)
		}
	}

	return ordered, nil
}
