package core

type VehicleCompareQuery struct {
	VehicleID uint    `json:"vehicle_id"`
	Limit     int     `json:"limit"`
	PriceDiff float64 `json:"price_diff"` // ± percentage (e.g. 0.15 = 15%)
}
