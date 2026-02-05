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

func (a *StringArray) Scan(value interface{}) error {
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
	VIN              string `json:"vin" gorm:"uniqueIndex"`

	// Basic Info
	Title            string  `json:"title"`
	VehicleClass     string  `json:"vehicle_class"` // e.g., Car, Motorbike
	Category         string  `json:"category"`      // e.g., Cabrio, Limousine
	Condition        string  `json:"condition"`     // New, Used
	Location         string  `json:"location"`
	Price            float64 `json:"price" gorm:"index"`
	Currency         string  `json:"currency" gorm:"default:'EUR'"`
	PriceRatingLabel string  `json:"price_rating_label"`

	// Core Attributes
	Year              int       `json:"year" gorm:"index"`
	FirstRegistration time.Time `json:"first_registration"`
	Mileage           int       `json:"mileage" gorm:"index"`
	EngineType        string    `json:"engine_type"`
	Transmission      string    `json:"transmission"`
	BodyType          string    `json:"body_type"`
	Displacement      int       `json:"displacement"` // in ccm
	PowerHP           int       `json:"power_hp"`
	PowerKW           int       `json:"power_kw"`
	FuelType          string    `json:"fuel_type"`
	Color             string    `json:"color"`
	Doors             int       `json:"doors"`
	Seats             int       `json:"seats"`
	Gearbox           string    `json:"gearbox"` // manual, automatic

	// Features & History
	Features               StringArray `json:"features" gorm:"type:text"` // JSON array of features
	DamageUnrepaired       bool        `json:"damage_unrepaired"`
	Roadworthy             bool        `json:"roadworthy"`
	AccidentDamaged        bool        `json:"accident_damaged"`
	NumberOfPreviousOwners int         `json:"number_of_previous_owners"`
	Warranty               bool        `json:"warranty"`

	// Media
	Images StringArray `json:"images" gorm:"type:text"` // JSON array of image URLs

	// Metadata
	ExternalID string `json:"external_id" gorm:"index"` // inner_id from source
	SourceURL  string `json:"source_url"`
}
