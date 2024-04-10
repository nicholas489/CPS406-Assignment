<template>
  <div class="box">

<PaymentForm />

  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import PaymentForm from '../organisms/PayBill.vue';



</script>
  
  <style scoped>
    .box {
  display: flex;
  align-items: center;
  justify-content: center;
  border: 2px solid #ccc;
  border-radius: 8px;
  height: 300px;
  padding: 20px;
}
  
p {
  color: #333;
  font-size: 18px; /* Slightly larger text */
  margin-bottom: 24px; /* Space before the form */
}

/* Style the form components here if needed */

input[type="text"], /* Select text-type inputs */
input[type="number"] { /* Select number-type inputs */
  width: 100%; /* Full width */
  padding: 12px 20px;
  margin: 8px 0;
  display: inline-block;
  border: 1px solid #ccc;
  border-radius: 4px;
  box-sizing: border-box; /* Border box sizing */
}

button {
  width: 100%; /* Full width */
  background-color: #4CAF50; /* Green background */
  color: white;
  padding: 14px 20px;
  margin: 8px 0;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #45a049; /* Darker green on hover */
}

@media (min-width: 600px) { /* Responsive breakpoint */
  .box {
    padding: 60px; /* More padding on larger screens */
  }
}
</style>