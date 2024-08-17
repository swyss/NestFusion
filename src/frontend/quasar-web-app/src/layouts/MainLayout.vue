<template>
  <q-layout :view="layoutView">
    <!-- HEADER -->
    <app-header></app-header>
    <!-- DRAWERS -->
    <app-drawer></app-drawer>
    <!-- CONTENTS -->
    <q-page-container>
      <router-view/>
    </q-page-container>
    <!-- FOOTER -->
    <app-footer></app-footer>
  </q-layout>
</template>

<script setup>
import AppHeader from "layouts/main/header/AppHeader.vue";
import AppDrawer from "layouts/main/drawers/AppDrawer.vue";
import AppFooter from "layouts/main/footer/AppFooter.vue";

// Extracted constant for layout view setting
const layoutView = "hHh lpR fFf";

defineOptions({
  name: "MainLayout",
  components: {AppHeader, AppDrawer, AppFooter},
});
</script>
