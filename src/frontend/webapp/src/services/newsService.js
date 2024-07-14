import axios from 'axios';

const API_KEY = 'API-Key'; // Replace
const BASE_URL = 'https://newsapi.org/v2/top-headlines';

export const getNews = async (country) => {
  const response = await axios.get(`${BASE_URL}?country=${country}&apiKey=${API_KEY}`);
  return response.data.articles;
};
