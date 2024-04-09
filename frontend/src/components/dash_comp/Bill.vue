<template>
  <div class="box">
<p> THis is the bill you need to pay</p>


<PaymentForm />

  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import PaymentForm from '../organisms/amountDue.vue';

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
  </style>