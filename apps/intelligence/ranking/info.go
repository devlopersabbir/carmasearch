package ranking

type Vehicle struct {
	ID              string
	Price           float64
	Mileage         float64
	Year            int
	PowerKW         float64
	SimilarityScore float64 // from elastic search or ML
	DealScore       float64 // optional precomputed score
	PopularityScore float64 // optional business signal
}

type RankingResult struct {
	Vehicle
	FinalScore float64
	Metadata   map[string]float64
}
