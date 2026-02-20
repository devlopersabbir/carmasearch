package elastic

import (
	"log"

	"github.com/carmasearch/carma-server/api/vehicle/core"
	esCore "github.com/carmasearch/carma-server/internal/elastic/core"
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

func (s *vehicleService) CompareVehicle(input *esCore.VehicleSearchQuery) ([]core.Vehicle, error) {
	log.Println("=========log=========", input)
	ids, err := s.service.SearchSimilarVehicles(input)
	log.Println("=========ids=========", ids)
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return []core.Vehicle{}, nil
	}

	vehicles, err := s.repo.GetVehiclesByIDs(ids)
	if err != nil {
		return nil, err
	}

	return vehicles, nil
}
