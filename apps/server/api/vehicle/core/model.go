package core

import (
	"time"

	utils "github.com/carmasearch/carma-server/internal/utils"
	"gorm.io/gorm"
)

type Vehicle struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Identifiers
	VIN        string `json:"vin" gorm:"uniqueIndex"`
	ExternalID string `json:"external_id" gorm:"index"`
	SourceURL  string `json:"source_url"`

	Slug  string `json:"slug" gorm:"uniqueIndex"`
	Title string `json:"title" gorm:"index"`

	// Make & Model
	Make         string `json:"make" gorm:"index:idx_make_model,priority:1"`
	Model        string `json:"model" gorm:"index:idx_make_model,priority:2"`
	ModelVersion string `json:"model_version" gorm:"index"`

	VehicleClass string `json:"vehicle_class" gorm:"index"`
	Category     string `json:"category" gorm:"index"`
	Condition    string `json:"condition" gorm:"index"`

	// Location
	City    string `json:"city" gorm:"index"`
	ZipCode string `json:"zip_code" gorm:"index"`
	Country string `json:"country" gorm:"index"`

	// Pricing
	Price         float64 `json:"price" gorm:"index"`
	DiscountPrice float64 `json:"discount_price" gorm:"index"`
	Currency      string  `json:"currency" gorm:"default:EUR"`

	// Registration & Mileage
	Year              int       `json:"year" gorm:"index"`
	FirstRegistration time.Time `json:"first_registration" gorm:"index"`
	Mileage           int       `json:"mileage" gorm:"index"`

	// Technical Specs
	FuelType     string `json:"fuel_type" gorm:"index"`
	Transmission string `json:"transmission" gorm:"index"`
	EngineType   string `json:"engine_type" gorm:"index"`

	PowerHP      int `json:"power_hp"`
	PowerKW      int `json:"power_kw"`
	Displacement int `json:"displacement"`

	Doors int `json:"doors" gorm:"index"`
	Seats int `json:"seats" gorm:"index"`

	// Colors & Interior
	ExteriorColor    string `json:"exterior_color" gorm:"index"`
	InteriorColor    string `json:"interior_color" gorm:"index"`
	InteriorMaterial string `json:"interior_material" gorm:"index"`

	// Ownership
	PreviousOwners int `json:"previous_owners" gorm:"index"`

	// Emissions
	CO2Emission   int    `json:"co2_emission" gorm:"index"`
	EmissionClass string `json:"emission_class" gorm:"index"`

	// Safety Features (BOOLEAN = FAST FILTER)
	ABS                    bool `json:"abs" gorm:"index"`
	ESP                    bool `json:"esp" gorm:"index"`
	TractionControl        bool `json:"traction_control" gorm:"index"`
	EmergencyBrakeAssist   bool `json:"emergency_brake_assist" gorm:"index"`
	BlindSpotAssist        bool `json:"blind_spot_assist" gorm:"index"`
	LaneAssist             bool `json:"lane_assist" gorm:"index"`
	TrafficSignRecognition bool `json:"traffic_sign_recognition" gorm:"index"`
	ISOFIX                 bool `json:"isofix" gorm:"index"`

	// Comfort
	HeatedSteeringWheel bool `json:"heated_steering_wheel"`
	StartStopSystem     bool `json:"start_stop_system"`
	HeatedSeats         bool `json:"heated_seats"`
	ElectricSeats       bool `json:"electric_seats"`
	SportSeats          bool `json:"sport_seats"`

	// Exterior Features
	FogLights          bool `json:"fog_lights"`
	AdaptiveHeadlights bool `json:"adaptive_headlights"`
	RainSensor         bool `json:"rain_sensor"`

	// Infotainment
	Radio            bool `json:"radio"`
	NavigationSystem bool `json:"navigation_system"`
	VoiceControl     bool `json:"voice_control"`
	Bluetooth        bool `json:"bluetooth"`
	USB              bool `json:"usb"`
	AppleCarPlay     bool `json:"apple_carplay"`
	AndroidAuto      bool `json:"android_auto"`

	// Seller
	SellerType    string `json:"seller_type" gorm:"index"`
	SellerName    string `json:"seller_name" gorm:"index"`
	SellerCity    string `json:"seller_city"`
	SellerCountry string `json:"seller_country"`

	// Media
	Images utils.StringArray `json:"images" gorm:"type:json"`

	// Listing
	ListingStatus string `json:"listing_status" gorm:"index"`
}
