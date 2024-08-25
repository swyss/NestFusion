import {defineStore} from 'pinia';
import {api} from 'boot/axios'; // Use the configured API instance

export const useAuthStore = defineStore('authStore', {
  state: () => ({
    user: null, // The authenticated user's data
    token: null, // The authentication token
    loading: false, // Loading state indicator
    error: null // Error state indicator
  }),
  getters: {
    isAuthenticated: (state) => !!state.token, // Check if the user is authenticated based on the presence of a token
    isAdmin: (state) => state.user && state.user.role === 'admin' // Check if the authenticated user is an admin
  },
  actions: {
    async login({email, password}) {
      this.loading = true; // Set loading state to true during the login process
      try {
        const response = await api.post('/api/login', {email, password}); // Attempt to login the user via the API
        this.token = response.data.token; // Store the received token in the state
        this.user = response.data.user; // Store the authenticated user data in the state
      } catch (error) {
        this.error = error.response?.data ?? 'Unexpected error'; // Capture and store any errors
      } finally {
        this.loading = false; // Reset loading state after the login process completes
      }
    },
    logout() {
      this.token = null; // Clear the token to log the user out
      this.user = null; // Clear the user data to log the user out
    },
    async fetchUser() {
      this.loading = true; // Set loading state to true while fetching the user data
      try {
        const response = await api.get('/api/user'); // Fetch the authenticated user's data from the API
        this.user = response.data; // Store the fetched user data in the state
      } catch (error) {
        this.error = error.response?.data ?? 'Unexpected error'; // Capture and store any errors
      } finally {
        this.loading = false; // Reset loading state after fetching the user data
      }
    },
    async updateUser(updatedUser) {  // Method to update the user profile
      this.loading = true; // Set loading state to true during the update process
      try {
        const response = await api.put(`/api/user/${this.user.id}`, updatedUser); // Update the user profile via the API
        this.user = response.data; // Store the updated user data in the state
        this.$q.notify({type: 'positive', message: 'Profile updated successfully!'}); // Display a success notification
      } catch (error) {
        this.error = error.response?.data ?? 'Unexpected error'; // Capture and store any errors
        this.$q.notify({type: 'negative', message: 'Failed to update profile!'}); // Display an error notification
      } finally {
        this.loading = false; // Reset loading state after the update process completes
      }
    }
  }
});
