package core

type VehicleSearchQuery struct {
	// Basic vehicle Info
	VehicleID    *string  `json:"vehicle_id"`
	ListingURL   *string  `json:"listing_url"`
	Title        *string  `json:"title"`
	Make         *string  `json:"make"`
	Model        *string  `json:"model"`
	FuelType     []string `json:"fuel_type"`
	Transmission []string `json:"transmission"`
	SellerName   *string  `json:"seller_name"`

	// price & registration range
	PriceFrom        *float64 `json:"price_from"`
	PriceTo          *float64 `json:"price_to"`
	RegistrationFrom *int     `json:"registration_from"`
	RegistrationTo   *int     `json:"registration_to"`
	MileageFrom      *int     `json:"mileage_from"`
	MileageTo        *int     `json:"mileage_to"`
	CO2EmissionFrom  *int     `json:"co2_emission_from"`
	CO2EmissionTo    *int     `json:"co2_emission_to"`

	// vehicle details
	PreviousOwners    *int     `json:"previous_owners"`
	Seats             *int     `json:"seats"`
	Doors             *int     `json:"doors"`
	ExteriorColors    []string `json:"exterior_colors"`
	InteriorColors    []string `json:"interior_colors"`
	InteriorMaterials []string `json:"interior_materials"`

	// safty features
	ABS                    *bool `json:"abs"`
	ESP                    *bool `json:"esp"`
	TractionControl        *bool `json:"traction_control"`
	EmergencyBrakeAssist   *bool `json:"emergency_brake_assist"`
	BlindSpotAssist        *bool `json:"blind_spot_assist"`
	LaneAssist             *bool `json:"lane_assist"`
	TrafficSignRecognition *bool `json:"traffic_sign_recognition"`
	ISOFIX                 *bool `json:"isofix"`

	// comfort features
	HeatedSteeringWheel *bool `json:"heated_steering_wheel"`
	StartStopSystem     *bool `json:"start_stop_system"`
	HeatedSeats         *bool `json:"heated_seats"`
	ElectricSeats       *bool `json:"electric_seats"`
	SportSeats          *bool `json:"sport_seats"`

	// Exterior Features
	FogLights          *bool `json:"fog_lights"`
	AdaptiveHeadlights *bool `json:"adaptive_headlights"`
	RainSensor         *bool `json:"rain_sensor"`

	// Infotainment
	Radio            *bool `json:"radio"`
	NavigationSystem *bool `json:"navigation_system"`
	VoiceControl     *bool `json:"voice_control"`
	Bluetooth        *bool `json:"bluetooth"`
	USB              *bool `json:"usb"`
	AppleCarPlay     *bool `json:"apple_carplay"`
	AndroidAuto      *bool `json:"android_auto"`

	// Meta
	Page      int    `json:"page"`
	Limit     int    `json:"limit"`
	SortBy    string `json:"sort_by"`    // price, registration, mileage
	SortOrder string `json:"sort_order"` // asc, desc
}
