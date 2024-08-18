<template>
  <div>
    <h1>User Details</h1>
    <div v-if="loading">Loading...</div>
    <div v-if="error">{{ error }}</div>
    <div v-if="user">
      <p><strong>Name:</strong> {{ user.name }}</p>
      <p><strong>Email:</strong> {{ user.email }}</p>
      <router-link :to="'/user/' + user.id + '/edit'">Edit</router-link>
      <button @click="deleteUser(user.id)">Delete</button>
    </div>
  </div>
</template>

<script>
import { useUserStore } from 'stores/user-store';
import { useRoute, useRouter } from 'vue-router';

export default {
  setup() {
    const route = useRoute();
    const router = useRouter();
    const userStore = useUserStore();

    userStore.fetchUser(route.params.id);

    const deleteUser = async (id) => {
      await userStore.deleteUser(id);
      await router.push('/');
    };

    return {
      ...userStore,
      deleteUser,
    };
  },
};
</script>
