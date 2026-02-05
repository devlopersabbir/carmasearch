package vehicle

import (
	"errors"

	"github.com/carmasearch/carma-server/api/vehicle/core"
	"github.com/carmasearch/carma-server/api/vehicle/domain"
	"github.com/carmasearch/carma-server/arch/network"
)

type service struct {
	network.BaseService
	repo domain.Repository
}

func NewService(repo domain.Repository) domain.Service {
	return &service{
		BaseService: network.NewBaseService(),
		repo:        repo,
	}
}

func (s *service) CreateVehicle(vehicle *core.Vehicle) error {
	// Add validation logic here if needed
	if vehicle.Title == "" {
		return errors.New("title is required")
	}
	return s.repo.Create(vehicle)
}

func (s *service) GetVehicle(id uint) (*core.Vehicle, error) {
	return s.repo.FindByID(id)
}

func (s *service) UpdateVehicle(vehicle *core.Vehicle) error {
	// Ensure vehicle exists
	existing, err := s.repo.FindByID(vehicle.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("vehicle not found")
	}
	return s.repo.Update(vehicle)
}

func (s *service) DeleteVehicle(id uint) error {
	return s.repo.Delete(id)
}

func (s *service) ListVehicles(limit, offset int) ([]core.Vehicle, int64, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}
	return s.repo.List(limit, offset)
}

func (s *service) SearchVehicles(query string) ([]core.Vehicle, error) {
	// TODO: Implement ElasticSearch logic
	// For now, return empty or basic DB search if desired,
	// but user specifically asked for Elastic later, so we keep this placeholder.
	return nil, errors.New("search not implemented yet")
}
