import { defineStore } from 'pinia'

export const useAppInfoStore = defineStore('counter', {
  state: () => ({
    counter: 0
  }),

  getters: {
    doubleCount (state) {
      return state.counter * 2
    }
  },

  actions: {
    increment () {
      this.counter++
    }
  }
})
