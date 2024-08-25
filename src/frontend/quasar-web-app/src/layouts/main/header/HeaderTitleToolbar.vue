<script setup>
defineOptions({
  name: "TitleToolbar",
});

// Imports
import {useAppPropertyStore} from "stores/app/app-properties-store";

// Variables
const appProperties = useAppPropertyStore();

</script>

<template>
  <div class="q-pt-md">
    <q-toolbar>
      <q-toolbar-title>
        <q-avatar><img alt="logo" src="src/assets/icons/app_icon.png"></q-avatar>
        NestFusion
      </q-toolbar-title>
      <q-space></q-space>
      <q-btn
        dense
        flat
        padding="sm"
        square
        @click="appProperties.toggleDarkMode"
      >
        <q-icon name="bi-brilliance"/>
      </q-btn>
    </q-toolbar>
  </div>
</template>
<style scoped></style>
