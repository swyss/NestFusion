<template>
  <q-page>
    <q-list bordered separator>
      <q-item v-for="user in users" :key="user.id" @click="selectUser(user.id)">
        <q-item-section>
          <q-item-label>{{ user.name }}</q-item-label>
          <q-item-label caption>{{ user.email }}</q-item-label>
        </q-item-section>
      </q-item>
    </q-list>
    <q-btn label="Neuen Benutzer hinzufügen" @click="createUser" />
  </q-page>
</template>

<script>
import { useUserStore } from 'src/stores/user-store';

export default {
  setup() {
    const userStore = useUserStore();

    // Fetch users when component is mounted
    userStore.fetchUsers();

    return {
      users: userStore.users,
      selectUser: (id) => {
        // Navigation to user details or edit page
      },
      createUser: () => {
        // Navigation to user creation page
      },
    };
  },
};
</script>
