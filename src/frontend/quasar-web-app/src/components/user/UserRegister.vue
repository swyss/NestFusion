<template>
  <q-card class="bg-light-green justify-center items-center q-pa-lg">
    <q-card-section class="row items-center q-pb-none">
      <div class="text-h6 text-white">Register</div>
      <q-space />
      <q-btn text-color="white" icon="bi-x" flat round dense v-close-popup />
    </q-card-section>

    <q-card-section>
      <div class="column">
        <div class="row">
          <q-card bordered class="q-pa-lg shadow-1" square>
            <q-card-section>
              <q-form class="q-gutter-md">
                <q-input v-model="name" clearable filled label="name" square type="name"/>
                <q-input v-model="email" clearable filled label="email" square type="email"/>
                <q-input v-model="password" clearable filled label="password" square type="password"/>
                <q-input v-model="passwordConfirm" clearable filled label="confirm password" square type="password"/>
              </q-form>
            </q-card-section>
            <q-card-actions class="q-px-md">
              <q-btn :loading="authStore.loading" class="full-width" color="light-green-7" label="Register" size="lg"
                     unelevated @click="handleRegister"/>
            </q-card-actions>
            <q-card-section class="text-center q-pa-none">
              <p class="text-grey-6">Have an account? <span @click="switchToLoginModal" style="cursor: pointer; color: #42b983;">Login</span></p>
            </q-card-section>
          </q-card>
        </div>
      </div>
    </q-card-section>
  </q-card>
</template>

<script setup>
import { useRouter } from 'vue-router';
import { ref } from 'vue';
import axios from 'axios';
import { useQuasar } from 'quasar';
import { useAuthStore } from "stores/user/auth-store";
import {hideRegisterDialog, showLoginDialog} from 'src/utils/user/modalFunctions';

const router = useRouter();
const $q = useQuasar();
const name = ref('');
const email = ref('');
const password = ref('');
const passwordConfirm = ref('');
const authStore = useAuthStore();

const SUCCESS_MESSAGE = 'Registration successful!';
const FAILURE_MESSAGE = 'Registration failed!';

const notify = (type, message) => {
  $q.notify({ type, message });
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
  if (password.value !== passwordConfirm.value) {
    notify('negative', 'Passwords do not match!');
    return;
  }
  registerUser({
    name: name.value,
    email: email.value,
    password: password.value
  });
};

const switchToLoginModal = () => {
  hideRegisterDialog();
  showLoginDialog();
};
</script>
