<template>
  <div class="box">
    <div class="stats">
      <div class="stat-item">
        <div class="stat-value">{{ totalMembers }}</div>
        <div class="stat-label">Total Members</div>
      </div>
      <div class="stat-item">
        <div class="stat-value">{{ sessionsCompleted }}</div>
        <div class="stat-label">Sessions Completed</div>
      </div>
      <div class="stat-item">
        <div class="stat-value">{{ upcomingClasses }}</div>
        <div class="stat-label">Upcoming Classes</div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const totalMembers = ref(150); // Example data, replace with real data later
const sessionsCompleted = ref(1200); // Example data, replace with real data later
const upcomingClasses = ref(5); // Example data, replace with real data later

// Additional script logic to handle dynamic data fetching will go here
</script>

<style scoped>
.box {
  display: flex;
  align-items: center;
  justify-content: space-around; /* Changed to space-around for even spacing */
  flex-wrap: wrap; /* Allow the stats to wrap if needed */
  border: 2px solid #ccc;
  border-radius: 8px;
  padding: 20px;
  height: auto; /* Changed to auto to accommodate content */
}

.stats {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-around; /* Evenly distribute space around items */
  width: 100%;
}

.stat-item {
  text-align: center; /* Center the text of stat items */
}

.stat-value {
  font-size: 2em; /* Large font size for the statistic value */
  font-weight: bold; /* Bold font for emphasis */
  color: #4CAF50; /* Color for visual appeal */
  margin-bottom: 0.25em; /* Small margin below the value */
}

.stat-label {
  font-size: 1em; /* Standard font size for labels */
  color: #555; /* Darker text color for contrast */
}

@media (max-width: 600px) {
  .stats {
    flex-direction: column; /* Stack the items on small screens */
  }

  .stat-item {
    margin-bottom: 1em; /* Add margin for stacked layout */
  }
}
</style>
