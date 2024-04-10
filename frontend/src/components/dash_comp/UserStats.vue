<template>
    <div class="box">
      <div class="user-stats">
        <div class="stat-item">
          <div class="stat-value">{{ classesAttended }}</div>
          <div class="stat-label">Classes Attended</div>
        </div>
        <div class="stat-item">
          <div class="stat-value">{{ upcomingUserClasses }}</div>
          <div class="stat-label">Upcoming Classes</div>
        </div>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref } from 'vue';
  
  const classesAttended = ref(25); // Example data, replace with real data later
  const upcomingUserClasses = ref(3); // Example data, replace with real data later
  
  // Logic for fetching real user stats will go here
  </script>
  
  <style scoped>
  .box {
    display: flex;
    align-items: center;
    justify-content: space-around;
    flex-wrap: wrap;
    border: 2px solid #ccc;
    border-radius: 8px;
    padding: 20px;
    height: auto;
  }
  
  .user-stats {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-around;
    width: 100%;
  }
  
  .stat-item {
    text-align: center;
  }
  
  .stat-value {
    font-size: 2em;
    font-weight: bold;
    color: #4CAF50;
    margin-bottom: 0.25em;
  }
  
  .stat-label {
    font-size: 1em;
    color: #555;
  }
  
  @media (max-width: 600px) {
    .user-stats {
      flex-direction: column;
    }
  
    .stat-item {
      margin-bottom: 1em;
    }
  }
  </style>
  