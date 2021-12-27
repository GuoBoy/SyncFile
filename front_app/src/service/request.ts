import axios, { AxiosResponse } from "axios";
import { ElLoadingService, ElMessage } from "element-plus";
import { BaseGetApi, BasePostApi } from "./api";

const instance = axios.create({
  baseURL: "http://localhost:3748",
  timeout: 9000,
});

instance.interceptors.request.use(
  (config) => {
    const token: string = window.localStorage.getItem("token") || "";
    if (token && config.headers) {
      config.headers["accessToken"] = token;
    }
    ElLoadingService();
    return config;
  },
  (err) => {
    Promise.reject(err);
    ElLoadingService().close();
    ElMessage.error("请求失败：" + err);
  }
);

instance.interceptors.response.use(
  (resp) => {
    ElLoadingService().close();
    return resp.data;
  },
  (err) => {
    Promise.reject(err);
    ElLoadingService().close();
    ElMessage.error("请求失败：" + err);
  }
);

export function get(api: BaseGetApi): Promise<AxiosResponse> {
  return new Promise((resolve, reject) => {
    if (api.params !== undefined) {
      instance
        .get(api.url, {
          params: api.params,
        })
        .then((res) => resolve(res))
        .catch((err) => reject(err));
    } else {
      instance
        .get(api.url)
        .then((res) => resolve(res))
        .catch((err) => reject(err));
    }
  });
}

export function post(api: BasePostApi): Promise<AxiosResponse> {
  return new Promise((resolve, reject) => {
    if (api.data !== undefined) {
      instance
        .post(api.url, api.data)
        .then((res) => resolve(res))
        .catch((err) => reject(err));
    } else {
      instance
        .post(api.url)
        .then((res) => resolve(res))
        .catch((err) => reject(err));
    }
  });
}
