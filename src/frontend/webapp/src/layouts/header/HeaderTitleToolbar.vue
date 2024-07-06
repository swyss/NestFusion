<script setup>
defineOptions({
  name: "TitleToolbar",
});
// imports
import { ref } from "vue";
import { useAppPropertyStore } from "stores/AppProperty";
import { useAppInfoStore } from "stores/AppInfo";
// import UserLogin from "pages/user/LoginPage.vue";
// import InfoPage from "pages/general/InfoPage.vue";
// variable
const appProperties = useAppPropertyStore();
const appInfos = useAppInfoStore();
const loginInception = ref(false);
const infoInception = ref(false);

// functions
function clickLoginBtn() {
  loginInception.value = true;
}

appInfos.printInfo = function () {

};

function clickInfoBtn() {
  infoInception.value = true;
  appInfos.printInfo();
}
</script>

<template>
  <div class="q-pt-md">
    <q-toolbar>
      <q-toolbar-title>
        <q-avatar> </q-avatar>
        Title
      </q-toolbar-title>
      <q-space></q-space>
      <q-btn dense flat padding="sm md" square @click="clickLoginBtn">
        <q-icon name="bi-person-square" />
      </q-btn>
      <q-dialog v-model="loginInception">
<!--        <user-login></user-login>-->
      </q-dialog>
      <q-btn dense flat padding="sm" square @click="clickInfoBtn">
        <q-icon name="bi-info-circle" />
      </q-btn>
      <q-dialog
        v-model="infoInception"
        persistent
        transition-hide="scale"
        transition-show="scale"
      >
<!--        <info-page></info-page>-->
      </q-dialog>
      <q-btn
        dense
        flat
        padding="sm"
        square
        @click="appProperties.toggleDarkMode"
      >
        <q-icon name="bi-moon-stars" />
      </q-btn>
    </q-toolbar>
  </div>
</template>

<style scoped></style>
