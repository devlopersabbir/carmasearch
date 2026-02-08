package core

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

// StringArray is a custom type for handling array of strings in DB (e.g. JSON)
type StringArray []string

func (a *StringArray) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, a)
}

func (a StringArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

type Vehicle struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	// Vehicle Identifiers
	Make             string `json:"make" gorm:"index:idx_make_model"`
	Model            string `json:"model" gorm:"index:idx_make_model"`
	ModelDescription string `json:"model_description"`
	ModelVersion     string `json:"model_verison" gorm:"index"`
	VIN              string `json:"vin" gorm:"uniqueIndex"`

	// Basic Info
	Title        string `json:"title"`
	Slug         string `json:"slug" gorm:"uniqueIndex"` // need to be unique
	VehicleClass string `json:"vehicle_class"`           // e.g., Car, Motorbike
	Category     string `json:"category"`                // e.g., Cabrio, Limousine
	Condition    string `json:"condition"`               // New, Used

	// location
	City    string `json:"city" gorm:"index"`
	ZipCode string `json:"zip_code" gorm:"index"`
	Country string `json:"country" gorm:"index"`

	// Pricing
	Price            float64 `json:"price" gorm:"index"`
	DiscountPrice    float64 `json:"discountPrice" gorm:"index"`
	Currency         string  `json:"currency" gorm:"default:'EUR'"`
	PriceRatingLabel string  `json:"price_rating_label"`

	// Core Attributes
	Year              int       `json:"year" gorm:"index"`
	FirstRegistration time.Time `json:"first_registration"`
	Mileage           int       `json:"mileage" gorm:"index"`
	EngineType        string    `json:"engine_type"` // e.g., Hybrid, Diesel, Petrol, Electric
	Transmission      string    `json:"transmission"`
	BodyType          string    `json:"body_type"`
	Displacement      int       `json:"displacement"` // in ccm
	PowerHP           int       `json:"power_hp"`
	PowerKW           int       `json:"power_kw"`
	FuelType          string    `json:"fuel_type"` // e.g., Diesel, Petrol, Electric
	Color             string    `json:"color"`
	Doors             int       `json:"doors"`
	Seats             int       `json:"seats"`
	Gearbox           string    `json:"gearbox"` // manual, automatic

	// Features & History
	Features               StringArray `json:"features" gorm:"type:jsonb"` // JSON array of features
	DamageUnrepaired       bool        `json:"damage_unrepaired"`
	Roadworthy             bool        `json:"roadworthy"`
	AccidentDamaged        bool        `json:"accident_damaged"`
	NumberOfPreviousOwners int         `json:"number_of_previous_owners"`
	Warranty               bool        `json:"warranty"`

	// Emissions & Consumption
	EmissionClass       string  `json:"emission_class"` // Euro 6d
	CO2Emission         int     `json:"co2_emission"`   // g/km
	ConsumptionCombined float64 `json:"consumption_combined"`
	ConsumptionCity     float64 `json:"consumption_city"`
	ConsumptionHighway  float64 `json:"consumption_highway"`

	// Seller
	SellerType    string `json:"seller_type"` // dealer | private
	SellerName    string `json:"seller_name"`
	SellerCity    string `json:"seller_city"`
	SellerCountry string `json:"seller_country"`

	// Media
	Images StringArray `json:"images" gorm:"type:jsonb"` // JSON array of image URLs

	// Metadata
	ExternalID string `json:"external_id" gorm:"index"` // inner_id from source
	SourceURL  string `json:"source_url"`

	ListingStatus string `json:"listing_status"` // active, sold, reserved
}
