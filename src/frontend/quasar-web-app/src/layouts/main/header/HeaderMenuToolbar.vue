<script setup>
// Component options
defineOptions({
  name: "MenuToolbar",
});

// Imports
import { useAppPropertyStore } from "stores/app-properties-store";

// Constants
const TABS = [
  { name: "account", label: "account", link: "account" },
  { name: "login", label: "login", link: "login" },
];

// Variables
const appProperties = useAppPropertyStore();
const tabs = TABS;

// Methods
const goToHome = () => {
  $router.replace("/");
};

const goToUser = () => {
  $router.replace("/user");
};

// Methods to toggle drawers
const toggleLeftDrawer = () => {
  appProperties.toggleLeftDrawer();
};

const toggleRightDrawer = () => {
  appProperties.toggleRightDrawer();
};
</script>

<template>
  <div class="q-py-md">
    <q-toolbar>
      <q-btn dense flat padding="sm md" square @click="toggleLeftDrawer">
        <q-icon name="bi-list" />
      </q-btn>
      <q-separator color="on_primary" inset vertical />
      <q-btn dense flat padding="sm md" square @click="goToHome">
        <q-icon name="bi-house" />
      </q-btn>
      <q-btn dense flat padding="sm md" square @click="goToUser">
        <q-icon color="accent" name="bi-star" />
      </q-btn>
      <q-separator color="on_primary" inset vertical />
      <q-tabs align="left" v-model="tabs" inline-label shrink stretch>
        <q-route-tab
          v-for="tab in tabs"
          :key="tab.name"
          v-bind="tab"
          :to="tab.link"
        />
      </q-tabs>
      <q-space></q-space>
      <q-separator color="$on_primary" inset vertical />
      <q-btn dense flat padding="sm md" square @click="toggleRightDrawer">
        <q-icon name="bi-sliders" />
      </q-btn>
    </q-toolbar>
  </div>
</template>

<style scoped></style>
