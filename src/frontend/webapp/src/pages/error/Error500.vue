<template>
  <q-page padding class="flex flex-center">
    <div class="text-center">
      <img src="public/statics/img/funny-robot.png" alt="Funny 500" style="width: 300px;" />
      <h1>500 - Interner Serverfehler</h1>
      <p>Es sieht so aus, als ob unser Server gerade eine Pause macht. Bitte versuche es später noch einmal.</p>
      <q-btn color="primary" @click="goHome">Zurück zur Startseite</q-btn>
    </div>
  </q-page>
</template>

<script setup>
defineOptions({
  name: 'ErrorServer'
});
</script>
