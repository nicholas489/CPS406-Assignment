<template>
    <div class="box">
      <p>Start a Session</p>
      <div class="button-container">
        <button @click="goToCreateSession" class="go-button">Go -></button>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { useRouter } from 'vue-router';
  
  const router = useRouter();
  
  function goToCreateSession() {
    router.push({ name: 'create-class' });
  }
  </script>
  
  <style scoped>
  .box {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    border: 2px solid #ccc;
    border-radius: 8px;
    height: 300px;
    padding: 20px;
    text-align: center; /* Center align the text */
  }
  
  .button-container {
    width: 100%; /* Take the full width of the parent */
    display: flex;
    justify-content: center; /* Center the button horizontally */
    padding-top: 20px; /* Optional: add space between the text and the button */
  }
  
  .go-button {
    padding: 10px 15px;
    background-color: #28bc84;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }
  </style>
  