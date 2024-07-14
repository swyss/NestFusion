<template>
  <div v-if="newsStore.articles.length">
    <h2>Top News in {{ country }}</h2>
    <ul>
      <li v-for="article in newsStore.articles" :key="article.url">
        <a :href="article.url" target="_blank">{{ article.title }}</a>
      </li>
    </ul>
  </div>
  <div v-else-if="newsStore.loading">Loading...</div>
  <div v-else-if="newsStore.error">Error: {{ newsStore.error.message }}</div>
  <div v-else>
    <p>No news available</p>
  </div>
</template>

<script>
import { useNewsStore } from 'src/stores/news-collector';
import {  onMounted } from 'vue';

export default {
  name: 'NewsComponent',
  props: {
    country: {
      type: String,
      required: true,
    },
  },
  setup(props) {
    const newsStore = useNewsStore();
    onMounted(() => {
      newsStore.fetchNews(props.country);
    });

    return { newsStore };
  },
};
</script>

<style scoped>
</style>
