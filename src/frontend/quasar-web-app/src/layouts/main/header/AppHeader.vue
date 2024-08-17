<script setup>
defineOptions({
  name: "AppHeader",
});

// Importing components
import HeaderTitleToolbar from "layouts/main/header/HeaderTitleToolbar.vue";
import HeaderMenuToolbar from "layouts/main/header/HeaderMenuToolbar.vue";

// Constants for components
const TitleToolbar = HeaderTitleToolbar;
const MenuToolbar = HeaderMenuToolbar;
</script>
<template>
  <q-header class="bg-primary text-black" elevated>
    <TitleToolbar />
    <MenuToolbar />
  </q-header>
</template>
<style scoped></style>
