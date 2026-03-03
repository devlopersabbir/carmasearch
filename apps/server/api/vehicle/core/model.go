package core

import (
	"time"

	"github.com/carmasearch/carma-server/internal/utils"
	"gorm.io/gorm"
)

type Vehicle struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// ================= PRIMARY =================
	UniqueID   string `gorm:"primaryKey;column:unique_id;type:varchar(500)"`
	VehicleID  string `gorm:"column:vehicle_id;index"`
	DataSource string `gorm:"column:data_source;index"`
	ListingURL string `gorm:"column:listing_url;type:text;index"`

	ScrapedAt time.Time `gorm:"column:scraped_at;type:date;index"`

	IsVehicleAvailable bool `gorm:"column:is_vehicle_available;index"`

	Images utils.StringArray `gorm:"column:images;type:jsonb"`

	// ================= BASIC INFO =================
	Title        *string `gorm:"column:title"`
	Slug         *string `gorm:"column:slug"`
	Subtitle     *string `gorm:"column:subtitle"`
	Description  *string `gorm:"column:description"`
	VehicleTitle *string `gorm:"column:vehicle_title"`

	Make         *string `gorm:"column:make"`
	Model        *string `gorm:"column:model"`
	ModelVersion *string `gorm:"column:model_version"`
	ModelRange   *string `gorm:"column:model_range"`
	TrimLine     *string `gorm:"column:trim_line"`
	VehicleType  *string `gorm:"column:vehicle_type"`
	Category     *string `gorm:"column:category"`
	BodyType     *string `gorm:"column:body_type"`

	MakeID            *string `gorm:"column:make_id"`
	ModelID           *string `gorm:"column:model_id"`
	ModelGenerationID *string `gorm:"column:model_generation_id"`
	ModelVariantID    *string `gorm:"column:model_variant_id"`

	// ================= PRICING =================
	Price          *string `gorm:"column:price"`
	PriceInfo      *string `gorm:"column:price_info"`
	PriceLabel     *string `gorm:"column:price_label"`
	MonthlyPayment *string `gorm:"column:monthly_payment"`
	DownPayment    *string `gorm:"column:down_payment"`
	NewCarPrice    *string `gorm:"column:newCarPrice"`

	// ================= ENGINE =================
	EngineType      *string `gorm:"column:engine_type"`
	MotorTypeID     *string `gorm:"column:motor_type_id"`
	PowerKW         *string `gorm:"column:power_kw"`
	PowerHP         *string `gorm:"column:power_hp"`
	PowerDisplay    *string `gorm:"column:power_display"`
	DisplacementCCM *string `gorm:"column:displacement_ccm"`
	Cylinders       *string `gorm:"column:cylinders"`
	Gears           *string `gorm:"column:gears"`
	Torque          *string `gorm:"column:torque"`
	TopSpeed        *string `gorm:"column:topSpeed"`
	Acceleration    *string `gorm:"column:acceleration"`

	// ================= FUEL & TRANSMISSION =================
	FuelType         *string `gorm:"column:fuel_type"`
	FuelCategory     *string `gorm:"column:fuel_category"`
	PrimaryFuel      *string `gorm:"column:primary_fuel"`
	Transmission     *string `gorm:"column:transmission"`
	TransmissionType *string `gorm:"column:transmission_type"`
	DriveTrain       *string `gorm:"column:drive_train"`
	AllWheelDrive    *bool   `gorm:"column:all_wheel_drive"`
	FrontWheelDrive  *bool   `gorm:"column:front_wheel_drive"`
	RearWheelDrive   *bool   `gorm:"column:rear_wheel_drive"`

	// ================= MILEAGE & HISTORY =================
	MileageKM          *string `gorm:"column:mileage_km"`
	MileageDisplay     *string `gorm:"column:mileage_display"`
	FirstRegistration  *string `gorm:"column:first_registration"`
	ProductionYear     *string `gorm:"column:production_year"`
	ConstructionYear   *string `gorm:"column:construction_year"`
	LastInspection     *string `gorm:"column:last_inspection"`
	NextInspection     *string `gorm:"column:next_inspection"`
	PreviousOwners     *string `gorm:"column:previous_owners"`
	FullServiceHistory *string `gorm:"column:full_service_history"`

	// ================= DIMENSIONS =================
	TotalLength        *string `gorm:"column:total_length"`
	TotalWidth         *string `gorm:"column:total_width"`
	TotalHeight        *string `gorm:"column:total_height"`
	WheelBase          *string `gorm:"column:wheel_base"`
	GrossVehicleWeight *string `gorm:"column:gross_vehicle_weight"`
	Payload            *string `gorm:"column:payload"`
	FuelTankVolume     *string `gorm:"column:fuel_tank_volume"`

	// ================= ELECTRIC =================
	BatteryCapacity  *string `gorm:"column:battery_capacity"`
	BatteryType      *string `gorm:"column:battery_type"`
	ElectricRange    *string `gorm:"column:electric_range"`
	MaxChargingPower *string `gorm:"column:max_charging_power"`
	ChargingTime80   *string `gorm:"column:charging_time_80"`

	// ================= LOCATION =================
	CountryCode *string `gorm:"column:country_code"`
	City        *string `gorm:"column:city"`
	PostalCode  *string `gorm:"column:postal_code"`
	Street      *string `gorm:"column:street"`
	SellerName  *string `gorm:"column:seller_name"`

	// ================= SAFETY FLAGS =================
	ABS                    *bool `gorm:"column:abs"`
	ESP                    *bool `gorm:"column:esp"`
	TractionControl        *bool `gorm:"column:traction_control"`
	DriverAirbag           *bool `gorm:"column:driver_airbag"`
	PassengerAirbag        *bool `gorm:"column:passenger_airbag"`
	SideAirbag             *bool `gorm:"column:side_airbag"`
	HeadAirbag             *bool `gorm:"column:head_airbag"`
	BlindSpotAssist        *bool `gorm:"column:blind_spot_assist"`
	LaneAssist             *bool `gorm:"column:lane_assist"`
	CollisionWarningSystem *bool `gorm:"column:collision_warning_system"`
	EmergencyBrakeAssist   *bool `gorm:"column:emergency_brake_assist"`

	// ================= COMFORT =================
	AirConditioning       *bool `gorm:"column:air_conditioning"`
	ClimateControl        *bool `gorm:"column:climate_control"`
	HeatedSeats           *bool `gorm:"column:heated_seats"`
	ElectricSeats         *bool `gorm:"column:electric_seats"`
	CruiseControl         *bool `gorm:"column:cruise_control"`
	AdaptiveCruiseControl *bool `gorm:"column:adaptive_cruise_control"`
	ParkingSensorsFront   *bool `gorm:"column:parking_sensors_front"`
	ParkingSensorsRear    *bool `gorm:"column:parking_sensors_rear"`
	ParkingCamera         *bool `gorm:"column:parking_camera"`

	// ================= INFOTAINMENT =================
	NavigationSystem *bool `gorm:"column:navigation_system"`
	Bluetooth        *bool `gorm:"column:bluetooth"`
	AppleCarplay     *bool `gorm:"column:apple_carplay"`
	AndroidAuto      *bool `gorm:"column:android_auto"`
	DigitalCockpit   *bool `gorm:"column:digital_cockpit"`
	SoundSystem      *bool `gorm:"column:sound_system"`

	// ================= RAW / TRACKING =================
	Enriched        *string           `gorm:"column:enriched"`
	ExtraAttributes utils.StringArray `gorm:"-"`
}
