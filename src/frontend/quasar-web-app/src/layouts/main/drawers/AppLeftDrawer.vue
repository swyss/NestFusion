<script setup>
// imports
import {useAppPropertyStore} from "stores/app/app-properties-store";
import {useRoute} from "vue-router";

// Get the app property store instance
const appProperties = useAppPropertyStore();
// Get the current route to handle active link highlighting
const route = useRoute();
</script>

<template>
  <q-drawer
    :breakpoint="500"
    :class="$q.dark.isActive ? 'bg-grey-9' : 'bg-grey-3'"
    :model-value="appProperties.leftDrawerState"
    :width="200"
    bordered
    side="left"
  >
    <q-scroll-area class="fit">
      <q-list>
        <template v-for="(menuItem, index) in appProperties.menuListLeft" :key="index">
          <q-item
            v-ripple
            :active="route.path === menuItem.route"
            :to="menuItem.route"
            clickable
            tag="router-link"
          >
            <q-item-section avatar>
              <q-icon :name="menuItem.icon"/>
            </q-item-section>
            <q-item-section>
              {{ menuItem.label }}
            </q-item-section>
          </q-item>
          <q-separator v-if="menuItem.separator" :key="'sep' + index"/>
        </template>
      </q-list>
    </q-scroll-area>
  </q-drawer>
</template>

<style scoped></style>
