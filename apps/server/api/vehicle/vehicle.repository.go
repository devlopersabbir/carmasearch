package vehicle

import (
	"encoding/json"
	"strings"

	"github.com/carmasearch/carma-server/api/vehicle/core"
	"github.com/carmasearch/carma-server/api/vehicle/domain"
	esClient "github.com/carmasearch/carma-server/internal/database"
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

func (r *repository) Create(vehicle *core.Vehicle) error {
	// create
	if err := r.db.Create(vehicle).Error; err != nil {
		return err
	}
	data, _ := json.Marshal(vehicle)
	_, err := esClient.ESClient.Index(
		esClient.ESIndexName,
		strings.NewReader(string(data)),
		esClient.ESClient.Index.WithRefresh("true"),
	)
	if err != nil {
		panic(err)
	}
	return nil
}

func (r *repository) FindByID(id uint) (*core.Vehicle, error) {
	var vehicle core.Vehicle
	if err := r.db.First(&vehicle, id).Error; err != nil {
		return nil, err
	}
	return &vehicle, nil
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
