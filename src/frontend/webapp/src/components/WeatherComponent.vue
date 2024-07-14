<template>
  <div v-if="weatherStore.weather">
    <h2>Weather in {{ city }}</h2>
    <p>Temperature: {{ weatherStore.weather.main.temp }} °C</p>
    <p>Weather: {{ weatherStore.weather.weather[0].description }}</p>
  </div>
  <div v-else-if="weatherStore.loading">Loading...</div>
  <div v-else-if="weatherStore.error">Error: {{ weatherStore.error.message }}</div>
  <div v-else>
    <p>No data available</p>
  </div>
</template>

<script>
import { useWeatherStore } from 'src/stores/weater-collector';
import { onMounted } from 'vue';

export default {
  name: 'WeatherComponent',
  props: {
    city: {
      type: String,
      required: true,
    },
  },
  setup(props) {
    const weatherStore = useWeatherStore();
    onMounted(() => {
      weatherStore.fetchWeather(props.city);
    });

    return { weatherStore };
  },
};
</script>

<style scoped>
</style>
