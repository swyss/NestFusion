<template>
  <q-page padding>
    <h1>Weather Information</h1>
    <WeatherComponent city="Basel" />
  </q-page>
</template>

<script>
import WeatherComponent from 'src/components/WeatherComponent.vue';

export default {
  components: {
    WeatherComponent,
  },
};
</script>

<style scoped>
</style>
