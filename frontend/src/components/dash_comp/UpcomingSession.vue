<template>
  <div class="box">
    <div class="upcoming-classes-container">
      <h2 class="upcoming-classes-title">Upcoming Classes</h2>
      <ul class="upcoming-classes-list">
        <li v-for="(classInfo, index) in upcomingClasses" :key="index" class="class-item">
          <div class="class-name"><strong>Class:</strong> {{ classInfo.name }}</div>
          <div class="class-coach"><strong>Coach:</strong> {{ classInfo.coach }}</div>
          <div class="class-date"><strong>Date:</strong> {{ classInfo.date }}</div>
          <div class="class-time"><strong>Time:</strong> {{ classInfo.time }}</div>
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const upcomingClasses = ref([
  { name: 'Yoga Basics', coach: 'John Doe', date: '2024-04-12', time: '10:00 AM' },
  { name: 'Advanced Pilates', coach: 'Jane Smith', date: '2024-04-15', time: '12:00 PM' },
  { name: 'Cardio Kickboxing', coach: 'Emily Johnson', date: '2024-04-18', time: '5:00 PM' },
  // Add more classes with their respective coaches as needed
]);

// Additional script logic to handle dynamic data fetching will go here later
</script>

<style scoped>
.box {
  display: block;
  max-height: 300px;
  overflow-y: auto;
  border: 2px solid #ccc;
  border-radius: 8px;
  padding: 20px;
  background-color: #f9fafb; /* Light background color for contrast */
}

.upcoming-classes-container {
  width: 100%;
}

.upcoming-classes-title {
  margin-bottom: 1rem;
  text-align: center;
}

.upcoming-classes-list {
  list-style-type: none;
  padding: 0;
  margin: 0;
}

.class-item {
  margin-bottom: 1rem;
  padding: 0.75rem;
  border-bottom: 1px solid #e5e7eb;
  background-color: #fff; /* White background for each class item */
  border-radius: 4px; /* Rounded corners for class items */
  box-shadow: 0 2px 4px rgba(0,0,0,0.05); /* Subtle shadow for depth */
}

.class-item:last-child {
  border-bottom: none; /* Remove bottom border for the last item */
}

.class-name, .class-coach, .class-date, .class-time {
  color: #333; /* Dark text for readability */
  margin: 4px 0; /* Margin for spacing between lines */
}

.class-name {
  color: #28bc84; /* Use accent color for class name */
  font-weight: bold; /* Bold font for emphasis */
}
</style>
