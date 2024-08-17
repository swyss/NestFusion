<script setup>
import AppLeftDrawer from "./AppLeftDrawer.vue";
import AppRightDrawer from "./AppRightDrawer.vue";

defineOptions({
  name: "AppDrawer",
  components: { AppLeftDrawer, AppRightDrawer },
});
</script>

<template>
  <AppLeftDrawer />
  <AppRightDrawer />
</template>

<style scoped></style>
