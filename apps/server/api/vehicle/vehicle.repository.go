package vehicle

import (
	"context"

	"github.com/carmasearch/carma-server/api/vehicle/core"
	"github.com/carmasearch/carma-server/api/vehicle/domain"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) domain.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(c context.Context, v *core.Vehicle) error {
	return r.db.Create(v).Error
}

func (r *repository) FindByID(id uint) (*core.Vehicle, error) {
	var vehicle core.Vehicle
	if err := r.db.First(&vehicle, id).Error; err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func (r *repository) FindByIDs(ids []uint64) ([]*core.Vehicle, error) {
	var vehicles []*core.Vehicle

	if len(ids) == 0 {
		return vehicles, nil
	}

	if err := r.db.
		Where("id IN ?", ids).
		Find(&vehicles).Error; err != nil {
		return nil, err
	}

	return vehicles, nil
}

// Find by Slug
func (r *repository) FindBySlug(slug string) (*core.Vehicle, error) {
	var vehicle core.Vehicle

	if err := r.db.First(&vehicle, "slug = ?", slug).Error; err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func (r *repository) Update(vehicle *core.Vehicle) error {
	return r.db.Save(vehicle).Error
}

func (r *repository) Delete(id uint) error {
	var vehicle core.Vehicle
	return r.db.Delete(&vehicle, id).Error
}

func (r *repository) List(limit, offset int) ([]core.Vehicle, int64, error) {
	var vehicle core.Vehicle
	var vehicles []core.Vehicle
	var count int64

	if err := r.db.Model(&vehicle).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Limit(limit).Offset(offset).Find(&vehicles).Error; err != nil {
		return nil, 0, err
	}

	return vehicles, count, nil
}

// FindPaginated fetches a page of vehicles without a COUNT query —
// intended for the bulk-sync loop where total count is not needed.
func (r *repository) FindPaginated(c context.Context, limit, offset int) ([]*core.Vehicle, error) {
	var vehicles []*core.Vehicle
	if err := r.db.WithContext(c).
		Order("id ASC").
		Limit(limit).
		Offset(offset).
		Find(&vehicles).Error; err != nil {
		return nil, err
	}
	return vehicles, nil
}
