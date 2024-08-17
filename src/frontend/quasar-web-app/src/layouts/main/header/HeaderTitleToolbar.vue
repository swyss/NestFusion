<script setup>
defineOptions({
  name: "TitleToolbar",
});

// Imports
import { ref } from "vue";
import { useAppPropertyStore } from "stores/app-properties-store";
import { useAppInfoStore } from "stores/app-info-store";

// Variables
const appProperties = useAppPropertyStore();
const appInfos = useAppInfoStore();
const isLoginDialogVisible = ref(false);
const isInfoDialogVisible = ref(false);

// Functions
function handleLoginButtonClick() {
  isLoginDialogVisible.value = true;
}

function handleInfoButtonClick() {
  isInfoDialogVisible.value = true;
  appInfos.printInfo();
}

appInfos.printInfo = function () {
  // Add printing info logic here if needed
};
</script>

<template>
  <div class="q-pt-md">
    <q-toolbar>
      <q-toolbar-title>
        <q-avatar></q-avatar>
        Title
      </q-toolbar-title>
      <q-space></q-space>
      <q-btn dense flat padding="sm md" square @click="handleLoginButtonClick">
        <q-icon name="bi-person-square" />
      </q-btn>
      <q-dialog v-model="isLoginDialogVisible">
        <!--        <user-login></user-login>-->
      </q-dialog>
      <q-btn dense flat padding="sm" square @click="handleInfoButtonClick">
        <q-icon name="bi-info-circle" />
      </q-btn>
      <q-dialog
        v-model="isInfoDialogVisible"
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
