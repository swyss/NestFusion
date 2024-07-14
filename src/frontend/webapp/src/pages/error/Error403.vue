<template>
  <q-page padding class="flex flex-center">
    <div class="text-center">
      <img src="public/statics/img/funny-security.png" alt="Funny 403" style="width: 300px;" />
      <h1>403 - Zugriff verweigert</h1>
      <p>Entschuldigung, du hast nicht die notwendigen Rechte, um diese Seite zu sehen. Bitte gehe zurück.</p>
      <q-btn color="primary" @click="goHome">Zurück zur Startseite</q-btn>
    </div>
  </q-page>
</template>

<script setup>
defineOptions({
  name: 'ErrorAccessDenied'
});
</script>
