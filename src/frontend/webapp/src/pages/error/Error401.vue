<template>
  <q-page padding class="flex flex-center">
    <div class="text-center">
      <img src="public/statics/img/funny-lock.png" alt="Funny 401" style="width: 300px;" />
      <h1>401 - Nicht autorisiert</h1>
      <p>Hey, du bist nicht eingeloggt! Melde dich an, um fortzufahren.</p>
      <q-btn color="primary" @click="goLogin">Zum Login</q-btn>
    </div>
  </q-page>
</template>

<script setup>
defineOptions({
  name: 'ErrorNotAuthorized'
});
</script>
