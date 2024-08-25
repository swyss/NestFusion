<template>
  <q-page class="admin-page">
    <q-card>
      <q-card-section>
        <div class="text-h6">Admin Management</div> <!-- Display the title "Admin Management" -->
      </q-card-section>
      <q-card-section>
        <q-table
          :columns="columns"
          :rows="userStore.users"
          row-key="id">
          <template v-slot:body-cell="props">
            <!-- Default slot for all body cells -->
            <q-td :props="props">
              <template v-if="props.col.name === 'actions'">
                <q-btn flat icon="delete" @click="onDeleteUser(props.row.id)"/> <!-- Button to delete a user -->
              </template>
              <template v-else>
                {{ props.row[props.col.name] }}
              </template>
            </q-td>
          </template>
        </q-table>
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script>
import {useUserStore} from 'stores/user/user-store'; // Import the user store
import {onMounted} from 'vue'; // Import the onMounted lifecycle hook

// Extracted columns constant
const COLUMNS = [
  {name: 'name', required: true, label: 'Name', align: 'left', field: row => row.name}, // Column for the user's name
  {name: 'email', required: true, label: 'Email', align: 'left', field: row => row.email}, // Column for the user's email
  {name: 'actions', label: 'Actions', align: 'center', field: 'actions'} // Column for the action buttons
];

// Extracted function for deleting a user
const deleteUser = async (userStore, userId) => {
  await userStore.deleteUser(userId); // Call the store action to delete a user
};

export default {
  setup() {
    const userStore = useUserStore(); // Initialize the user store

    const onDeleteUser = (userId) => deleteUser(userStore, userId);

    onMounted(async () => {
      await userStore.fetchUsers(); // Fetch the list of users when the component is mounted
    });

    return {userStore, columns: COLUMNS, onDeleteUser}; // Return the store, columns, and delete function to the template
  }
};
</script>
