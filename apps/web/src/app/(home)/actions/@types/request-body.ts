export type CompareVehiclesRequestBody = {
  listing_url: string;
  page?: number;
  page_size?: number;
  // for the advance search
  registration_from?: string;
  registration_until?: string;
  mileage_from?: string;
  mileage_until?: string;
  exterior_colors?: string[];
  interior_colors?: string[];
  interior_materials?: string[];
};
