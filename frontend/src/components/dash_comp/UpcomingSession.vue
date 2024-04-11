<template>
    <div class="box">
        <div class="upcoming-classes-container">
            <h2 class="upcoming-classes-title">Upcoming Classes</h2>
            <ul class="upcoming-classes-list">
                <li v-for="(classInfo, index) in upcomingClasses" :key="index" class="class-item">
                    <div><strong>Class:</strong> {{ classInfo.name }}</div>
                    <div><strong>Date/Time:</strong> {{ formatDate(classInfo.CreatedAt) }}</div>
                    <div><strong>Location:</strong> {{ classInfo.location }}</div>
                </li>
            </ul>
        </div>
    </div>
</template>

<script setup lang="ts">
import {onMounted, type Ref, ref} from 'vue';
    type Event = {
        ID: number;
        CreatedAt: string;
        UpdatedAt: string;
        DeletedAt: string | null;
        name: string;
        coach_email: string;
        location: string;
        cost: number;
        Users: [number];
        event_expenses: number;
        coach_expenses: number;
    };

const upcomingClasses: Ref<Event[]> = ref([]);
onMounted(async () => {
    const response = await fetch('/api/event');
    upcomingClasses.value = await response.json();
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

// Additional script logic to handle data passing can be added here later
</script>

<style scoped>
.box {
    display: block; /* Changed from flex to block for a vertical layout */
    max-height: 300px; /* Max height for scrollable area */
    overflow-y: auto; /* Enables vertical scrolling if content overflows */
    border: 2px solid #ccc;
    border-radius: 8px;
    padding: 20px;
}

.upcoming-classes-container {
    width: 100%;
}

.upcoming-classes-title {
    margin-bottom: 1rem;
    text-align: center;
}

.upcoming-classes-list {
    list-style-type: none; /* Remove list item bullets */
    padding: 0; /* Remove padding */
    margin: 0; /* Remove margin */
}

.class-item {
    margin-bottom: 1rem; /* Space between items */
    padding: 0.5rem;
    border-bottom: 1px solid #e5e7eb; /* Light border between items */
}
</style>
