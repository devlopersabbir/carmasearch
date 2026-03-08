"use client";
import axios from "axios";
import {
  compareVehiclesRequestBody,
  CompareVehiclesRequestBody,
} from "./@types/request-body";

const Baseurl = "http://localhost:8080/api/v1";
export async function CompareVehicles<T extends CompareVehiclesRequestBody>(
  query: T,
) {
  // safe perse with zod
  const parsedQuery = compareVehiclesRequestBody.safeParse(query);

  if (!parsedQuery.success) {
    throw new Error("Invalid query");
  }
  // TODO: call api to get data
  const { data } = await axios.post(`${Baseurl}/vehicles/search`, {
    ...query,
    page: query.page || 5,
    page_size: query.page_size || 10,
  });
  return data;
}
