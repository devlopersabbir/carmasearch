package ranking

import (
	"github.com/devloeprsabbir/go-elasticsearch/models"
)

// FromModel converts a models.Vehicle into the lightweight ranking.Vehicle
// struct. It parses numeric fields that are stored as strings in the DB.
func FromModel(v models.Vehicle, esScore float64) Vehicle {
	price := parseFloat(v.Price)

	mileage := parseFloat(v.MileageKM)
	if mileage == 0 {
		mileage = parseFloat(v.MileageInKM)
	}

	year := extractYear(v.ProductionYear, v.ConstructionYear, v.FirstRegistration)

	powerKW := parseFloat(v.PowerKW)

	return Vehicle{
		ID:              v.UniqueID,
		Price:           price,
		Mileage:         mileage,
		Year:            year,
		PowerKW:         powerKW,
		SimilarityScore: esScore,
		DealScore:       0, // extend later with ML signal
		PopularityScore: 0, // extend later with view-count signal
	}
}

// MarketAvgPrice computes a simple market average from a slice of ranking
// vehicles, ignoring zero-price entries.
func MarketAvgPrice(vehicles []Vehicle) float64 {
	var sum float64
	var count int
	for _, v := range vehicles {
		if v.Price > 0 {
			sum += v.Price
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return sum / float64(count)
}
