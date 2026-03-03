package domain

import (
	"context"

	"github.com/carmasearch/carma-server/api/vehicle/core"
	esCore "github.com/carmasearch/carma-server/internal/elastic/core"
)

type Repository interface {
	Create(c context.Context, v *core.Vehicle) error
	FindByID(id uint) (*core.Vehicle, error)
	FindByIDs(ids []uint64) ([]*core.Vehicle, error)
	FindBySlug(slug string) (*core.Vehicle, error)
	Update(vehicle *core.Vehicle) error
	Delete(id uint) error
	List(limit, offset int) ([]core.Vehicle, int64, error)
	// FindPaginated is used by the bulk-sync loop to stream records in pages.
	FindPaginated(c context.Context, limit, offset int) ([]*core.Vehicle, error)
}

type Service interface {
	CreateVehicle(c context.Context, vehicle *core.Vehicle) error
	GetVehicle(id uint) (*core.Vehicle, error)
	UpdateVehicle(vehicle *core.Vehicle) error
	DeleteVehicle(id uint) error
	ListVehicles(limit, offset int) ([]core.Vehicle, int64, error)

	SearchAndCompare(
		c context.Context,
		req *esCore.VehicleSearchAndCompare,
	) (int64, []*core.Vehicle, error)

	// BulkSyncToElastic re-indexes all PostgreSQL vehicles into Elasticsearch
	// in batches. Returns the total number indexed.
	BulkSyncToElastic(c context.Context, batchSize int) (int, error)
}
