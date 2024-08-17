import { defineStore } from "pinia";
import axios from "axios";

// Constants
const API_BASE_URL = "/api/users";

// Extracted async API handler
async function handleApiCall(context, action) {
  context.loading = true;
  try {
    await action();
  } catch (error) {
    context.error = error.response.data;
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
      await handleApiCall(this, async () => {
        const response = await axios.get(API_BASE_URL);
        this.users = response.data;
      });
    },
    async fetchUser(userId) {
      await handleApiCall(this, async () => {
        const response = await axios.get(`${API_BASE_URL}/${userId}`);
        this.user = response.data;
      });
    },
    async createUser(newUser) {
      await handleApiCall(this, async () => {
        const response = await axios.post(API_BASE_URL, newUser);
        this.users.push(response.data);
      });
    },
    async updateUser(userId, updatedUser) {
      await handleApiCall(this, async () => {
        const response = await axios.put(
          `${API_BASE_URL}/${userId}`,
          updatedUser
        );
        this.users = this.users.map((user) =>
          user.id === userId ? response.data : user
        );
      });
    },
    async deleteUser(userId) {
      await handleApiCall(this, async () => {
        await axios.delete(`${API_BASE_URL}/${userId}`);
        this.users = this.users.filter((user) => user.id !== userId);
      });
    },
  },
});
