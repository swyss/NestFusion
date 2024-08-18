<template>
  <div>
    <h1>{{ isEdit ? "Edit User" : "New User" }}</h1>
    <form @submit.prevent="saveUser">
      <div>
        <label for="name">Name:</label>
        <input type="text" v-model="userData.name" required />
      </div>
      <div>
        <label for="email">Email:</label>
        <input type="email" v-model="userData.email" required />
      </div>
      <button type="submit">{{ isEdit ? "Update" : "Create" }}</button>
    </form>
  </div>
</template>

<script>
import { useUserStore } from 'stores/user-store';
import { useRoute, useRouter } from 'vue-router';
import { ref, onMounted } from 'vue';

export default {
  setup() {
    const route = useRoute();
    const router = useRouter();
    const userStore = useUserStore();
    const userData = ref({ name: '', email: '' });
    const isEdit = ref(false);

    onMounted(async () => {
      if (route.params.id) {
        await userStore.fetchUser(route.params.id);
        userData.value = { ...userStore.user };
        isEdit.value = true;
      }
    });

    const saveUser = async () => {
      if (isEdit.value) {
        await userStore.updateUser(route.params.id, userData.value);
      } else {
        await userStore.createUser(userData.value);
      }
      await router.push('/');
    };

    return {
      userData,
      isEdit,
      saveUser,
    };
  },
};
</script>
