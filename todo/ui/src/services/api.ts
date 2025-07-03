import axios from "axios";

export const api = axios.create({
  baseURL: "http://localhost:8080/api/v1",
});

const requests = {
  getTodos: (offset: number = 0, limit: number = 10) => api.get(`/todos?offset=${offset}&limit=${limit}`),
  getTodo: (id: number) => api.get(`/todos/${id}`),
  createTodo: (data: { title: string }) => api.post("/todos", data),
  updateTodo: (id: number, data: { title: string, completed: boolean }) => api.put(`/todos/${id}`, data),
  deleteTodo: (id: number) => api.delete(`/todos/${id}`),
}

export default requests;