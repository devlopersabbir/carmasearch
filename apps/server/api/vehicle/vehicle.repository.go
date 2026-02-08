package vehicle

import (
	"github.com/carmasearch/carma-server/api/vehicle/core"
	"github.com/carmasearch/carma-server/api/vehicle/domain"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

var vehicle core.Vehicle

func NewRepository(db *gorm.DB) domain.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(vehicle *core.Vehicle) error {
	return r.db.Create(vehicle).Error
}

func (r *repository) FindByID(id uint) (*core.Vehicle, error) {
	if err := r.db.First(&vehicle, id).Error; err != nil {
		return nil, err
	}
	return &vehicle, nil
}

// Find by Slug
func (r *repository) FindByTitle(slug string) (*core.Vehicle, error) {
	if err := r.db.First(&vehicle, slug).Error; err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func (r *repository) Update(vehicle *core.Vehicle) error {
	return r.db.Save(vehicle).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&core.Vehicle{}, id).Error
}

func (r *repository) List(limit, offset int) ([]core.Vehicle, int64, error) {
	var vehicles []core.Vehicle
	var count int64

	if err := r.db.Model(&core.Vehicle{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Limit(limit).Offset(offset).Find(&vehicles).Error; err != nil {
		return nil, 0, err
	}

	return vehicles, count, nil
}
