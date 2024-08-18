<template>
  <q-page>
    <q-form @submit="createUser">
      <q-input v-model="user.name" label="Name" />
      <q-input v-model="user.email" label="E-Mail" />
      <q-btn label="Erstellen" type="submit" />
    </q-form>
  </q-page>
</template>

<script>
import { useUserStore } from 'src/stores/user-store';
import { useRouter } from 'vue-router';

export default {
  setup() {
    const userStore = useUserStore();
    const router = useRouter();

    const user = {
      name: '',
      email: ''
    };

    return {
      user,
      createUser: () => {
        userStore.createUser(user);
        router.push('/users');
      },
    };
  },
};
</script>
