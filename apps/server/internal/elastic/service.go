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
	esClient  *elasticsearch.Client
	indexName string
	db        *gorm.DB
}

func NewVehicleCompareService(esClient *elasticsearch.Client, indexName string, db *gorm.DB) domain.VehicleCompareService {
	return &vehicleService{
		esClient:  esClient,
		indexName: indexName,
		db:        db,
	}
}

func (s *vehicleService) CompareVehicle(input *esCore.VehicleSearchQuery) ([]core.Vehicle, error) {
	log.Println("=========log=========", input)
	ids, err := s.SearchSimilarVehicles(input)
	log.Println("=========ids=========", ids)
	if err != nil {
		return nil, err
	}

	if len(ids) == 0 {
		return []core.Vehicle{}, nil
	}

	vehicles, err := s.GetVehiclesByIDs(ids)
	if err != nil {
		return nil, err
	}

	return vehicles, nil
}
