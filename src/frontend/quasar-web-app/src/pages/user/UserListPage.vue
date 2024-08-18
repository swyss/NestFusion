<template>
  <q-page>
    <q-card>
      <q-card-section>
        <div class="text-h6">User List</div>
      </q-card-section>

      <q-card-section>
        <q-list bordered>
          <q-item v-for="user in users" :key="user.id">
            <q-item-section>{{ user.name }}</q-item-section>
            <q-item-section>{{ user.email }}</q-item-section>
            <q-item-section side>
              <q-btn flat icon="edit" @click="editUser(user.id)" />
              <q-btn flat icon="delete" @click="handleDeleteUser(user.id)" />
            </q-item-section>
          </q-item>
        </q-list>
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script setup>
import { onMounted } from 'vue';
import { useUserStore } from 'stores/user/user-store';

const userStore = useUserStore();

const { fetchUsers, deleteUser, users } = userStore;

// Load users when the component is mounted
onMounted(() => {
  fetchUsers();
});

function editUser(userId) {
  // Implement logic to edit the user, e.g., navigate to an edit page
  console.log(`Editing user with ID: ${userId}`);
}

// Renamed local function to avoid conflict
function handleDeleteUser(userId) {
  if (confirm('Are you sure you want to delete this user?')) {
    deleteUser(userId);
  }
}
</script>
