package seed

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/carmasearch/carma-server/api/vehicle/core"
	"gorm.io/gorm"
)

var (
	makes  = []string{"BMW", "Audi", "Mercedes-Benz", "Volkswagen", "Toyota", "Ford"}
	models = map[string][]string{
		"BMW":           {"320", "X3", "X5", "118"},
		"Audi":          {"A3", "A4", "A6", "Q5"},
		"Mercedes-Benz": {"C200", "E220", "GLC"},
		"Volkswagen":    {"Golf", "Passat", "Polo"},
		"Toyota":        {"Corolla", "Camry", "Yaris"},
		"Ford":          {"Focus", "Fiesta", "Kuga"},
	}

	cities        = []string{"Berlin", "Munich", "Hamburg", "Stuttgart", "Frankfurt"}
	colors        = []string{"Black", "White", "Grey", "Blue", "Red"}
	bodyTypes     = []string{"Sedan", "Hatchback", "SUV"}
	fuels         = []string{"Petrol", "Diesel", "Hybrid"}
	transmissions = []string{"Manual", "Automatic"}
	priceRatings  = []string{"Good", "Fair", "High"}
	featuresPool  = []string{
		"Navigation system",
		"Bluetooth",
		"Parking sensors",
		"LED headlights",
		"Cruise control",
		"Heated seats",
		"Lane assist",
	}
)

func SeedVehicles(db *gorm.DB, count int) error {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < count; i++ {
		make := randomFrom(makes)
		model := randomFrom(models[make])

		year := rand.Intn(8) + 2017
		mileage := rand.Intn(120_000-5_000) + 5_000

		vehicle := core.Vehicle{
			Make:             make,
			Model:            model,
			ModelDescription: fmt.Sprintf("%s %s", make, model),
			ModelVersion:     "Base",
			VIN:              randomVIN(),

			Title: fmt.Sprintf("%s %s %d", make, model, year),
			Slug:  fmt.Sprintf("%s-%s-%d-%d", make, model, year, i),

			VehicleClass: "Car",
			Category:     randomFrom(bodyTypes),
			Condition:    "Used",

			City:    randomFrom(cities),
			ZipCode: fmt.Sprintf("%05d", rand.Intn(99999)),
			Country: "DE",

			Price:            float64(rand.Intn(45_000-9_000) + 9_000),
			DiscountPrice:    0,
			Currency:         "EUR",
			PriceRatingLabel: randomFrom(priceRatings),

			Year:              year,
			FirstRegistration: time.Date(year, time.Month(rand.Intn(12)+1), 1, 0, 0, 0, 0, time.UTC),
			Mileage:           mileage,

			EngineType:   randomFrom(fuels),
			FuelType:     randomFrom(fuels),
			Transmission: randomFrom(transmissions),
			Gearbox:      randomFrom(transmissions),

			BodyType:     randomFrom(bodyTypes),
			Displacement: rand.Intn(2000-1400) + 1400,
			PowerHP:      rand.Intn(220-110) + 110,
			PowerKW:      rand.Intn(160-80) + 80,

			Color: randomFrom(colors),
			Doors: 5,
			Seats: 5,

			Features: randomFeatures(),

			DamageUnrepaired:       false,
			Roadworthy:             true,
			AccidentDamaged:        false,
			NumberOfPreviousOwners: rand.Intn(3) + 1,
			Warranty:               rand.Intn(2) == 1,

			EmissionClass:       "Euro 6d",
			CO2Emission:         rand.Intn(160-95) + 95,
			ConsumptionCombined: roundFloat(rand.Float64()*3+4, 1),
			ConsumptionCity:     roundFloat(rand.Float64()*3+5, 1),
			ConsumptionHighway:  roundFloat(rand.Float64()*2+4, 1),

			SellerType:    "dealer",
			SellerName:    fmt.Sprintf("%s Autohaus", make),
			SellerCity:    randomFrom(cities),
			SellerCountry: "DE",

			Images: core.StringArray{
				"https://img.mobile.de/car1.jpg",
				"https://img.mobile.de/car2.jpg",
			},

			ExternalID:    fmt.Sprintf("ext-%d", rand.Int63()),
			SourceURL:     "https://suchen.mobile.de",
			ListingStatus: "active",
		}

		if err := db.Create(&vehicle).Error; err != nil {
			return err
		}
	}

	return nil
}

func randomFrom[T any](list []T) T {
	return list[rand.Intn(len(list))]
}

func randomFeatures() core.StringArray {
	count := rand.Intn(5) + 3
	seen := map[string]bool{}
	var features core.StringArray

	for len(features) < count {
		f := randomFrom(featuresPool)
		if !seen[f] {
			seen[f] = true
			features = append(features, f)
		}
	}
	return features
}

func randomVIN() string {
	const chars = "ABCDEFGHJKLMNPRSTUVWXYZ0123456789"
	b := make([]byte, 17)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func roundFloat(v float64, precision int) float64 {
	p := math.Pow(10, float64(precision))
	return math.Round(v*p) / p
}
