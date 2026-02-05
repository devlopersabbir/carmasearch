package domain

import "time"

type VehicleDomain struct {
	ID uint

	PostedAt  *time.Time
	UpdatedAt *time.Time
}
