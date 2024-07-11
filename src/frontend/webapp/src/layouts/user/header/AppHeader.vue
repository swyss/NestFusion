<script setup>
defineOptions({
  name: "UserHeader",
});
// imports
import TitleToolbar from "layouts/main/header/HeaderTitleToolbar.vue";
import MenuToolbar from "layouts/user/header/HeaderMenuToolbar.vue";
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
