<script setup>
// Component options
defineOptions({
  name: "MenuToolbar",
});

// Imports
import {useRouter} from 'vue-router'; // Importing useRouter hook
import {useAppPropertyStore} from "stores/app/app-properties-store";

// Constants
const ICON_NAMES = {
  list: "bi-list",
  house: "bi-house",
  star: "bi-star",
  sliders: "bi-sliders",
};

// Store variable
const appProperties = useAppPropertyStore();

// Router instance
const router = useRouter();

// Method to navigate
const navigateTo = (route) => {
  router.replace(route); // Using the router instance to navigate
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
        <q-icon :name="ICON_NAMES.list"/>
      </q-btn>
      <q-separator color="on_primary" inset vertical/>
      <q-btn dense flat padding="sm md" square @click="() => navigateTo('/')">
        <q-icon :name="ICON_NAMES.house"/>
      </q-btn>
      <q-btn dense flat padding="sm md" square @click="() => navigateTo('/user')">
        <q-icon :name="ICON_NAMES.star" color="accent"/>
      </q-btn>
      <q-separator color="on_primary" inset vertical/>
      <q-tabs align="left" inline-label shrink stretch>
        <q-route-tab
          v-for="tab in appProperties.linkList"
          :key="tab.name"
          :to="tab.link"
          v-bind="tab"
        />
      </q-tabs>
      <q-space></q-space>
      <q-btn dense flat padding="sm md" square>
        <q-icon name="bi-person"/>
      </q-btn>
      <q-separator color="$on_primary" inset vertical/>
      <q-btn dense flat padding="sm md" square @click="toggleRightDrawer">
        <q-icon :name="ICON_NAMES.sliders"/>
      </q-btn>
    </q-toolbar>
  </div>
</template>
<style scoped></style>
