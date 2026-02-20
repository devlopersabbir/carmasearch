package domain

import (
	"github.com/carmasearch/carma-server/api/vehicle/core"
)

type Repository interface {
	Create(vehicle *core.Vehicle) error
	FindByID(id uint) (*core.Vehicle, error)
	FindBySlug(slug string) (*core.Vehicle, error)
	Update(vehicle *core.Vehicle) error
	Delete(id uint) error
	List(limit, offset int) ([]core.Vehicle, int64, error)
	// Add more complex search methods here as needed, or use specific filter structs
}

type Service interface {
	CreateVehicle(vehicle *core.Vehicle) error
	GetVehicle(id uint) (*core.Vehicle, error)
	UpdateVehicle(vehicle *core.Vehicle) error
	DeleteVehicle(id uint) error
	ListVehicles(limit, offset int) ([]core.Vehicle, int64, error)
	// Placeholder for ElasticSearch integration
	SearchVehicles(filters map[string]interface{}) ([]core.Vehicle, error)
}
