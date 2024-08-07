<template>
  <q-layout view="hHh lpR fFf">
    <!-- HEADER -->
    <app-header></app-header>
    <!-- DRAWERS -->
    <app-drawer></app-drawer>
    <!-- CONTENTS -->
    <q-page-container >
      <router-view />
    </q-page-container>
    <!-- FOOTER -->
    <app-footer></app-footer>
  </q-layout>
</template>

<script setup>
import AppHeader from "layouts/header/AppHeader.vue";
import AppDrawer from "layouts/drawers/AppDrawer.vue";
import AppFooter from "layouts/footer/AppFooter.vue";
//
defineOptions({
  name: "MainLayout",
  components: { AppHeader, AppDrawer, AppFooter },
});
</script>
