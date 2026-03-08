import { z } from "zod";

export const compareVehiclesRequestBody = z.object({
  listing_url: z.string(),
  page: z.number().optional(),
  page_size: z.number().optional(),
  registration_from: z.string().optional(),
  registration_until: z.string().optional(),
  mileage_from: z.string().optional(),
  mileage_until: z.string().optional(),
  exterior_colors: z.array(z.string()).optional(),
  interior_colors: z.array(z.string()).optional(),
  interior_materials: z.array(z.string()).optional(),
});

export type CompareVehiclesRequestBody = z.infer<
  typeof compareVehiclesRequestBody
>;
