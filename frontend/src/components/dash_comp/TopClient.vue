<template>
  <div class="box">
    <div class="top-clients-container">
      <h2 class="top-clients-title">Top Clients</h2>
      <ul class="top-clients-list">
        <li v-for="(member, index) in topClients" :key="index">
          <span>{{ member.name }} - {{ member.attendanceCount }} visits</span>
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const topClients = ref([
  // Fake members data
  { name: 'Alex Johnson', attendanceCount: 34 },
  { name: 'Morgan Smith', attendanceCount: 28 },
  { name: 'Jamie Doe', attendanceCount: 25 },
  // Add more fake members as needed
]);

// More script logic can be added here later
</script>

<style scoped>
.box {
  display: flex;
  flex-direction: column;
  align-items: center; 
  justify-content: flex-start; /* Align to start for the scrolling content */
  border: 2px solid #ccc;
  border-radius: 8px;
  height: 300px; /* Fixed height to enable scrolling */
  padding: 20px;
  overflow-y: auto; /* Enables vertical scrolling */
  background-color: white;
}

.top-clients-container {
  width: 100%;
}

.top-clients-title {
  margin-bottom: 1rem;
  text-align: center;
}

.top-clients-list {
  list-style-type: none;
  padding: 0;
  margin: 0; /* Remove default margins to prevent unexpected overflow */
  width: 100%; /* Ensure the list uses the full width of the box */
}

.top-clients-list li {
  margin-bottom: 0.5rem;
  background-color: #f3f3f3;
  padding: 0.5rem;
  border-radius: 4px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}
</style>

