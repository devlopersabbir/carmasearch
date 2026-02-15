package elastic

import (
	"encoding/json"
	"io"

	"github.com/carmasearch/carma-server/internal/elastic/core"
)

func parseCompareResults(body io.Reader) ([]core.VehicleCompareResult, error) {
	var raw struct {
		Hits struct {
			Hits []struct {
				Score  float64 `json:"_score"`
				Source struct {
					ID      uint    `json:"id"`
					Title   string  `json:"title"`
					Make    string  `json:"make"`
					Model   string  `json:"model"`
					Year    int     `json:"year"`
					Price   float64 `json:"price"`
					Mileage int     `json:"mileage"`
					Color   int     `json:"color"`
					City    string  `json:"city"`
				} `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(body).Decode(&raw); err != nil {
		return nil, err
	}

	results := make([]core.VehicleCompareResult, 0)
	for _, h := range raw.Hits.Hits {
		results = append(results, core.VehicleCompareResult{
			ID:      h.Source.ID,
			Title:   h.Source.Title,
			Make:    h.Source.Make,
			Model:   h.Source.Model,
			Year:    h.Source.Year,
			Price:   h.Source.Price,
			Mileage: h.Source.Mileage,
			City:    h.Source.City,
			Score:   h.Score,
		})
	}

	return results, nil
}
