import { defineStore } from "pinia";
import { useQuasar } from "quasar";

export const useAppInfoStore = defineStore("infos", () => {
  const $q = useQuasar();
  const msg = "Hello, from APP";
  // functions
  function printInfo() {
    console.log(msg);
    console.log($q.version);
    console.log($q.lang.getLocale());
    console.log($q.platform.is);
  }
  //
  return { msg, printInfo };
});
