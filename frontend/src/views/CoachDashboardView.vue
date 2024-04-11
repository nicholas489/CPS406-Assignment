<!-- Dashboard.vue -->
<script setup lang="ts">

import Finance from '@/components/dash_comp/Bill.vue'
import ClubStats from '@/components/dash_comp/ClubStats.vue'
import UpcomingSession from '@/components/dash_comp/UpcomingSession.vue'
import NewSession from '@/components/dash_comp/CreateSession.vue'
import TopClient from '@/components/dash_comp/TopClient.vue'
import OverView from '@/components/dash_comp/Overview.vue'
import Button from 'primevue/button'
import {useAuthStore} from "@/stores/authStore";
import {useRouter} from "vue-router";



const authStore = useAuthStore();
const router = useRouter();
const logout = () => {
  authStore.logout();
  router.push('login');
};

</script>

<template>
  <div class="dashboard">
    <div class="grid grid-cols-2 gap-4">
      <ClubStats/>
      <UpcomingSession />
      <NewSession/>
      <Finance/>
      
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
