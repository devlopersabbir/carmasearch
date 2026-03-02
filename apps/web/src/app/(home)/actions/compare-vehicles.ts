"use client";
import axios from "axios";

const Baseurl = "http://localhost:8080/api/v1";
export async function CompareVehicles<T extends object>(query: T) {
  console.log("query", query);
  // TODO: call api to get data
  const data = await axios.post(`${Baseurl}/vehicles/search`, query);
  console.log("data", data);
  return data;
}
