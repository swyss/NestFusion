<template>
  <q-list>
    <q-item-label header> Essential Links</q-item-label>
    <EssentialLink
      v-for="link in links"
      :key="link.title"
      v-bind="link"
      title=""
    />
  </q-list>
</template>

<script setup>
import EssentialLink from "components/EssentialLink.vue";
defineProps(["links"]);
</script>
