<script setup>
defineOptions({
  name: "MenuToolbar",
});
// imports
import { useAppPropertyStore } from "stores/app-properties-store";
// variables
const appProperties = useAppPropertyStore();
const tab_menu = [
  { name: "account", label: "account", link: "account" },
  { name: "login", label: "login", link: "login" },
  { name: "components", label: "components", link: "components" },
];
// functions
</script>
<template>
  <div class="q-py-md">
    <q-toolbar>
      <q-btn
        dense
        flat
        padding="sm md"
        square
        @click="appProperties.toggleLeftDrawer()"
      >
        <q-icon name="bi-list" />
      </q-btn>
      <q-separator color="on_primary" inset vertical />
      <q-btn dense flat padding="sm md" square>
        <q-icon name="bi-house" @click="$router.replace('/')" />
      </q-btn>
      <q-btn dense flat padding="sm md" square>
        <q-icon color="accent" name="bi-star" @click="$router.replace('/user')"/>
      </q-btn>
      <q-separator color="on_primary" inset vertical />
      <q-tabs align="left" v-model="tab_menu" inline-label shrink stretch>
        <q-route-tab
          v-for="tab in tab_menu"
          :key="tab.name"
          v-bind="tab"
          :to="tab.link"
        />
      </q-tabs>
      <q-space></q-space>
      <q-separator color="$on_primary" inset vertical />
      <q-btn
        dense
        flat
        padding="sm md"
        square
        @click="appProperties.toggleRightDrawer()"
      >
        <q-icon name="bi-sliders" />
      </q-btn>
    </q-toolbar>
  </div>
</template>

<style scoped></style>
