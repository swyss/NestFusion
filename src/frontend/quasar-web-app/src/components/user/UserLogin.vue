<template>
  <q-page class="bg-light-green window-height window-width row justify-center items-center">
    <div class="column">
      <div class="row">
        <h5 class="text-h5 text-white q-my-md">Login</h5>
      </div>
      <div class="row">
        <q-card bordered class="q-pa-lg shadow-1" square>
          <q-card-section>
            <q-form class="q-gutter-md">
              <q-input v-model="email" clearable filled label="Email" square type="email"/>
              <q-input v-model="password" clearable filled label="Password" square type="password"/>
            </q-form>
          </q-card-section>
          <q-card-actions class="q-px-md">
            <q-btn :loading="authStore.loading" class="full-width" color="light-green-7" label="Login" size="lg"
                   unelevated @click="handleLogin"/>
          </q-card-actions>
          <q-card-section class="text-center q-pa-none">
            <p class="text-grey-6">Not registered? Create an Account</p>
          </q-card-section>
        </q-card>
      </div>
    </div>
  </q-page>
</template>

<script>
import {useAuthStore} from 'stores/user/auth-store';
import {ref} from 'vue';
import {useRouter} from 'vue-router';

export default {
  setup() {
    // Constants
    const SUCCESS_ROUTE = '/';

    // State
    const authStore = useAuthStore();
    const router = useRouter();
    const email = ref('');
    const password = ref('');

    // Functions
    const showErrorNotification = (message) => {
      this.$q.notify({type: 'negative', message});
    };

    const login = async (loginPayload) => {
      await authStore.login(loginPayload);
      return !authStore.error;
    };

    const handleLogin = async () => {
      const loginPayload = {email: email.value, password: password.value};
      const isLoggedIn = await login(loginPayload);
      if (isLoggedIn) {
        await router.push(SUCCESS_ROUTE);
      } else {
        showErrorNotification(authStore.error);
      }
    };

    return {authStore, email, password, handleLogin};
  }
};
</script>
