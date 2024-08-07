import { defineStore } from 'pinia'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    User: null,
    loggedIn: true // change to false
  }),

  getters: {
    isLoggedIn (state) {
      return state.loggedIn
    },
    user (state) {
      return state.User
    }
  },

  actions: {
    login (username, password) {
      // make the login request
      // if successful, set the user and loggedIn
      this.User = { username }
      this.loggedIn = true
    },
    logout () {
      // make the logout request
      // if successful, set the user and loggedIn
      this.User = null
      this.loggedIn = false
    },
    async register (username, password) {
      // make the request to register the user
      // if successful, set the user and loggedIn
      this.User = {username}
      this.loggedIn = true
    }
  }
})

