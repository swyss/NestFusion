import { defineStore } from "pinia";
import { useQuasar } from "quasar";
import {computed, ref} from "vue";



export const useAppPropertyStore= defineStore("properties", () => {
  const $q = useQuasar();
  const leftDrawerOpen = ref(false);
  const rightDrawerOpen = ref(false);
  //
  const leftDrawerState = computed(() => leftDrawerOpen.value)
  const rightDrawerState = computed(() => rightDrawerOpen.value)
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
    init,
    leftDrawerState,
    rightDrawerState,
    toggleLeftDrawer,
    toggleRightDrawer,
    toggleDarkMode,
  };
});
