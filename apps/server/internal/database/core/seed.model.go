package core

import "time"

type SeedHistory struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"uniqueIndex"`
	CreatedAt time.Time
}
