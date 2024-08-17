import { defineStore } from "pinia";
import axios from "axios";

// Constants
const API_BASE_URL = "/api/users";

// Function to generate API URLs
function generateApiUrl(endpoint = "") {
  return `${API_BASE_URL}/${endpoint}`;
}

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
        const response = await axios.get(generateApiUrl());
        this.users = response.data;
      });
    },
    async fetchUser(userId) {
      await apiHandler(this, async () => {
        const response = await axios.get(generateApiUrl(userId));
        this.user = response.data;
      });
    },
    async createUser(newUser) {
      await apiHandler(this, async () => {
        const response = await axios.post(generateApiUrl(), newUser);
        this.users.push(response.data);
      });
    },
    async updateUser(userId, updatedUser) {
      await apiHandler(this, async () => {
        const response = await axios.put(generateApiUrl(userId), updatedUser);
        this.users = this.users.map((user) =>
          user.id === userId ? response.data : user
        );
      });
    },
    async deleteUser(userId) {
      await apiHandler(this, async () => {
        await axios.delete(generateApiUrl(userId));
        this.users = this.users.filter((user) => user.id !== userId);
      });
    },
  },
});
