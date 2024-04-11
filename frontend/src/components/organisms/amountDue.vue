

<script setup lang="ts">
import {onMounted, ref, watch} from 'vue';
import { useRouter } from 'vue-router';
import { useToast } from 'primevue/usetoast';
import { useAuthStore } from "@/stores/authStore";

const toast = useToast();
const router = useRouter();
const balance = ref(0); // Placeholder value
const paymentAmount = ref('');
const authStore = useAuthStore();

const update = async () => {
    const response = await fetch(`/api/user/${authStore.id}`);
    const body = await response.json();
    balance.value = body.balance;
};

watch(() => authStore.isAuthenticated, async (newVal) => {
    if (newVal) {
        await update();
    }
}, {immediate: true});

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
  // Check if the payment amount is empty
  if (paymentAmount.value == '') {
    toast.add({
      severity: 'warn', // Using 'warn' severity for an empty field warning
      summary: 'Amount Required',
      detail: 'Please enter an amount to proceed with the payment',
      life: 3000
    });
  } else if (!isNaN(Number(paymentAmount.value))) {
    // Proceed with the payment if the amount is a valid number
    router.push({ name: 'payment', query: { amount: paymentAmount.value } });
  } else {
    // Display an error if the entered amount is not a valid number
    toast.add({
      severity: 'error',
      summary: 'Invalid Amount',
      detail: 'Please enter a valid number for the amount to pay',
      life: 3000
    });
  }
}
</script>

<template>
  <div class="flex flex-col gap-4 payment-form">
    <div class="font-bold">Balance: ${{ balance }}</div>
    <div class="flex items-center gap-2">
      <label for="amount" class="block">Enter amount to pay:</label>
      <input type="text" id="amount" v-model="paymentAmount" @input="validateAmount" class="w-32 border-gray-300 rounded-md shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50" placeholder="Amount" />
    </div>
    <Button  class='payment-button' label="Proceed to Payment" @click="submitPayment" />
  </div>
</template>
<style scoped>
/* Add more styles as needed */
</style>

