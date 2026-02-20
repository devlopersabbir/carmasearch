package seed

import (
	"math/rand"
	"time"

	"gorm.io/gorm"
)

var ()

func SeedVehicles(db *gorm.DB, count int) error {
	rand.Seed(time.Now().UnixNano())

	return nil
}
