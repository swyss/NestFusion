<script setup>
defineOptions({
  name: 'UserAccount'
});
import { ref } from 'vue'

const form = ref({
  username: '',
  first_name: '',
  last_name: '',
  email: '',
  password: ''
})

const save = () => {
  console.log('save', form.value)
}

</script>

<template>
  <q-page class="flex flex-center">
    <q-card inline style="width:500px; padding:30px">
      <q-card-section class="bg-blue-grey-1 text-white">
        <div class="text-left" >
          <q-input stack-label label="Username" v-model="form.username" />
          <q-input stack-label label="First Name" v-model="form.first_name" />
          <q-input stack-label label="Last Name" v-model="form.last_name" />
          <q-input stack-label label="Email" v-model="form.email" />
          <q-input hide-underline stack-label label="Password" v-model="form.password" type='password'/>
        </div>
      </q-card-section>
      <q-separator />
      <q-card-actions class="float-right">
        <q-btn no-caps color="blue" label="Save" @click='save()' />
      </q-card-actions>
    </q-card>
  </q-page>
</template>

<style scoped>

</style>


