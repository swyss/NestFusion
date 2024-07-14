<template>
  <q-page padding class="flex flex-center">
    <div class="text-center">
      <img src="public/statics/img/funny-image.png" alt="Funny 404" style="width: 300px;" />
      <h1>404 - Seite nicht gefunden</h1>
      <p>Oops! Du scheinst dich verlaufen zu haben. Lass uns zurückgehen.</p>
      <q-btn color="primary" @click="goHome">Zurück zur Startseite</q-btn>
    </div>
  </q-page>
</template>

<script setup>
defineOptions({
  name: 'ErrorNotFound'
});
</script>
