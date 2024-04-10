<template>
  <div class="box">
    <div class="upcoming-classes-container">
      <h2 class="upcoming-classes-title">Upcoming Classes</h2>
      <ul class="upcoming-classes-list">
        <li v-for="(classInfo, index) in upcomingClasses" :key="index" class="class-item">
          <div><strong>Class:</strong> {{ classInfo.name }}</div>
          <div><strong>Date:</strong> {{ classInfo.date }}</div>
          <div><strong>Time:</strong> {{ classInfo.time }}</div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const upcomingClasses = ref([
  // Placeholder data for upcoming classes
  { name: 'Yoga Basics', date: '2024-04-12', time: '10:00 AM' },
  { name: 'Advanced Pilates', date: '2024-04-15', time: '12:00 PM' },
  { name: 'Cardio Kickboxing', date: '2024-04-18', time: '5:00 PM' },
  // Add more placeholder data as needed
]);

// Additional script logic to handle data passing can be added here later
</script>

<style scoped>
.box {
  display: block; /* Changed from flex to block for a vertical layout */
  max-height: 300px; /* Max height for scrollable area */
  overflow-y: auto; /* Enables vertical scrolling if content overflows */
  border: 2px solid #ccc;
  border-radius: 8px;
  padding: 20px;
}

.upcoming-classes-container {
  width: 100%;
}

.upcoming-classes-title {
  margin-bottom: 1rem;
  text-align: center;
}

.upcoming-classes-list {
  list-style-type: none; /* Remove list item bullets */
  padding: 0; /* Remove padding */
  margin: 0; /* Remove margin */
}

.class-item {
  margin-bottom: 1rem; /* Space between items */
  padding: 0.5rem;
  border-bottom: 1px solid #e5e7eb; /* Light border between items */
}
</style>
