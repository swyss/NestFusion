import { defineStore } from "pinia";
import { api } from "boot/axios"; // Verwende die konfigurierte API-Instanz

// Constants
const API_BASE_URL = "/users"; // baseURL ist bereits in der api Instanz konfiguriert

// Extracted async API handler
async function apiHandler(context, action) {
  context.loading = true;
  try {
    await action();
  } catch (error) {
    context.error = error.response?.data ?? "Unexpected error";
  } finally {
    context.loading = false;
  }
}

// Pinia Store
export const useUserStore = defineStore("userStore", {
  state: () => ({
    users: [],
    user: null,
    loading: false,
    error: null,
  }),
  actions: {
    async fetchUsers() {
      await apiHandler(this, async () => {
        const response = await api.get(API_BASE_URL);
        this.users = response.data;
      });
    },
    async fetchUser(userId) {
      await apiHandler(this, async () => {
        const response = await api.get(`${API_BASE_URL}/${userId}`);
        this.user = response.data;
      });
    },
    async createUser(newUser) {
      await apiHandler(this, async () => {
        const response = await api.post(API_BASE_URL, newUser);
        this.users.push(response.data);
      });
    },
    async updateUser(userId, updatedUser) {
      await apiHandler(this, async () => {
        const response = await api.put(`${API_BASE_URL}/${userId}`, updatedUser);
        this.users = this.users.map((user) =>
          user.id === userId ? response.data : user
        );
      });
    },
    async deleteUser(userId) {
      await apiHandler(this, async () => {
        await api.delete(`${API_BASE_URL}/${userId}`);
        this.users = this.users.filter((user) => user.id !== userId);
      });
    },
  },
});
