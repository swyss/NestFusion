<template>
  <q-page>
    <q-form @submit="updateUser">
      <q-input v-model="user.name" label="Name" />
      <q-input v-model="user.email" label="E-Mail" />
      <q-btn label="Speichern" type="submit" />
      <q-btn label="Löschen" color="negative" @click="deleteUser" />
    </q-form>
  </q-page>
</template>

<script>
import { useUserStore } from 'src/stores/user-store';
import { useRoute, useRouter } from 'vue-router';

export default {
  setup() {
    const userStore = useUserStore();
    const route = useRoute();
    const router = useRouter();

    const userId = route.params.id;
    userStore.fetchUser(userId);

    return {
      user: userStore.user,
      updateUser: () => {
        userStore.updateUser(userId, userStore.user);
        router.push('/users');
      },
      deleteUser: () => {
        userStore.deleteUser(userId);
        router.push('/users');
      },
    };
  },
};
</script>
