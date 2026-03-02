import axios from "axios";

export const Baseurl = "http://localhost:8080/api/v1/";
export const api = {
  async post<T extends object, R extends object>(url: string, data: T) {
    const res = await axios.post<R>(`${Baseurl}${url}`, data, {
      headers: {
        "Content-Type": "application/json",
      },
    });
    return res.data;
  },
  async get<T extends object, R extends object>(url: string, params?: T) {
    const res = await axios.get<R>(`${Baseurl}${url}`, {
      headers: {
        "Content-Type": "application/json",
      },
      params,
    });
    return res.data;
  },
};
