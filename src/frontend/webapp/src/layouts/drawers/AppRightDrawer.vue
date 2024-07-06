<script>
import { ref } from 'vue'

const menuList = [
  {
    icon: 'bi-inboxes',
    label: 'Inbox',
    separator: true
  },
  {
    icon: 'bi-layout-wtf',
    label: 'Outbox',
    separator: false
  },
  {
    icon: 'bi-trash3',
    label: 'Trash',
    separator: false
  },
  {
    icon: 'bi-bug',
    label: 'Spam',
    separator: true
  },
  {
    icon: 'bi-gear',
    label: 'Settings',
    separator: false
  },
  {
    icon: 'bi-chat-left-text',
    label: 'Send Feedback',
    separator: false
  },
  {
    icon: 'bi-question-diamond',
    iconColor: 'primary',
    label: 'Help',
    separator: false
  }
]

export default {
  setup() {
    return {
      drawer: ref(false),
      menuList
    }
  }
}
</script>

<template>
  <q-drawer
    v-model="drawer"
    show-if-above
    :width="200"
    :breakpoint="500"
    bordered
    :class="$q.dark.isActive ? 'bg-grey-9' : 'bg-grey-3'"
  >
    <q-scroll-area class="fit">
      <q-list>

        <template v-for="(menuItem, index) in menuList" :key="index">
          <q-item clickable :active="menuItem.label === 'Outbox'" v-ripple>
            <q-item-section avatar>
              <q-icon :name="menuItem.icon"/>
            </q-item-section>
            <q-item-section>
              {{ menuItem.label }}
            </q-item-section>
          </q-item>
          <q-separator :key="'sep' + index" v-if="menuItem.separator"/>
        </template>

      </q-list>
    </q-scroll-area>
  </q-drawer>
</template>

<style scoped></style>
