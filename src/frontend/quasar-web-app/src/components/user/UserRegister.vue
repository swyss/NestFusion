<template>
  <q-page class="bg-light-green row justify-center items-center q-pa-xl" style="max-height: fit-content; min-height: fit-content">
    <div class="column">
      <div class="row">
        <h5 class="text-h5 text-white q-my-md">Register</h5>
      </div>
      <div class="row">
        <q-card bordered class="q-pa-lg shadow-1" square>
          <q-card-section>
            <q-form class="q-gutter-md">
              <q-input v-model="name" clearable filled label="name" square type="name"/>
              <q-input v-model="email" clearable filled label="email" square type="email"/>
              <q-input v-model="password" clearable filled label="password" square type="password"/>
            </q-form>
          </q-card-section>
          <q-card-actions class="q-px-md">
            <q-btn :loading="authStore.loading" class="full-width" color="light-green-7" label="Register" size="lg"
                   unelevated @click="handleRegister"/>
          </q-card-actions>
          <q-card-section class="text-center q-pa-none">
            <p class="text-grey-6">Have an account? Login</p>
          </q-card-section>
        </q-card>
      </div>
    </div>
  </q-page>
</template>

<script setup>
import {useRouter} from 'vue-router';
import {ref} from 'vue';
import axios from 'axios';
import {useQuasar} from 'quasar';
import {useAuthStore} from "stores/user/auth-store";

const router = useRouter();
const $q = useQuasar();
const name = ref('');
const email = ref('');
const password = ref('');
const authStore = useAuthStore();

const SUCCESS_MESSAGE = 'Registration successful!';
const FAILURE_MESSAGE = 'Registration failed!';

const notify = (type, message) => {
  $q.notify({type, message});
};

const registerUser = async (userData) => {
  try {
    await axios.post('/api/user', userData);
    notify('positive', SUCCESS_MESSAGE);
    await router.push('/login');
  } catch (error) {
    notify('negative', FAILURE_MESSAGE);
  }
};

const handleRegister = () => {
  registerUser({
    name: name.value,
    email: email.value,
    password: password.value
  });
};
</script>
