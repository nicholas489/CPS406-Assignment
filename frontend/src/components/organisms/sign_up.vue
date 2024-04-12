<template>
    <div class="flex flex-col">
      <h1 class="text-3xl text-center">Sign Up</h1>
      <form class="flex flex-col space-y-8 mt-6 w-full" @submit.prevent="submitForm">
        <FloatLabel>
          <InputText id="username" v-model="username" placeholder="Enter your email"/>
          <label for="username">Email</label>
        </FloatLabel>
          <FloatLabel>
          <InputText id="name" v-model="name" placeholder="Enter your name"/>
          <label for="name">Name</label>
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
  
  <script setup lang="ts">
  import { ref } from 'vue';
  import InputText from 'primevue/inputtext';
  import FloatLabel from 'primevue/floatlabel';
  import Button from 'primevue/button';
  import Toast from 'primevue/toast';
  import {useToast} from "primevue/usetoast";
  import {useAuthStore} from "@/stores/authStore";
  import {useRouter} from "vue-router";

  const toast = useToast();
  const authStore = useAuthStore();
  const router = useRouter();
  // States for form inputs
  const username = ref('');
  const password = ref('');
  const confirmPassword = ref('');
  const coach = ref(false);
  const name = ref('');
  
  // Placeholder function for form submission
  const submitForm = async () => {
    if (password.value !== confirmPassword.value) {
        toast.add({severity: 'error', summary: 'Sign Up', detail: 'Passwords do not match!', life: 3000});
        return;
    }
     const rawResponse = await fetch(`/api/signup/${coach.value ? 'coach' : 'user'}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            name: name.value,
            email: username.value,
            password: password.value
        })
    });
    if (rawResponse.ok) {
        await authStore.pushToast('success', 'Signup', 'Signup Successful')
        await router.push('/login')
    } else {
        const body = await rawResponse.json();
        toast.add({severity: 'error', summary: 'Signup', detail: body.error, life: 3000});
    }
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
  