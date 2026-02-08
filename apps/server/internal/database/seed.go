package database

import (
	"fmt"

	"github.com/carmasearch/carma-server/api/vehicle/seed"
	"github.com/carmasearch/carma-server/internal/database/core"
	"gorm.io/gorm"
)

func RunVehicleSeedOnce(db *gorm.DB) error {
	const seedName = "vehicle_seed_v1"

	var count int64
	db.Model(&core.SeedHistory{}).
		Where("name = ?", seedName).
		Count(&count)

	if count > 0 {
		fmt.Println("🚫 Vehicle seed already executed, skipping")
		return nil
	}

	if err := seed.SeedVehicles(db, 50); err != nil {
		return err
	}

	return db.Create(&core.SeedHistory{
		Name: seedName,
	}).Error
}
