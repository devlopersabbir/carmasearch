package ranking

import "math"

// PredictFairPrice predicts market fair price using simple intelligence logic
func PredictFairPrice(
	marketAvgPrice float64,
	vehicleYear int,
	currentYear int,
	mileage float64,
) float64 {

	// Age penalty
	ageFactor := float64(currentYear-vehicleYear) * 0.02

	// Mileage penalty
	mileageFactor := mileage / 200000.0

	predicted := marketAvgPrice * (1 - ageFactor - mileageFactor)

	if predicted < 0 {
		return marketAvgPrice * 0.5
	}

	return predicted
}

// PriceScore calculates how good the listing price is
func PriceScore(listingPrice, predictedPrice float64) float64 {
	if predictedPrice <= 0 {
		return 0
	}

	score := 1 - math.Abs(listingPrice-predictedPrice)/predictedPrice

	if score < 0 {
		return 0
	}

	if score > 1 {
		return 1
	}

	return score
}
