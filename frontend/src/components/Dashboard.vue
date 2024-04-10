<!-- Dashboard.vue -->
<script setup lang="ts">
import Schedule from './dash_comp/UpcomingSession.vue'
import Booking from './dash_comp/Bookings.vue'
import Finance from './dash_comp/Finances.vue'
import BoxFour from './dash_comp/ClubStats.vue'
import Button from 'primevue/button'

import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/authStore'

const authStore = useAuthStore();
const router = useRouter();


const logout = () => {
  authStore.logout();
  router.push('login');

};


</script>

<template>
  <div class="dashboard">
    <h2>Welcome this is your Dashboard</h2>
    <div class="grid grid-cols-2 gap-4">
      <Schedule />
      <Booking />
      <Finance />
      <BoxFour /> <!-- Remove or rename later just for testing reasons -->
    </div>

    <div class="logout-container">
      <Button @click="logout" label="Logout"></Button>
    </div>

  </div>
</template>

<style scoped>
.dashboard {
  padding: 20px;
  display: flex;
  flex-direction: column;
  align-items: center; /* This will center your logout button */
}

.grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 20px;
  width: 100%; /* Set to the full width of .dashboard */
}

/* Add this to make sure your h2 is styled correctly */
h2 {
  padding-bottom: 10px;
  margin-bottom: 10px;
  text-align: center; /* Center the header text */
}

.logout-button-container {
  width: 100%; /* Full width to center the button with text-align */
  text-align: center; /* Center-aligns the button */
}

.logout-button {
  margin-top: 20px; 
  padding: 10px 20px;
  background-color: #f44336; 
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

/* Additional styling as needed */
</style>

