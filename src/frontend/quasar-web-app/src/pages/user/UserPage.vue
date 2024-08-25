<template>
  <q-page class="flex flex-center" style="background-color: azure">
    <img
      alt="Quasar logo"
      src="~assets/icons/app_background.png"
      style="width: 200px; height: 200px"
    />
    <q-dialog v-model="isLoginDialogVisible" persistent transition-hide="scale" transition-show="scale">
      <UserLogin></UserLogin>
    </q-dialog>
    <q-dialog v-model="isInfoDialogVisible" persistent transition-hide="scale" transition-show="scale">
    </q-dialog>
  </q-page>
</template>

<script setup>
import {ref} from "vue";
import UserLogin from "components/user/UserLogin.vue";

defineOptions({
  name: "IndexPage",
});
const isLoginDialogVisible = ref(true);
const isInfoDialogVisible = ref(false);

// Functions
function handleLoginButtonClick() {
  isLoginDialogVisible.value = true;
}

function handleInfoButtonClick() {
  isInfoDialogVisible.value = true;
  appInfos.printInfo();
}
</script>
