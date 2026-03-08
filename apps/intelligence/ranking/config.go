package ranking

// RankingConfig allows tuning ranking behavior without code change
type RankingConfig struct {
	PriceWeight      float64
	MileageWeight    float64
	YearWeight       float64
	SimilarityWeight float64
	PopularityWeight float64
}

// DefaultRankingConfig is good for marketplace MVP
func DefaultRankingConfig() RankingConfig {
	return RankingConfig{
		PriceWeight:      0.4,
		MileageWeight:    0.2,
		YearWeight:       0.2,
		SimilarityWeight: 0.2,
		PopularityWeight: 0.1,
	}
}
