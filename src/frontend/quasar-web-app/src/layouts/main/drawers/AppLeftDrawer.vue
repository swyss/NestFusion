<script setup>
// imports
import { useAppPropertyStore } from "stores/app-properties-store";

// constants
const MENU_LIST = [
  { icon: "bi-inboxes", label: "Inbox", separator: true },
  { icon: "bi-layout-wtf", label: "Outbox", separator: false },
  { icon: "bi-trash3", label: "Trash", separator: false },
  { icon: "bi-bug", label: "Spam", separator: true },
  { icon: "bi-gear", label: "Settings", separator: false },
  { icon: "bi-chat-left-text", label: "Send Feedback", separator: false },
  {
    icon: "bi-question-diamond",
    iconColor: "primary",
    label: "Help",
    separator: false,
  },
];

// setup
const appProperties = useAppPropertyStore();
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
        <template v-for="(menuItem, index) in MENU_LIST" :key="index">
          <q-item v-ripple :active="menuItem.label === 'Outbox'" clickable>
            <q-item-section avatar>
              <q-icon :name="menuItem.icon" />
            </q-item-section>
            <q-item-section>
              {{ menuItem.label }}
            </q-item-section>
          </q-item>
          <q-separator v-if="menuItem.separator" :key="'sep' + index" />
        </template>
      </q-list>
    </q-scroll-area>
  </q-drawer>
</template>

<style scoped></style>
