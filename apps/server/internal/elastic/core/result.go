package core

type VehicleCompareResult struct {
	ID      uint    `json:"id"`
	Title   string  `json:"title"`
	Make    string  `json:"make"`
	Model   string  `json:"model"`
	Year    int     `json:"year"`
	Price   float64 `json:"price"`
	Mileage int     `json:"mileage"`
	Color   int     `json:"color"`
	City    string  `json:"city"`
	Score   float64 `json:"score"` // ES relevance score
}
