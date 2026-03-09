package ranking

import (
	"math"
	"strconv"
	"strings"
)

// Vehicle is the lightweight view used by the ranking engine.
// Only numeric fields needed for scoring are kept here.
type Vehicle struct {
	ID              string
	Price           float64
	Mileage         float64
	Year            int
	PowerKW         float64
	SimilarityScore float64 // from Elasticsearch or ML
	DealScore       float64 // optional precomputed score
	PopularityScore float64 // optional business signal
}

// RankingResult is a Vehicle plus its computed scores.
type RankingResult struct {
	Vehicle
	FinalScore float64
	Metadata   map[string]float64
}

// --- helpers ---

// parseFloat safely converts a string to float64, returning 0 on failure.
func parseFloat(s string) float64 {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return v
}

// parseInt safely converts a string to int, returning 0 on failure.
func parseInt(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return v
}

// extractYear returns the first 4-digit year found among the candidate year
// fields (ProductionYear, ConstructionYear, FirstRegistration).
func extractYear(productionYear, constructionYear, firstRegistration string) int {
	for _, s := range []string{productionYear, constructionYear, firstRegistration} {
		s = strings.TrimSpace(s)
		if len(s) >= 4 {
			if v, err := strconv.Atoi(s[:4]); err == nil && v > 1900 {
				return v
			}
		}
	}
	return 0
}

// DealQuality classifies a listing price against the predicted fair price.
// Returns "great", "fair", or "overpriced".
func DealQuality(listingPrice, predictedPrice float64) string {
	if predictedPrice <= 0 {
		return "unknown"
	}
	diff := math.Abs(listingPrice-predictedPrice) / predictedPrice
	if listingPrice <= predictedPrice*0.95 {
		return "great" // 5 %+ below market
	}
	if diff <= 0.10 {
		return "fair" // within 10 % of market
	}
	return "overpriced"
}
