<script setup lang="ts">
import {onMounted, ref, watch} from 'vue';
import {useAuthStore} from "@/stores/authStore";

const authStore = useAuthStore();
const classesAttended = ref(69); // Example data, replace with real data later
const upcomingUserClasses = ref(0);
const update = async () => {
    const response = await fetch('/api/event');
    upcomingUserClasses.value = (await response.json()).length
    classesAttended.value = parseInt(await (await fetch(`/api/user/${authStore.id}/events/count`)).text())
}

watch(() => authStore.isAuthenticated, async (a) => {
    if (a) {
        await update();
    }
})

onMounted(async () => {
    await update();
})
// Logic for fetching real user stats will go here
</script>

<template>
    <div class="box">
        <div class="user-stats">
            <div class="stat-item">
                <div class="stat-value">{{ classesAttended }}</div>
                <div class="stat-lacbel">Classes Attended</div>
            </div>
            <div class="stat-item">
                <div class="stat-value">{{ upcomingUserClasses }}</div>
                <div class="stat-label">Upcoming Classes</div>
            </div>
        </div>
    </div>
</template>

<style scoped>
.box {
    display: flex;
    align-items: center;
    justify-content: space-around;
    flex-wrap: wrap;
    border: 2px solid #ccc;
    border-radius: 8px;
    padding: 20px;
    height: auto;
}

.user-stats {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-around;
    width: 100%;
}

.stat-item {
    text-align: center;
}

.stat-value {
    font-size: 2em;
    font-weight: bold;
    color: #4CAF50;
    margin-bottom: 0.25em;
}

.stat-label {
    font-size: 1em;
    color: #555;
}

@media (max-width: 600px) {
    .user-stats {
        flex-direction: column;
    }

    .stat-item {
        margin-bottom: 1em;
    }
}
</style>
  