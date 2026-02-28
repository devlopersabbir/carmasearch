package domain

import (
	"context"

	"github.com/carmasearch/carma-server/api/vehicle/core"
	esCore "github.com/carmasearch/carma-server/internal/elastic/core"
)

type VehicleCompareRepository interface {
	IndexVehcile(v *core.Vehicle)
	Search(ctx context.Context, req *esCore.VehicleSearchQuery) ([]uint64, int64, error)
	GetVehiclesByIDs(ids []uint64) ([]core.Vehicle, error)
}

type VehicleCompareService interface {
	CompareVehicle(vehicle *esCore.VehicleSearchQuery) ([]core.Vehicle, error)
	SearchSimilarVehicles(input *esCore.VehicleSearchQuery) ([]uint64, error)
}
