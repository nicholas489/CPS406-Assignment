<template>
  <div class="box">
    <div class="booking-container">
      <h2 class="booking-title">Book a Class</h2>
      <form @submit.prevent="bookClass" class="booking-form">
        <div class="form-group">
          <label for="classSelection">Select Class:</label>
          <select id="classSelection" v-model="selectedClass" class="form-control">
            <option 
              v-for="(classInfo, index) in upcomingClasses" 
              :key="index" 
              :value="classInfo">
                {{ classInfo.name }} - {{ classInfo.date }} - {{ classInfo.time }}
            </option>
          </select>
        </div>
        <button type="submit" class="submit-button">Confirm Booking</button>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const upcomingClasses = ref([
  { name: 'Yoga Basics', date: '2024-04-12', time: '10:00 AM' },
  { name: 'Advanced Pilates', date: '2024-04-15', time: '12:00 PM' },
  { name: 'Cardio Kickboxing', date: '2024-04-18', time: '5:00 PM' },
  // ...other classes
]);

const selectedClass = ref(upcomingClasses.value[0]); // Default to the first class

const bookClass = () => {
  // Logic to book the selected class
  // TODO: Need to make a toast here
  console.log('Booking confirmed for:', selectedClass.value);
  // Add your backend booking logic here
};
</script>

<style scoped>
.box {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
  border: 2px solid #ccc;
  border-radius: 8px;
  height: auto; /* Adjusted for content */
  padding: 20px;
  overflow-y: auto; /* For scrolling */
}

.booking-container {
  width: 100%;
}

.booking-title {
  margin-bottom: 1rem;
  text-align: center;
}

.form-group {
  margin-bottom: 1rem;
}

.form-control {
  width: 100%;
  padding: 0.5rem;
  margin-bottom: 1rem;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.submit-button {
  padding: 0.5rem 1rem;
  width: 100%;
  background-color: #28bc84; /* Green background */
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.submit-button:hover {
  background-color: #059669;
}
</style>
