package search

type VehicleESDoc struct {
	ID       uint    `json:"id"`
	Make     string  `json:"make"`
	Model    string  `json:"model"`
	Year     int     `json:"year"`
	Price    float64 `json:"price"`
	Mileage  int     `json:"mileage"`
	FuelType string  `json:"fuel_type"`
	Gearbox  string  `json:"gearbox"`
	BodyType string  `json:"body_type"`
	City     string  `json:"city"`
	Country  string  `json:"country"`
}
