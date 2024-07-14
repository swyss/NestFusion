<template>
  <q-page padding>
    <h1>Top News</h1>
    <NewsComponent country="ch" />
  </q-page>
</template>

<script>
import NewsComponent from 'src/components/NewsComponent.vue';

export default {
  components: {
    NewsComponent,
  },
};
</script>

<style scoped>
</style>
