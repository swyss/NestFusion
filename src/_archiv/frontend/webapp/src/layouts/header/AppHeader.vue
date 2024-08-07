<script setup>
defineOptions({
  name: "AppHeader",
});
// imports
import TitleToolbar from "layouts/header/HeaderTitleToolbar.vue";
import MenuToolbar from "layouts/header/HeaderMenuToolbar.vue";
// variables

// functions
</script>

<template>
  <q-header class="bg-primary text-black" elevated>
    <title-toolbar></title-toolbar>
    <menu-toolbar></menu-toolbar>
  </q-header>
</template>

<style scoped></style>
