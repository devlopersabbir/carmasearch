package core

import "github.com/carmasearch/carma-server/api/vehicle/core"

type CompareRequest struct {
	Url string `json:"url"`
	VehicleSearchAndCompare
}

type CompareRequestQuery struct {
	Page      int    `form:"page,default=1"`
	PageSize  int    `form:"page_size,default=20"`
	SortBy    string `form:"sort_by,default=created_at"`
	SortOrder string `form:"sort_order,default=desc"`
}

type VehicleSearchAndCompare struct {
	Url string `json:"url"`
	// Basic vehicle Info (Search Query Parameters)
	VehicleID    *string `json:"vehicle_id"`    // Vehicle ID to search a specific vehicle
	ListingURL   *string `json:"listing_url"`   // Vehicle listing URL for searching
	Title        *string `json:"title"`         // Vehicle title for search and comparison
	Make         *string `json:"make"`          // Make of the vehicle
	Model        *string `json:"model"`         // Model of the vehicle
	MakeID       *string `json:"make_id"`       // ID of the make for filtering
	ModelID      *string `json:"model_id"`      // ID of the model for filtering
	ModelVersion *string `json:"model_version"` // Model version for searching specific versions
	ModelRange   *string `json:"model_range"`   // Range of the model (like sedan, SUV, etc.)
	TrimLine     *string `json:"trim_line"`     // Trim line for detailed filtering of vehicle variants
	VehicleType  *string `json:"vehicle_type"`  // Type of vehicle (e.g., sedan, coupe)

	// Pricing and Registration range
	PriceFrom        *float64 `json:"price_from"`        // Minimum price to search for
	PriceTo          *float64 `json:"price_to"`          // Maximum price to search for
	RegistrationFrom *int     `json:"registration_from"` // Start year of registration
	RegistrationTo   *int     `json:"registration_to"`   // End year of registration

	// Mileage range
	MileageFrom *int `json:"mileage_from"` // Minimum mileage for searching
	MileageTo   *int `json:"mileage_to"`   // Maximum mileage for searching

	// CO2 Emission range
	CO2EmissionFrom *int `json:"co2_emission_from"` // Minimum CO2 emissions
	CO2EmissionTo   *int `json:"co2_emission_to"`   // Maximum CO2 emissions

	// Fuel and Transmission Filters
	FuelType     string  `json:"fuel_type"`    // List of fuel types (e.g., petrol, diesel, electric)
	Transmission string  `json:"transmission"` // List of transmission types (e.g., manual, automatic)
	DriveTrain   *string `json:"drive_train"`  // Drive train (e.g., FWD, RWD, AWD)

	// Vehicle Details for more detailed comparison
	PreviousOwners    *string `json:"previous_owners"`    // Number of previous owners
	Seats             *string `json:"seats"`              // Number of seats in the vehicle
	Doors             *string `json:"doors"`              // Number of doors in the vehicle
	ExteriorColors    *string `json:"exterior_colors"`    // List of possible exterior colors
	InteriorColors    *string `json:"interior_colors"`    // List of possible interior colors
	InteriorMaterials *string `json:"interior_materials"` // Materials used for interior (e.g., leather, fabric)

	// Safety Features
	ABS                    *bool `json:"abs"`                      // Whether the vehicle has ABS
	ESP                    *bool `json:"esp"`                      // Whether the vehicle has ESP
	TractionControl        *bool `json:"traction_control"`         // Whether the vehicle has traction control
	EmergencyBrakeAssist   *bool `json:"emergency_brake_assist"`   // Emergency brake assist system
	BlindSpotAssist        *bool `json:"blind_spot_assist"`        // Blind spot assist system
	LaneAssist             *bool `json:"lane_assist"`              // Lane assist system
	TrafficSignRecognition *bool `json:"traffic_sign_recognition"` // Traffic sign recognition system

	// Comfort Features
	HeatedSteeringWheel *bool `json:"heated_steering_wheel"` // Heated steering wheel
	StartStopSystem     *bool `json:"start_stop_system"`     // Start/Stop system
	HeatedSeats         *bool `json:"heated_seats"`          // Heated seats
	ElectricSeats       *bool `json:"electric_seats"`        // Electric adjustable seats
	SportSeats          *bool `json:"sport_seats"`           // Sport seats for more comfortable driving

	// Exterior Features (For comparison)
	FogLights          *bool `json:"fog_lights"`          // Whether the vehicle has fog lights
	AdaptiveHeadlights *bool `json:"adaptive_headlights"` // Adaptive headlights for improved vision
	RainSensor         *bool `json:"rain_sensor"`         // Rain sensor for automatic wiper activation

	// Infotainment
	Radio            *bool `json:"radio"`             // Whether the vehicle has a radio
	NavigationSystem *bool `json:"navigation_system"` // Navigation system
	VoiceControl     *bool `json:"voice_control"`     // Voice control for vehicle commands
	Bluetooth        *bool `json:"bluetooth"`         // Bluetooth connectivity
	USB              *bool `json:"usb"`               // USB connectivity
	AppleCarPlay     *bool `json:"apple_carplay"`     // Apple CarPlay integration
	AndroidAuto      *bool `json:"android_auto"`      // Android Auto integration

	// Location Information (for filtering searches by location)
	CountryCode *string `json:"country_code"` // Country of the vehicle
	City        *string `json:"city"`         // City where the vehicle is listed
	PostalCode  *string `json:"postal_code"`  // Postal code of the vehicle's location
	Street      *string `json:"street"`       // Street where the vehicle is located
	SellerName  *string `json:"seller_name"`  // Seller name for the vehicle listing

	// Additional Fields for more comparison capabilities
	CompareRequestQuery
}

type VehicleSearchQueryResponse struct {
	Total    uint64         `json:"total"`
	Page     int            `json:"page"`
	Pagesize int            `json:"page_size"`
	Vehicles []core.Vehicle `json:"vehicles"`
}
