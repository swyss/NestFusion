<template>
  <q-card>
    <q-card-section>
      <div class="text-h6">{{ labels.PROFILE_TITLE }}</div>
    </q-card-section>
    <q-card-section>
      <q-input v-model="localUser.name" :label="labels.NAME_LABEL" @input="updateName"/>
      <q-input v-model="localUser.email" :label="labels.EMAIL_LABEL" disabled type="email"/>
      <q-input v-model="password" :label="labels.NEW_PASSWORD_LABEL" type="password"/>
    </q-card-section>
    <q-card-actions align="right">
      <q-btn :label="labels.UPDATE_PROFILE_LABEL" :loading="loading" color="primary" @click="updateProfile"/>
    </q-card-actions>
  </q-card>
</template>

<script>
import {reactive, ref, watch} from 'vue';

const labels = {
  PROFILE_TITLE: "User Profile",
  NAME_LABEL: "Name",
  EMAIL_LABEL: "Email",
  NEW_PASSWORD_LABEL: "New Password",
  UPDATE_PROFILE_LABEL: "Update Profile"
};

export default {
  props: {
    user: Object,
    loading: Boolean,
    updateProfile: Function
  },
  setup(props, {emit}) {
    const localUser = reactive({...props.user});
    const password = ref('');

    watch(() => props.user, (newUser) => {
      Object.assign(localUser, newUser);
    });

    const updateName = (newName) => {
      emit('update:user', {...localUser, name: newName});
    };

    return {localUser, password, labels, updateName};
  }
};
</script>
