"use client";
import axios from "axios";
import { CompareVehiclesRequestBody } from "./@types/request-body";

const Baseurl = "http://localhost:8080/api/v1";
export async function CompareVehicles(query: CompareVehiclesRequestBody) {
  // TODO: call api to get data
  const { data } = await axios.post(`${Baseurl}/vehicles/search`, {
    ...query,
    page: query.page || 5,
    page_size: query.page_size || 10,
  });
  return data;
}
