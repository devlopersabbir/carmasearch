package domain

import (
	"github.com/carmasearch/carma-server/api/vehicle/core"
	esCore "github.com/carmasearch/carma-server/internal/elastic/core"
)

type VehicleCompareRepository interface {
	GetVehiclesByIDs(ids []uint) ([]core.Vehicle, error)
}

type VehicleCompareService interface {
	CompareVehicle(vehicle *esCore.VehicleSearchQuery) ([]core.Vehicle, error)
	SearchSimilarVehicles(input *esCore.VehicleSearchQuery) ([]uint, error)
}
