import { ref } from 'vue';

// Create refs to control dialog visibility
const isLoginDialogVisible = ref(false);
const isRegisterDialogVisible = ref(false);

// Function to show the login dialog
const showLoginDialog = () => {
  isLoginDialogVisible.value = true;
};

// Function to hide the login dialog
const hideLoginDialog = () => {
  isLoginDialogVisible.value = false;
};

// Function to show the register dialog
const showRegisterDialog = () => {
  isRegisterDialogVisible.value = true;
};

// Function to hide the register dialog
const hideRegisterDialog = () => {
  isRegisterDialogVisible.value = false;
};

export {
  isLoginDialogVisible,
  isRegisterDialogVisible,
  showLoginDialog,
  hideLoginDialog,
  showRegisterDialog,
  hideRegisterDialog
};
