<template>
    <div class="upcoming-appointments box">
      <div class="header">
        <div class="icon-calendar"></div>
        <h3>Upcoming Appointments</h3>
      </div>
      <!-- ... rest of your component -->
    </div>
  </template>
  
  <script setup>
  // ... your script setup
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
  
  .header {
    display: flex;
    align-items: center;
    /* ... */
  }

  .icon-calendar {
    display: inline-block;
    width: 20px; /* Width of the icon */
    height: 20px; /* Height of the icon */
    position: relative;
    border: 2px solid #666; /* Border color */
    border-radius: 3px; /* Rounded corners */
  }
  
  .icon-calendar:before {
    content: "";
    display: block;
    width: 60%; /* Width of the upper box to resemble the binding of a calendar */
    height: 4px; /* Height of the upper box */
    background-color: #666; /* Color of the upper box */
    position: absolute;
    top: -6px; /* Positioning just above the calendar */
    left: 50%;
    transform: translateX(-50%);
  }
  
  .icon-calendar:after {
    content: "";
    display: block;
    width: 100%; /* Width of the bottom part of the calendar */
    height: 4px; /* Height of the bottom part */
    background-color: #666; /* Color of the bottom part */
    position: absolute;
    bottom: 4px; /* Positioning from the bottom inside the calendar */
    left: 0;
  }
  
  
  /* Play Button Icon Styles */
  .start-session-button:before {
    .start-session-button:before {
        content: "";
        display: inline-block;
        width: 0;
        height: 0;
        border-top: 7px solid transparent; /* Creates the top part of the triangle */
        border-bottom: 7px solid transparent; /* Creates the bottom part of the triangle */
        border-left: 10px solid #666; /* Creates the base of the triangle, which is the 'play' symbol */
        vertical-align: middle;
        margin-right: 5px; /* Adjust the spacing between icon and button text if needed */
      }
      
  }
  
  /* ... other styles ... */
  </style>
  