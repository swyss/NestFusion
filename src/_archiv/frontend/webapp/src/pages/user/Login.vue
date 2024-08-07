<script setup>
defineOptions({
  name: 'UserLogin'
});
import { ref } from 'vue'

const form = ref({
  username: '',
  password: ''
})

const onSubmit = () => {
  console.log('onSubmit', form.value)
}

</script>

<template>
  <div class="fixed fixed-center">
    <form v-on:submit.prevent="onSubmit()">
      <q-card square  class="flex-center text-center" style="width: 400px; padding:50px">
        <q-card-section class="bg-blue-grey-1 text-white">
          <div class="text-left" >
            <q-input id="username" type="text" v-model.trim="form.username" stack-label label="User Name" required autofocus/>
            <q-input id="password" type="password" v-model="form.password" stack-label label="Password" required/><br>
          </div>
        </q-card-section>
        <q-separator />
        <q-card-actions align="around">
          <q-btn type="submit" class="fit" color="blue">Login</q-btn>
        </q-card-actions>
      </q-card>
    </form>
  </div>
</template>

<style scoped>

</style>
