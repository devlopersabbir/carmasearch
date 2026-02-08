package migrations

import (
	"github.com/carmasearch/carma-server/api/vehicle/core"
	seed "github.com/carmasearch/carma-server/internal/database/core"
	"gorm.io/gorm"
)

func Automigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&seed.SeedHistory{},
		&core.Vehicle{},
	)
}
