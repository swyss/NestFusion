<template>
  <q-layout view="hHh lpR fFf">
    <!-- HEADER -->
    <user-header></user-header>
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
import UserHeader from "layouts/user/header/AppHeader.vue";
import AppDrawer from "layouts/main/drawers/AppDrawer.vue";
import AppFooter from "layouts/main/footer/AppFooter.vue";
//
defineOptions({
  name: "UserLayout",
  components: { UserHeader, AppDrawer, AppFooter },
});
</script>
