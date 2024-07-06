import { defineStore } from "pinia";
import { useQuasar } from "quasar";
import { ref } from "vue";

export const useAppPropertyStore= defineStore("properties", () => {
  const $q = useQuasar();
  const leftDrawerOpen = ref(false);
  const rightDrawerOpen = ref(false);
  //
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
  //
  return {
    leftDrawerOpen,
    rightDrawerOpen,
    init,
    toggleLeftDrawer,
    toggleRightDrawer,
    toggleDarkMode,
  };
});
