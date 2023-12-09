import axios from "axios";
import { handleError, handleResponse } from "./response";

export const BACKEND_URL = process.env.REACT_APP_BACKEND_URL;

export const fetcher = ([path, data = {}]) => get(path, data);

export const getFileData = (file) => {
  const formData = new FormData();
  formData.append("file", file);
  return { data: formData, headers: { "Content-Type": "multipart/form" } };
};

export const post = (path, data, headers = {}) =>
  axios
    .post(`${BACKEND_URL}/${path}`, data, {
      headers: headers,
    })
    .then(handleResponse)
    .catch(handleError);
export const put = (path, data, headers = {}) =>
  axios
    .put(`${BACKEND_URL}/${path}`, data, {
      headers: headers,
    })
    .then(handleResponse)
    .catch(handleError);

export const patch = (path, data) =>
  axios
    .patch(`${BACKEND_URL}/${path}`, data)
    .then(handleResponse)
    .catch(handleError);

export const get = (path, data) => {
  return axios
    .get(`${BACKEND_URL}/${path}`, {
      params: data,
    })
    .then(handleResponse)
    .catch(handleError);
};

export const remove = (path) =>
  axios
    .delete(`${BACKEND_URL}/${path}`)
    .then(handleResponse)
    .catch(handleError);
