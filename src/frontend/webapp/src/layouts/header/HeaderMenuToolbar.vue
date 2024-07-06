<script setup>
defineOptions({
  name: "MenuToolbar",
});
// imports
import { useAppPropertyStore } from "stores/AppProperty";
// variables
const appProperties = useAppPropertyStore();
const tab_menu = [
  { name: "mails", label: "Mails", link: "/mails" },
  { name: "alarms", label: "Alarms", link: "/alarms" },
  { name: "movies", label: "Movies", link: "/movies" },
  { name: "photos", label: "Photos", link: "/photos" },
  { name: "videos", label: "Videos", link: "/videos" },
  { name: "music", label: "Music", link: "/music" },
  { name: "calendar", label: "Calendar", link: "/calendar" },
  { name: "addressbook", label: "Address Book", link: "/addressbook" },
  { name: "settings", label: "Settings", link: "/settings" },
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
        <q-icon color="accent" name="bi-star" />
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
