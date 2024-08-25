import {defineStore} from "pinia";
import {api} from "boot/axios"; // Use the configured API instance

// Constants
const API_BASE_URL = "/users"; // baseURL is already configured in the API instance

// Extracted async API handler
async function apiHandler(context, action) {
  context.loading = true; // Set loading state to true before the action
  try {
    await action(); // Execute the provided async action
  } catch (error) {
    context.error = error.response?.data ?? "Unexpected error"; // Capture and store any errors
  } finally {
    context.loading = false; // Ensure loading state is reset after the action completes
  }
}

// Pinia Store
export const useUserStore = defineStore("userStore", {
  state: () => ({
    users: [], // List of users
    user: null, // Single user details
    loading: false, // Loading state indicator
    error: null, // Error state indicator
  }),
  actions: {
    async fetchUsers() {
      await apiHandler(this, async () => {
        const response = await api.get(API_BASE_URL); // Fetch all users from the API
        this.users = response.data; // Store the fetched users in the state
      });
    },
    async fetchUser(userId) {
      await apiHandler(this, async () => {
        const response = await api.get(`${API_BASE_URL}/${userId}`); // Fetch a single user by ID from the API
        this.user = response.data; // Store the fetched user details in the state
      });
    },
    async createUser(newUser) {
      await apiHandler(this, async () => {
        const response = await api.post(API_BASE_URL, newUser); // Create a new user via the API
        this.users.push(response.data); // Add the newly created user to the users list in the state
      });
    },
    async updateUser(userId, updatedUser) {
      await apiHandler(this, async () => {
        const response = await api.put(`${API_BASE_URL}/${userId}`, updatedUser); // Update a user by ID via the API
        this.users = this.users.map((user) =>
          user.id === userId ? response.data : user
        ); // Replace the updated user in the users list
      });
    },
    async deleteUser(userId) {
      await apiHandler(this, async () => {
        await api.delete(`${API_BASE_URL}/${userId}`); // Delete a user by ID via the API
        this.users = this.users.filter((user) => user.id !== userId); // Remove the deleted user from the users list
      });
    },
  },
});
