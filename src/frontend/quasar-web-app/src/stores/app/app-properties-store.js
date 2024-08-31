import {defineStore} from "pinia";
import {useQuasar} from "quasar";
import {computed, ref} from "vue";

export const useAppPropertyStore = defineStore("properties", () => {
  const $q = useQuasar();
  const leftDrawerOpen = ref(false);
  const rightDrawerOpen = ref(false);

  // HEADER Link List
  const linkList = ref([
    {name: "tasks", label: "tasks", link: "tasks"},
    {name: "weather", label: "weather", link: "weather"},
  ]);
  // Menu list LEFT
  const menuListLeft = ref([
    {icon: "bi-inboxes", label: "Inbox", route: "/inbox", separator: true},
    {icon: "bi-layout-wtf", label: "Outbox", route: "/outbox", separator: false},
    {icon: "bi-trash3", label: "Trash", route: "/trash", separator: false},
    {icon: "bi-bug", label: "Spam", route: "/spam", separator: true},
    {icon: "bi-gear", label: "Settings", route: "/settings", separator: false},
    {icon: "bi-chat-left-text", label: "Send Feedback", route: "/feedback", separator: false},
    {icon: "bi-question-diamond", label: "Help", route: "/help", separator: false},
  ]);
  // Menu list RIGHT
  const menuListRight = ref([
    {icon: "bi-person-gear", label: "Admin", route: "/admin", separator: true},
    {icon: "bi-gear", label: "Settings", route: "/settings", separator: false},
    {icon: "bi-chat-left-text", label: "Send Feedback", route: "/feedback", separator: false},
    {icon: "bi-question-diamond", label: "Help", route: "/help", separator: false},
  ]);

  const leftDrawerState = computed(() => leftDrawerOpen.value);
  const rightDrawerState = computed(() => rightDrawerOpen.value);

  function toggleLeftDrawer() {
    leftDrawerOpen.value = !leftDrawerOpen.value;
  }

  function toggleRightDrawer() {
    rightDrawerOpen.value = !rightDrawerOpen.value;
  }

  function toggleDarkMode() {
    $q.dark.toggle();
  }

  function init() {
    leftDrawerOpen.value = false;
    rightDrawerOpen.value = false;
    $q.dark.set(false);
  }

  return {
    init,
    leftDrawerState,
    rightDrawerState,
    toggleLeftDrawer,
    toggleRightDrawer,
    toggleDarkMode,
    menuListLeft,
    menuListRight,
    linkList,
  };
});
