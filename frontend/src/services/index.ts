import axios, { AxiosInstance } from "axios";

import { type Benchmark } from "@/types/benchmark";
import { type Project, ProjectSchema } from "@/types/project";
import { type User } from "@/types/user";

import { token } from "@/store/auth";

class ProjectService {
  constructor(
    protected httpClient: AxiosInstance,
    protected userService: UserService
  ) {}

  async get(id: string): Promise<Project> {
    const config = {
      headers: {
        Authorization: `Bearer ${token.value}`
      }
    };

    const data = await this.httpClient
      .get<Project>(`/projects/${id}`, config)
      .then(async (response) => {
        return ProjectSchema.parse(response.data);
      });

    return data;
  }

  async getAll() {
    const config = {
      headers: {
        Authorization: `Bearer ${token.value}`
      }
    };

    const data = await this.httpClient
      .get<Project[]>("/projects", config)
      .then((response) => response.data);

    return data;
  }

  async create(data: any) {
    const config = {
      headers: {
        Authorization: `Bearer ${token.value}`
      }
    };
    return this.httpClient.post("/projects", data, config);
  }

  async update(id: string, data: any) {
    const config = {
      headers: {
        Authorization: `Bearer ${token.value}`
      }
    };
    return this.httpClient
      .put(`/projects/${id}`, data, config)
      .then((response: any) => response.data);
  }

  async delete(id: string) {
    const config = {
      headers: {
        Authorization: `Bearer ${token.value}`
      }
    };
    return this.httpClient
      .delete(`/projects/${id}`, config)
      .then((response: any) => response.data);
  }

  async updateMetrics(id: string, data: any) {
    const config = {
      headers: {
        Authorization: `Bearer ${token.value}`
      }
    };
    return this.httpClient
      .put(`/projects/${id}/metrics`, data, config)
      .then((response: any) => response.data);
  }

  async markAsCompleted(id: string, date: Date) {
    const config = {
      headers: {
        Authorization: `Bearer ${token.value}`
      }
    };

    return this.httpClient
      .patch(`/projects/${id}/completed`, { completionDate: date }, config)
      .then((response: any) => response.data);
  }
}

class BenchmarkService {
  constructor(protected httpClient: AxiosInstance) {}

  async get(id: string) {
    const config = {
      headers: {
        Authorization: `Bearer ${token.value}`
      }
    };
    return this.httpClient
      .get(`/benchmarks/${id}`, config)
      .then(async (response: any) => {
        return response.data;
      });
  }

  async getAll() {
    const config = {
      headers: {
        Authorization: `Bearer ${token.value}`
      }
    };
    return await this.httpClient
      .get<Benchmark[]>("/benchmarks", config)
      .then((response) => response.data);
  }

  async create(data: any) {
    const config = {
      headers: {
        Authorization: `Bearer ${token.value}`
      }
    };
    return this.httpClient.post("/benchmarks", data, config);
  }

  async update(id: string, data: any) {
    const config = {
      headers: {
        Authorization: `Bearer ${token.value}`
      }
    };
    return this.httpClient
      .put(`/benchmarks/${id}`, data, config)
      .then((response: any) => response.data);
  }

  async delete(id: string) {
    const config = {
      headers: {
        Authorization: `Bearer ${token.value}`
      }
    };
    return this.httpClient
      .delete(`/benchmarks/${id}`, config)
      .then((response: any) => response.data);
  }
}

class UserService {
  constructor(protected httpClient: AxiosInstance) {}

  async get(id: string) {
    const config = {
      headers: {
        Authorization: `Bearer ${token.value}`
      }
    };
    return this.httpClient
      .get<User>(`/users/${id}`, config)
      .then(async (response) => {
        return response.data;
      });
  }

  async getAll() {
    const config = {
      headers: {
        Authorization: `Bearer ${token.value}`
      }
    };
    return this.httpClient
      .get<User[]>("/users", config)
      .then((response) => response.data);
  }

  async create(data: User) {
    const config = {
      headers: {
        Authorization: `Bearer ${token.value}`
      }
    };
    return this.httpClient.post<User>("/users", data, config);
  }

  async uploadAvatar(file: File) {
    const config = {
      headers: {
        Authorization: `Bearer ${token.value}`,
        "Content-Type": "multipart/form-data"
      }
    };

    const formData = new FormData();
    formData.append("avatar", file);

    return this.httpClient
      .post<{ imageUrl: string }>(`/users/avatar`, formData, config)
      .then((response) => response.data);
  }

  async update(id: string, data: User) {
    const config = {
      headers: {
        Authorization: `Bearer ${token.value}`
      }
    };
    return this.httpClient
      .put<User>(`/users/${id}`, data, config)
      .then((response) => response.data);
  }

  async delete(id: string) {
    const config = {
      headers: {
        Authorization: `Bearer ${token.value}`
      }
    };
    return this.httpClient
      .delete(`/users/${id}`, config)
      .then((response: any) => response.data);
  }

  async me() {
    const config = {
      headers: {
        Authorization: `Bearer ${token.value}`
      }
    };

    return this.httpClient
      .get<User>("/users/me", config)
      .then((response) => response.data);
  }
}

const httpClient = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL as string
});

httpClient.interceptors.request.use(
  async (config) => {
    const token = localStorage.getItem("token");
    if (token) {
      config.headers["Authorization"] = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    console.info("interceptor error", error);
    return Promise.reject(error);
  }
);

class PreferenceService {
  constructor(protected httpClient: AxiosInstance) {}

  async getOrganisations(): Promise<string[]> {
    const config = {
      headers: {
        Authorization: `Bearer ${token.value}`
      }
    };

    return this.httpClient
      .get<string[]>(`/preferences/organisations`, config)
      .then((response) => response.data);
  }
}

const users = new UserService(httpClient);
const projects = new ProjectService(httpClient, users);
const benchmarks = new BenchmarkService(httpClient);
const preferences = new PreferenceService(httpClient);

export default {
  projects,
  benchmarks,
  users,
  preferences
};
