<template>
    <div class="box">
        <div class="booking-container">
            <h2 class="booking-title">RSVP Event</h2>
            <form @submit.prevent="bookClass" class="booking-form">
                <div class="form-group">
                    <label for="classSelection">Select Event:</label>
                    <select id="classSelection" v-model="selectedClass" class="form-control">
                        <option
                            v-for="(classInfo, index) in upcomingClasses"
                            :key="index"
                            :value="classInfo">
                            {{ classInfo.name }} - {{ formatDate(classInfo.CreatedAt) }} - ${{ classInfo.cost }}
                        </option>
                    </select>
                </div>
                <button type="submit" class="submit-button">Confirm Booking</button>
            </form>
        </div>
        <Toast/>
    </div>
</template>

<script setup lang="ts">
import Toast from 'primevue/toast';
import {onMounted, ref} from 'vue';
import {useAuthStore} from "@/stores/authStore";
import {useToast} from "primevue/usetoast";
import type {Ref} from 'vue'
import type {Event} from '@/types';

const toast = useToast();
const authStore = useAuthStore();
const upcomingClasses: Ref<Event[]> = ref([]);
let selectedClass: Ref<Event>;
onMounted(async () => {
    const response = await fetch('/api/event');
    upcomingClasses.value = await response.json();
    selectedClass = ref(upcomingClasses.value[0]);
});
const formatDate = (string: string) => {
    const date = new Date(string);
    date.setDate(date.getDate() + 7);
    return new Intl.DateTimeFormat('en-US', {
        year: 'numeric',
        month: 'long',
        day: '2-digit',
        hour: 'numeric',
        minute: '2-digit',
        second: '2-digit',
        hour12: true,
    }).format(date);
}

const bookClass = async () => {
    const event: Event = selectedClass.value;
    const rawResponse = await fetch('/api/event/join', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({email: authStore.email, event_id: event.ID})
    });
    const content = await rawResponse.json();
    if (rawResponse.ok) {
        toast.add({severity: 'success', summary: 'Booking', detail: 'Successfully RSVP\'d', life: 3000});
    } else {
        toast.add({severity: 'error', summary: 'Booking', detail: content.error, life: 3000});
    }
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
