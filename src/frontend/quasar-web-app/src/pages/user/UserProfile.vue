<template>
  <q-page class="profile-page">
    <UserProfileCard :loading="authStore.loading" :updateProfile="updateUserProfile" :user="authStore.user"/>
  </q-page>
</template>

<script>
defineOptions({
  name: "UserProfile",
});
import {useAuthStore} from 'stores/user/auth-store';
import UserProfileCard from 'components/user/UserProfileCard.vue';

export default {
  components: {UserProfileCard},
  setup() {
    const authStore = useAuthStore();

    const updateUserProfile = async () => {
      authStore.password = undefined;  // Initialize password
      const updatedUser = {
        name: authStore.user.name,
        password: authStore.password,
      };
      await authStore.updateUser(updatedUser);
    };

    return {authStore, updateUserProfile};
  }
};
</script>
