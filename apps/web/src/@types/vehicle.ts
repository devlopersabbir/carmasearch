export type ResponseVehicleData = {
  total: number;
  page: number;
  page_size: number;
  query_vehicle: VehicleData;
  results: VehicleData[];
};

export type VehicleData = {
  unique_id: string;
  vehicle_id: string;
  data_source: string;
  listing_url: string;
  images?: any;

  created_at?: string;
  scraped_at: string;
  updated_at: string;

  is_vehicle_available: boolean;

  title?: string;
  price?: string;
  make?: string;
  model?: string;
  model_version?: string;
  model_range?: string;
  price_info?: string;
  subtitle?: string;
  price_label?: string;
  trim_line?: string;
  vehicle_type?: string;
  category?: string;
  body_type?: string;
  make_id?: string;
  model_id?: string;
  model_generation_id?: string;
  model_variant_id?: string;
  motor_type_id?: string;
  trim_line_id?: string;
  air_conditioning_type?: string;
  sku?: string;
  hsn_tsn?: string;
  identifier?: string;

  power_kw?: string;
  power_hp?: string;
  power_display?: string;
  power_kw_display?: string;
  power_hp_display?: string;

  displacement_ccm?: string;
  displacement_display?: string;
  cylinders?: string;
  gears?: string;
  weight?: string;
  net_weight?: string;

  fuel_type?: string;
  fuel_category?: string;
  primary_fuel?: string;
  transmission?: string;
  transmission_type?: string;
  drive_train?: string;
  has_particle_filter?: string;

  fuel_consumption_combined?: string;
  fuel_consumption_urban?: string;
  fuel_consumption_extra_urban?: string;

  co2_emission?: string;
  co2_emission_combined?: string;
  co2_emissions_combined_fallback?: string;
  co2_emissions_combined_weighted?: string;
  co2_emissions_discharged?: string;

  co2_class?: string;
  co2_class_discharged?: string;
  emission_sticker?: string;
  emission_standard?: string;

  consumption_combined?: string;
  consumption_combined_weighted?: string;
  consumption_combined_discharged?: string;
  consumption_electric_combined?: string;
  consumption_electric_combined_weighted?: string;

  consumption_city?: string;
  consumption_city_discharged?: string;
  consumption_suburban?: string;
  consumption_suburban_discharged?: string;
  consumption_rural?: string;
  consumption_rural_discharged?: string;
  consumption_highway?: string;
  consumption_highway_discharged?: string;

  consumption_electric_city?: string;
  consumption_electric_suburban?: string;
  consumption_electric_rural?: string;
  consumption_electric_highway?: string;

  environment_pkw_envkv?: string;
  environment_eu_directive?: string;
  environment_bimschv35?: string;
  wltp?: string;
  energy_consumption?: string;

  consumption_costs?: string;
  consumption_costs_year?: string;
  fuel_price?: string;

  co2_costs?: string;
  co2_costs_average?: string;
  co2_costs_high?: string;
  co2_costs_low?: string;

  vehicle_tax?: string;

  engine_type?: string;
  other_energy_source?: string;

  mileage_km?: number;
  mileage_display?: string;
  mileage_detail?: string;
  mileage_in_km?: string;

  first_registration?: string;
  first_registration_raw?: string;
  first_registration_date?: string;
  production_year?: string;
  construction_year?: string;

  last_inspection?: string;
  next_inspection?: string;
  last_service_date?: string;
  last_service_mileage?: string;
  last_belt_service?: string;

  full_service_history?: string;

  offer_type?: string;
  condition?: string;
  damage_condition?: string;
  had_accident?: string;
  previous_owners?: string;
  is_rental?: string;
  non_smoking?: string;
  has_registration?: string;
  new_driver_suitable?: string;

  country_version?: string;
  original_market?: string;
  type?: string;

  seats?: string;
  doors?: string;

  color?: string;
  color_original?: string;
  manufacturer_color?: string;
  paint_type?: string;

  upholstery?: string;
  upholstery_color?: string;
  interior?: string;
  interior_color?: string;
  exterior_color?: string;
  interior_type?: string;
};
