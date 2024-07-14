import { defineStore } from 'pinia';
import axios from 'axios';

export const useNewsStore = defineStore('news', {
  state: () => ({
    articles: [],
    selectedCategory: 'general',
    apiKey: 'YOUR_API_KEY', // Ersetzen Sie dies durch Ihren tatsächlichen API-Schlüssel
  }),
  getters: {
    // Optional: Getter für spezifische Datenmanipulationen
  },
  actions: {
    async fetchNews() {
      const url = `https://newsapi.org/v2/top-headlines?country=de&category=${this.selectedCategory}&apiKey=${this.apiKey}`;
      try {
        const response = await axios.get(url);
        this.articles = response.data.articles;
      } catch (error) {
        console.error('Error fetching news:', error);
      }
    },
    changeCategory(newCategory) {
      this.selectedCategory = newCategory;
      this.fetchNews(); // Nachrichten neu laden, wenn die Kategorie geändert wird
    },
  },
});
