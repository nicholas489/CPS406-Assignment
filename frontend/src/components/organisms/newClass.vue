<template>
    <div class="max-w-md mx-auto my-10 p-6 border border-gray-300 rounded-lg bg-white shadow-md">
      <h2 class="text-lg font-semibold text-center mb-6">Event Submission</h2>    
  
      <div class="mb-4">
        <label for="eventName" class="block text-sm font-medium text-gray-700">Event Name:</label>
        <input type="text" id="eventName" v-model="event.name" class="mt-1 block w-full border-gray-300 rounded-md shadow-sm" />
      </div>

      <div class="mb-4">
        <label for="coachEmail" class="block text-sm font-medium text-gray-700">Coach Email:</label>
        <input type="email" id="coachEmail" v-model="event.coachEmail" class="mt-1 block w-full border-gray-300 rounded-md shadow-sm" />
      </div>
  
      <div class="mb-4">
        <label for="location" class="block text-sm font-medium text-gray-700">Location:</label>
        <input type="text" id="location" v-model="event.location" class="mt-1 block w-full border-gray-300 rounded-md shadow-sm" />
      </div>

      <div class="mb-4">
        <label for="dateTime" class="block text-sm font-medium text-gray-700">Date and Time:</label>
        <input type="datetime-local" id="dateTime" v-model="event.dateTime" class="mt-1 block w-full border-gray-300 rounded-md shadow-sm" />
      </div>
  
      <div class="mb-4">
        <label for="cost" class="block text-sm font-medium text-gray-700">Cost:</label>
        <input type="number" id="cost" v-model="event.cost" class="mt-1 block w-full border-gray-300 rounded-md shadow-sm" min="0" />
      </div>
  
      <div class="mb-4">
        <label for="eventExpenses" class="block text-sm font-medium text-gray-700">Event Expenses:</label>
        <input type="number" id="eventExpenses" v-model="event.eventExpenses" class="mt-1 block w-full border-gray-300 rounded-md shadow-sm" min="0" />
      </div>
  
      <div class="mb-4">
        <label for="coachExpenses" class="block text-sm font-medium text-gray-700">Coach Expenses:</label>
        <input type="number" id="coachExpenses" v-model="event.coachExpenses" class="mt-1 block w-full border-gray-300 rounded-md shadow-sm" min="0" />
      </div>
  
      <button type="submit" @click="handleSubmit" class="w-full text-white bg-green-500 hover:bg-green-600 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-opacity-50 rounded-md shadow-sm py-2 my-4">
        Submit Event
      </button>
    </div>
  </template>
  
  <script>
  import { useRouter } from 'vue-router';
  
  export default {
    data() {
      return {
        event: {
          name: '',
          coachEmail: '',
          location: '',
          cost: 0,
          eventExpenses: 0,
          coachExpenses: 0,
          dateTime: '',      
        },
      };
    },
    methods: {
      handleSubmit() {
        // Here you would handle the form submission
        console.log('Event data:', this.event);
        // Add logic to send data to backend and navigate to dashboard on success
      },
    },
  };
  </script>
  
  <style scoped>
    /* Your existing styles */
  </style>
  