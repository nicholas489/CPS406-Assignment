<template>
    <div class="flex flex-col">
      <h1 class="text-3xl text-center">Sign Up</h1>
      <form class="flex flex-col space-y-8 mt-6 w-full" @submit.prevent="submitForm">
        <FloatLabel>
          <InputText id="username" v-model="username" placeholder="Enter your email"/>
          <label for="username">Email</label>
        </FloatLabel>
        <FloatLabel>
          <InputText id="password" v-model="password" type="password" placeholder="Enter your password"/>
          <label for="password">Password</label>
        </FloatLabel>
        <FloatLabel>
          <InputText id="confirm-password" v-model="confirmPassword" type="password" placeholder="Confirm your password"/>
          <label for="confirm-password">Confirm Password</label>
        </FloatLabel>
        <div class="checkbox-container">
          <input type="checkbox" id="coach-checkbox" v-model="coach"/>
          <label for="coach-checkbox">Register as Coach?</label>
        </div>
        <Button label="Sign Up" type="submit"/>
      </form>
    </div>
    <Toast/>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  import InputText from 'primevue/inputtext';
  import FloatLabel from 'primevue/floatlabel';
  import Button from 'primevue/button';
  import Toast from 'primevue/toast';
  
  // States for form inputs
  const username = ref('');
  const password = ref('');
  const confirmPassword = ref('');
  const coach = ref(false);
  
  // Placeholder function for form submission
  const submitForm = async () => {
    // The sign-up logic will be implemented here by your backend team
  };
  </script>
  
  <style scoped>
  .checkbox-container {
    display: flex;
    align-items: center;
    gap: 0.5rem;
  }
  
  /* Additional styling as needed */
  </style>
  