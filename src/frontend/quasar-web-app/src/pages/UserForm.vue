<template>
  <q-card>
    <q-card-section>
      <div class="text-h6">{{ isEdit ? 'Edit User' : 'Create User' }}</div>
    </q-card-section>

    <q-card-section>
      <q-input v-model="form.name" label="Name" />
      <q-input v-model="form.email" label="Email" type="email" />
    </q-card-section>

    <q-card-actions align="right">
      <q-btn flat label="Cancel" @click="cancel" />
      <q-btn flat label="Save" color="primary" @click="save" />
    </q-card-actions>
  </q-card>
</template>

<script setup>
import { ref } from 'vue';
import { useUserStore } from 'src/stores/user-store';

const userStore = useUserStore();

const form = ref({
  name: '',
  email: ''
});

const isEdit = ref(false);

// Populate the form if editing an existing user
function loadUserForEdit(user) {
  isEdit.value = true;
  form.value = { ...user };
}

function save() {
  if (isEdit.value) {
    userStore.updateUser(userStore.user.id, form.value);
  } else {
    userStore.createUser(form.value);
  }
}

function cancel() {
  // Implement logic to cancel the operation
}
</script>
