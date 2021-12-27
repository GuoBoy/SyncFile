import { UserInfo } from "@/common/types";
import { post } from "./request";

export interface BasePostApi {
  url: string;
  data?: any;
}

export interface BaseGetApi {
  url: string;
  params?: any;
}

export function login(data: any, callback: any): void {
  post({ url: "/login", data: data })
    .then((res) => callback(res))
    .catch((err) => callback(err));
}

export function register(data: any, callback: any): void {
  post({ url: "/register", data: data })
    .then((res) => callback(res))
    .catch((err) => {
      callback(err);
    });
}
