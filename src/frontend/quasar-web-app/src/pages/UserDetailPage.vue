<template>
  <q-page>
    <q-card>
      <q-card-section>
        <div class="text-h6">User Details</div>
      </q-card-section>

      <q-card-section v-if="user">
        <div><strong>Name:</strong> {{ user.name }}</div>
        <div><strong>Email:</strong> {{ user.email }}</div>
      </q-card-section>

      <q-card-section v-else>
        <q-spinner size="30px" color="primary" />
        Loading...
      </q-card-section>

      <q-card-actions align="right">
        <q-btn flat label="Back" @click="goBack" />
      </q-card-actions>
    </q-card>
  </q-page>
</template>

<script setup>
import { onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useUserStore } from 'src/stores/user-store';

const route = useRoute();
const router = useRouter();
const userStore = useUserStore();

const { fetchUser, user } = userStore;

// Load the user based on the URL ID when the component is mounted
onMounted(() => {
  fetchUser(route.params.id);
});

function goBack() {
  router.back();
}
</script>
