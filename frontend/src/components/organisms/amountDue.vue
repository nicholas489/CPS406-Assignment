<template>
  <div class="flex flex-col gap-4 payment-form">
    <div class="font-bold">Total Amount Due: ${{ totalAmountDue }}</div>
    <div class="flex items-center gap-2">
      <label for="amount" class="block">Enter amount to pay:</label>
      <input type="text" id="amount" v-model="paymentAmount" @input="validateAmount" class="w-32 border-gray-300 rounded-md shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50" placeholder="Amount" />
    </div>
    <Button label="Proceed to Payment" @click="submitPayment" />
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import Button from 'primevue/button';
import { useRouter } from 'vue-router';
import { useToast } from 'primevue/usetoast';

const toast = useToast();
const totalAmountDue = 1000; // TODO: Backend need to change this value here
const paymentAmount = ref('');
const router = useRouter();

watch(paymentAmount, (newValue, oldValue) => {
  if (newValue !== '' && isNaN(Number(newValue))) {
    toast.add({
      severity: 'error',
      summary: 'Invalid Amount',
      detail: 'Please enter a valid number for the amount to pay',
      life: 3000
    });
    paymentAmount.value = oldValue; // Revert to the old value
  }
});

function validateAmount(event: any) {
 // Replace non-digits with an empty string
  paymentAmount.value = event.target.value.replace(/[^0-9.]+/g, '');
}

function submitPayment() {
  if (paymentAmount.value !== '' && !isNaN(Number(paymentAmount.value))) {
    router.push({ name: 'payment', query: { amount: paymentAmount.value } });
  } else {
    toast.add({
      severity: 'error',
      summary: 'Invalid Amount',
      detail: 'Please enter a valid number for the amount to pay',
      life: 3000
    });
  }
}
</script>

<style scoped>

</style>