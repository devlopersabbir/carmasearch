package domain

import (
	"github.com/carmasearch/carma-server/api/vehicle/core"
)

type CompareService interface {
	CompareVehicle(vehicle *core.Vehicle) ([]core.Vehicle, error)
}

type VehicleCompareRepository interface {
	GetVehiclesByIDs(ids []uint) ([]core.Vehicle, error)
}

type VehicleCompareService interface {
	CompareVehicle(vehicle *core.Vehicle) ([]core.Vehicle, error)
	SearchSimilarVehicles(input *core.Vehicle) ([]uint, error)
}
