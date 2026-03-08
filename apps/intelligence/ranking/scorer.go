package ranking

import "sort"

// RankVehicles ranks vehicles based on scoring intelligence
func RankVehicles(
	query Vehicle,
	candidates []Vehicle,
	marketAvgPrice float64,
	config RankingConfig,
	currentYear int,
) []RankingResult {
	results := make([]RankingResult, 0)

	// Calculate predicted fair price once
	predictedPrice := PredictFairPrice(
		marketAvgPrice,
		query.Year,
		currentYear,
		query.Mileage,
	)

	// Score each candidate vehicle
	for _, candidate := range candidates {
		priceScore := PriceScore(candidate.Price, predictedPrice)

		mileageScore := 1 - (candidate.Mileage / 200000.0)

		yearScore := float64(candidate.Year-2000) / float64(currentYear-2000)

		similarityScore := candidate.SimilarityScore

		popularityScore := candidate.PopularityScore

		finalScore :=
			priceScore*config.PriceWeight +
				mileageScore*config.MileageWeight +
				yearScore*config.YearWeight +
				similarityScore*config.SimilarityWeight +
				popularityScore*config.PopularityWeight

		results = append(results, RankingResult{
			Vehicle:    candidate,
			FinalScore: finalScore,
			Metadata: map[string]float64{
				"price_score":      priceScore,
				"mileage_score":    mileageScore,
				"year_score":       yearScore,
				"similarity_score": similarityScore,
				"popularity_score": popularityScore,
				"predicted_price":  predictedPrice,
			},
		})
	}

	// Sort by best ranking first
	sort.Slice(results, func(i, j int) bool {
		return results[i].FinalScore > results[j].FinalScore
	})

	return results
}
