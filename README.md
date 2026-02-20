# Vehicle Model Schema

This document provides an overview of the `Vehicle` model schema used in the `core` package.

## Overview

The `Vehicle` struct represents a vehicle entity with a variety of attributes, including identifiers, technical specifications, location, pricing, ownership, and safety features. This schema is used for managing and querying vehicle listings in a database.

## Schema Diagram

## Schema Diagram

+-------------------------+
| Vehicle |
+-------------------------+
| ID |
| CreatedAt |
| UpdatedAt |
| DeletedAt |
+-------------------------+
| VIN |
| ExternalID |
| SourceURL |
| Slug |
| Title |
+-------------------------+
| Make |
| Model |
| ModelVersion |
| VehicleClass |
| Category |
| Condition |
+-------------------------+
| City |
| ZipCode |
| Country |
+-------------------------+
| Price |
| DiscountPrice |
| Currency |
+-------------------------+
| Year |
| FirstRegistration |
| Mileage |
+-------------------------+
| FuelType |
| Transmission |
| EngineType |
| PowerHP |
| PowerKW |
| Displacement |
| Doors |
| Seats |
+-------------------------+
| ExteriorColor |
| InteriorColor |
| InteriorMaterial |
+-------------------------+
| PreviousOwners |
+-------------------------+
| CO2Emission |
| EmissionClass |
+-------------------------+
| ABS |
| ESP |
| TractionControl |
| EmergencyBrakeAssist |
| BlindSpotAssist |
| LaneAssist |
| TrafficSignRecognition |
| ISOFIX |
+-------------------------+
| HeatedSteeringWheel |
| StartStopSystem |
| HeatedSeats |
| ElectricSeats |
| SportSeats |
+-------------------------+
| FogLights |
| AdaptiveHeadlights |
| RainSensor |
+-------------------------+
| Radio |
| NavigationSystem |
| VoiceControl |
| Bluetooth |
| USB |
| AppleCarPlay |
| AndroidAuto |
+-------------------------+
| SellerType |
| SellerName |
| SellerCity |
| SellerCountry |
+-------------------------+
| Images |
+-------------------------+
| ListingStatus |
+-------------------------+

## Field Descriptions

### Identifiers

- **ID**: The unique identifier for the vehicle.
- **VIN**: Vehicle Identification Number (unique for each vehicle).
- **ExternalID**: An external identifier for the vehicle from a third-party source.
- **SourceURL**: The URL from where the vehicle data originates.
- **Slug**: A URL-friendly identifier for the vehicle.
- **Title**: The title or name of the vehicle listing.

### Make & Model

- **Make**: Manufacturer or brand of the vehicle.
- **Model**: The specific model of the vehicle.
- **ModelVersion**: Version or trim level of the vehicle.
- **VehicleClass**: The classification of the vehicle (e.g., sedan, SUV).
- **Category**: The vehicle's category (e.g., new, used).
- **Condition**: The condition of the vehicle (e.g., excellent, good, damaged).

### Location

- **City**: The city where the vehicle is located.
- **ZipCode**: The zip code of the vehicle's location.
- **Country**: The country where the vehicle is located.

### Pricing

- **Price**: The listed price of the vehicle.
- **DiscountPrice**: The discounted price if applicable.
- **Currency**: The currency in which the price is listed (default is EUR).

### Registration & Mileage

- **Year**: The manufacturing year of the vehicle.
- **FirstRegistration**: The date when the vehicle was first registered.
- **Mileage**: The total mileage the vehicle has been driven.

### Technical Specifications

- **FuelType**: Type of fuel used by the vehicle (e.g., petrol, diesel, electric).
- **Transmission**: The type of transmission (e.g., manual, automatic).
- **EngineType**: Type of engine used in the vehicle (e.g., inline, V6).
- **PowerHP**: The power of the engine in horsepower.
- **PowerKW**: The power of the engine in kilowatts.
- **Displacement**: The engine displacement in liters.
- **Doors**: The number of doors on the vehicle.
- **Seats**: The number of seats in the vehicle.

### Colors & Interior

- **ExteriorColor**: The color of the vehicle's exterior.
- **InteriorColor**: The color of the vehicle's interior.
- **InteriorMaterial**: The material used for the vehicle's interior (e.g., leather, fabric).

### Ownership

- **PreviousOwners**: The number of previous owners of the vehicle.

### Emissions

- **CO2Emission**: The carbon dioxide emissions of the vehicle in grams per kilometer.
- **EmissionClass**: The emission class of the vehicle (e.g., Euro 6).

### Safety Features

- **ABS**: Anti-lock Braking System (boolean).
- **ESP**: Electronic Stability Program (boolean).
- **TractionControl**: Traction control system (boolean).
- **EmergencyBrakeAssist**: Emergency brake assist (boolean).
- **BlindSpotAssist**: Blind spot assist (boolean).
- **LaneAssist**: Lane assist system (boolean).
- **TrafficSignRecognition**: Traffic sign recognition (boolean).
- **ISOFIX**: ISOFIX child seat connectors (boolean).

### Comfort

- **HeatedSteeringWheel**: Heated steering wheel feature (boolean).
- **StartStopSystem**: Start-stop system for fuel efficiency (boolean).
- **HeatedSeats**: Heated seats feature (boolean).
- **ElectricSeats**: Electrically adjustable seats (boolean).
- **SportSeats**: Sport seats (boolean).

### Exterior Features

- **FogLights**: Fog lights feature (boolean).
- **AdaptiveHeadlights**: Adaptive headlights (boolean).
- **RainSensor**: Rain sensor for wipers (boolean).

### Infotainment

- **Radio**: Radio system (boolean).
- **NavigationSystem**: Navigation system (boolean).
- **VoiceControl**: Voice control for infotainment (boolean).
- **Bluetooth**: Bluetooth connectivity (boolean).
- **USB**: USB connectivity (boolean).
- **AppleCarPlay**: Apple CarPlay compatibility (boolean).
- **AndroidAuto**: Android Auto compatibility (boolean).

### Seller

- **SellerType**: Type of seller (e.g., dealer, private seller).
- **SellerName**: Name of the seller.
- **SellerCity**: City of the seller.
- **SellerCountry**: Country of the seller.

### Media

- **Images**: A list of URLs to images of the vehicle.

### Listing

- **ListingStatus**: The status of the vehicle listing (e.g., active, sold, archived).

## Example JSON

```json
{
  "id": 1,
  "vin": "1HGBH41JXMN109186",
  "external_id": "EX12345",
  "source_url": "https://example.com/vehicle/1",
  "slug": "2020-toyota-corolla",
  "title": "2020 Toyota Corolla",
  "make": "Toyota",
  "model": "Corolla",
  "model_version": "SE",
  "vehicle_class": "Sedan",
  "category": "Used",
  "condition": "Excellent",
  "city": "Berlin",
  "zip_code": "10115",
  "country": "Germany",
  "price": 20000,
  "discount_price": 18000,
  "currency": "EUR",
  "year": 2020,
  "first_registration": "2020-05-01T00:00:00Z",
  "mileage": 50000,
  "fuel_type": "Petrol",
  "transmission": "Automatic",
  "engine_type": "Inline 4",
  "power_hp": 150,
  "power_kw": 110,
  "displacement": 2.0,
  "doors": 4,
  "seats": 5,
  "exterior_color": "Silver",
  "interior_color": "Black",
  "interior_material": "Leather",
  "previous_owners": 1,
  "co2_emission": 120,
  "emission_class": "Euro 6",
  "abs": true,
  "esp": true,
  "traction_control": true,
  "emergency_brake_assist": true,
  "blind_spot_assist": true,
  "lane_assist": true,
  "traffic_sign_recognition": true,
  "isofix": true,
  "heated_steering_wheel": true,
  "start_stop_system": true,
  "heated_seats": true,
  "electric_seats": true,
  "sport_seats": true,
  "fog_lights": true,
  "adaptive_headlights": true,
  "rain_sensor": true,
  "radio": true,
  "navigation_system": true,
  "voice_control": true,
  "bluetooth": true,
  "usb": true,
  "apple_carplay": true,
  "android_auto": true,
  "seller_type": "Dealer",
  "seller_name": "CarDealer GmbH",
  "seller_city": "Berlin",
  "seller_country": "Germany",
  "images": [
    "https://example.com/vehicle-image1.jpg",
    "https://example.com/vehicle-image2.jpg"
  ],
  "listing_status": "active"
}
```
