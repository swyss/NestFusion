<template>
  <q-page class="flex flex-center" style="background-color: azure">
    <img
      alt="Quasar logo"
      src="~assets/icons/app_background.png"
      style="width: 200px; height: 200px"
    />
    <q-dialog v-model="isLoginDialogVisible" persistent transition-hide="scale" transition-show="scale" backdrop-filter="sepia(90%)">
<UserLogin></UserLogin>
    </q-dialog>
    <q-dialog v-model="isRegisterDialogVisible" persistent transition-hide="scale" transition-show="scale" backdrop-filter="sepia(90%)">
      <UserRegister></UserRegister>
    </q-dialog>
  </q-page>
</template>
<script setup>
import UserLogin from "components/user/UserLogin.vue";
import UserRegister from "components/user/UserRegister.vue";
import { isLoginDialogVisible, isRegisterDialogVisible } from '/src/utils/user/modalFunctions';
import {ref} from "vue";
const toolbar= ref(false);

defineOptions({
  name: "UserPage",
});
</script>
