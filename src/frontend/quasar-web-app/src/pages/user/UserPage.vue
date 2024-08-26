<template>
  <q-page class="flex flex-center" style="background-color: azure">
    <img
      alt="Quasar logo"
      src="~assets/icons/app_background.png"
      style="width: 200px; height: 200px"
    />
    <q-dialog v-model="isLoginDialogVisible" transition-hide="scale" transition-show="scale" backdrop-filter="sepia(90%)">
      <UserLogin></UserLogin>
    </q-dialog>
    <q-dialog v-model="isInfoDialogVisible" transition-hide="scale" transition-show="scale" backdrop-filter="sepia(90%)">
      <UserRegister></UserRegister>
    </q-dialog>
  </q-page>
</template>

<script setup>
import {ref} from "vue";
import UserLogin from "components/user/UserLogin.vue";
import UserRegister from "components/user/UserRegister.vue";

defineOptions({
  name: "IndexPage",
});
const isLoginDialogVisible = ref(true);
const isInfoDialogVisible = ref(false);

</script>
