package migrations

import (
	"github.com/carmasearch/carma-server/api/vehicle/core"
	"gorm.io/gorm"
)

func Automigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&core.Vehicle{},
	)
}
