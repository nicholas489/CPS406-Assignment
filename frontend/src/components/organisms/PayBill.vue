<template>
    <div class="flex flex-col gap-4 payment-form">
        <div class="font-bold">Outstanding Balance: ${{ ownerTotalAmountDue }}</div>
        <div class="flex items-center gap-2">
            <label for="owner-amount" class="block">Enter amount to pay:</label>
            <input type="text" id="owner-amount" v-model="ownerPaymentAmount" @input="validateAmount"
                   class="w-32 border-gray-300 rounded-md shadow-sm focus:border-indigo-300 focus:ring focus:ring-indigo-200 focus:ring-opacity-50"
                   placeholder="Amount"/>
        </div>
        <Button class='payment-button' label="Pay Bill" @click="submitOwnerPayment"/>
    </div>
</template>

<script setup lang="ts">
import {onMounted, ref, watch} from 'vue';
import {useRouter} from 'vue-router';
import {useToast} from 'primevue/usetoast';
import {useAuthStore} from "@/stores/authStore";

const authStore = useAuthStore()
const toast = useToast();
const router = useRouter();
const ownerTotalAmountDue = ref(authStore.amountOwed); // Use ref to make it reactive
const ownerPaymentAmount = ref('');

watch(() => authStore.amountOwed, (a) => {
    ownerTotalAmountDue.value = a;
})

function validateAmount(event: any) {
    // Replace non-digits with an empty string
    ownerPaymentAmount.value = event.target.value.replace(/[^0-9.]+/g, '');
}

function submitOwnerPayment() {
    const payment = parseFloat(ownerPaymentAmount.value);
    if (ownerPaymentAmount.value === '') {
        toast.add({
            severity: 'warn',
            summary: 'Amount Required',
            detail: 'Please enter an amount to proceed with the payment',
            life: 3000
        });
    } else if (!isNaN(payment) && payment <= ownerTotalAmountDue.value) {
        authStore.amountOwed -= parseInt(ownerPaymentAmount.value);
        toast.add({
            severity: 'success',
            summary: 'Payment Successful',
            detail: `A payment of $${payment} has been applied.`,
            life: 3000
        });
        // Optionally reset the payment amount field
        ownerPaymentAmount.value = '';
    } else {
        toast.add({
            severity: 'error',
            summary: 'Invalid Amount',
            detail: 'Please enter a valid number that does not exceed the total amount due',
            life: 3000
        });
    }
}

function viewDetailedBill() {
    // Logic to view the detailed bill
    router.push({name: 'detailed-bill'});
}
</script>

<style scoped>
.details-button {
    /* Your styles for the view detailed bill button */
    background-color: #ffffff; /* Light background for contrast */
    color: #333333; /* Dark text for readability */
    border: 2px solid #4CAF50; /* Use a border to make the button stand out */
    /* Other styles like padding, font-size, border-radius, etc. */
}

/* Ensure rest of the styles match your design system */
</style>
  