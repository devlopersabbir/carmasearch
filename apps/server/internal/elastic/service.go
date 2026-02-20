package elastic

import (
	"github.com/carmasearch/carma-server/api/vehicle/core"
	"github.com/carmasearch/carma-server/internal/elastic/domain"
	"github.com/elastic/go-elasticsearch/v7"
	"gorm.io/gorm"
)

type vehicleService struct {
	service   domain.VehicleCompareService
	repo      domain.VehicleCompareRepository
	esClient  *elasticsearch.Client
	indexName string
	db        *gorm.DB
}

func (s *vehicleService) CompareVehicle(input *core.Vehicle) ([]core.Vehicle, error) {

	// 1️⃣ Query Elasticsearch for best matches
	ids, err := s.service.SearchSimilarVehicles(input)
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return []core.Vehicle{}, nil
	}

	// 2️⃣ Fetch full vehicles from DB in ranked order
	vehicles, err := s.repo.GetVehiclesByIDs(ids)
	if err != nil {
		return nil, err
	}

	return vehicles, nil
}
