<template>
  <div class="payment-container">
    <h2 class="payment-title">Pay Invoice</h2>

    <div class="payment-details">
      <!-- Payment method section -->
      <div class="payment-method">
        <!-- Card icons here -->
      </div>

      <!-- Payment amount section -->
      <div class="payment-amount">
        <label for="amount" class="block text-sm font-medium text-gray-700">Payment amount</label>
        <div class="mt-1 relative rounded-md shadow-sm">
          <input type="text" id="amount" class="form-input block w-full sm:text-sm border-gray-300" placeholder="Canadian Dollars" readonly />
        </div>
      </div>

      <!-- Name on card section -->
      <div class="input-group">
        <label for="cardName">Name on card</label>
        <input type="text" id="cardName" v-model="cardName" class="form-input" placeholder="Name on card">
      </div>

      <!-- Card number section -->
      <div class="input-group">
        <label for="cardNumber">Card number</label>
        <input type="text" id="cardNumber" v-model="cardNumber" class="form-input" placeholder="Card number" maxlength="16">
      </div>

      <!-- Expiry date and security code section -->
      <div class="flex">
        <div class="input-group flex-1">
          <label for="expiry">Expiry date</label>
          <input type="text" id="expiry" v-model="expiry" class="form-input" placeholder="MM/YY">
        </div>
        <div class="input-group flex-1">
          <label for="securityCode">Security code</label>
          <input type="text" id="securityCode" v-model="securityCode" class="form-input" placeholder="CVV" maxlength="3">
        </div>
      </div>

      <!-- ZIP/Postal code section -->
      <div class="input-group">
        <label for="zip">ZIP/Postal code</label>
        <input type="text" id="zip" v-model="postalCode" class="form-input" placeholder="ZIP Code">
      </div>

      <!-- PrimeVue Button for payment submission -->
      <Button label="Pay" @click="submitPayment" class="w-full mt-2" />
    </div>
    <Toast/>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import Button from 'primevue/button';
import Toast from 'primevue/toast';
import { useToast } from "primevue/usetoast";

const toast = useToast();

const cardName = ref('');
const cardNumber = ref('');
const expiry = ref('');
const securityCode = ref('');
const postalCode = ref('');

function submitPayment() {
  let errorMessage = '';

  if (!cardName.value) errorMessage += 'Name on card is required. ';
  // Update this line to check for exactly 16 digits
  if (!/^\d{16}$/.test(cardNumber.value)) errorMessage += 'Card number must be 16 digits and numeric. ';
  if (securityCode.value.length !== 3) errorMessage += 'Security code must be 3 digits. ';
  if (postalCode.value.length < 5 || postalCode.value.length > 8) errorMessage += 'Postal code must be between 5 and 8 characters. ';

  if (errorMessage) {
    toast.add({ severity: 'error', summary: 'Payment Error', detail: errorMessage.trim(), life: 3000 });
    return;
  }

  toast.add({ severity: 'success', summary: 'Payment Success', detail: 'Payment processed successfully', life: 3000 });
  // Implement payment processing logic here...
}

</script>

<style scoped>
.payment-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  max-width: 400px;
  margin: 2rem auto;
  padding: 2rem;
  border: 1px solid #ccc;
  border-radius: 0.5rem;
  background: #fff;
  box-shadow: 0 4px 8px rgba(0,0,0,0.1);
}

.payment-title {
  margin-bottom: 1rem;
}

.payment-method, .payment-amount, .input-group {
  width: 100%;
  margin-bottom: 1rem;
}

.form-input {
  padding: 0.5rem;
  border: 1px solid #ccc;
  border-radius: 0.375rem;
  width: 100%;
}

  .flex {
    display: flex;
    gap: 1rem;
  }

  .flex-1 {
    flex: 1;
  }

  .input-group label {
    display: block;
    margin-bottom: 0.5rem;
  }

  /* Import PrimeVue button styles if not globally imported */
  </style>
