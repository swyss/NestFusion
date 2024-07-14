import { defineStore } from 'pinia';
import { getWeather } from 'src/services/weatherService';

export const useWeatherStore = defineStore('weather', {
  state: () => ({
    weather: null,
    loading: false,
    error: null,
  }),
  actions: {
    async fetchWeather(city) {
      this.loading = true;
      try {
        this.weather = await getWeather(city);
      } catch (error) {
        this.error = error;
      } finally {
        this.loading = false;
      }
    },
  },
});
