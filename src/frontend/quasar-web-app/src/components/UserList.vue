<template>
  <div>
    <h1>User List</h1>
    <div v-if="loading">Loading...</div>
    <div v-if="error">{{ error }}</div>
    <ul v-if="users.length">
      <li v-for="user in users" :key="user.id">
        <router-link :to="'/user/' + user.id">{{ user.name }}</router-link>
      </li>
    </ul>
    <router-link to="/new-user">Add New User</router-link>
  </div>
</template>

<script>
import { useUserStore } from 'stores/user-store';

export default {
  setup() {
    const userStore = useUserStore();
    userStore.fetchUsers();

    return {
      ...userStore,
    };
  },
};
</script>
