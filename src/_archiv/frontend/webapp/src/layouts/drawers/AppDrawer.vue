<script setup>
import AppLeftDrawer from "layouts/drawers/AppLeftDrawer.vue";
import AppRightDrawer from "layouts/drawers/AppRightDrawer.vue";
//
defineOptions({
  name: "AppDrawer",
  components: { AppLeftDrawer, AppRightDrawer },
});
</script>

<template>
  <app-left-drawer></app-left-drawer>
  <app-right-drawer></app-right-drawer>
</template>

<style scoped></style>
