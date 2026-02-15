package elastic

import (
	"bytes"
	"context"
	"encoding/json"

	vehicleCore "github.com/carmasearch/carma-server/api/vehicle/core"
	"github.com/carmasearch/carma-server/api/vehicle/domain"
	"github.com/carmasearch/carma-server/internal/database"
	elasticCore "github.com/carmasearch/carma-server/internal/elastic/core"
)

const VehicleIndex = "vehicles"

type CompareRepository struct{}

func NewCompareRepository() domain.CompareRepository {
	return &CompareRepository{}
}

func (r *CompareRepository) FindSimilar(v *vehicleCore.Vehicle, limit int) ([]elasticCore.VehicleCompareResult, error) {
	if limit <= 0 {
		limit = 10
	}
	query := map[string]interface{}{
		"size": limit,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []interface{}{
					map[string]interface{}{
						"term": map[string]interface{}{
							"make": v.Make,
						},
					},
					map[string]interface{}{
						"term": map[string]interface{}{
							"model": v.Model,
						},
					},
				},
				"should": []interface{}{
					map[string]interface{}{
						"range": map[string]interface{}{
							"year": map[string]interface{}{
								"gte": v.Year - 2,
								"lte": v.Year + 2,
							},
						},
					},
					map[string]interface{}{
						"range": map[string]interface{}{
							"price": map[string]interface{}{
								"gte": v.Price * 0.85,
								"lte": v.Price * 1.15,
							},
						},
					},
					map[string]interface{}{
						"range": map[string]interface{}{
							"mileage": map[string]interface{}{
								"lte": v.Mileage + 20000,
							},
						},
					},
					map[string]interface{}{
						"term": map[string]interface{}{
							"color": v.Color,
						},
					},
				},
			},
		},
	}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(query)

	res, err := database.ESClient.Search(
		database.ESClient.Search.WithIndex(VehicleIndex),
		database.ESClient.Search.WithBody(&buf),
		database.ESClient.Search.WithContext(context.Background()),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	return parseCompareResults(res.Body)
}
