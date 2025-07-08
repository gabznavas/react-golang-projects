import axios from "axios";

const API_URL = "http://localhost:8080/api/v1";

const axiosInstance = axios.create({
  baseURL: API_URL,
});


export const requests = {
  getProjects: async () => {
    const response = await axiosInstance.get("/project");
    return response.data;
  },
  getProjectById: async (id: string) => {
    const response = await axiosInstance.get(`/project/${id}`);
    return response.data;
  },
  createProject: async (payload: { name: string, description: string }) => {
    const response = await axiosInstance.post("/project", payload);
    return response.data;
  },
  updateProject: async (id: string, payload: { name: string, description: string }) => {
    const response = await axiosInstance.put(`/project/${id}`, payload);
    return response.data;
  },
  deleteProject: async (id: string) => {
    const response = await axiosInstance.delete(`/project/${id}`);
    return response.data;
  },
};