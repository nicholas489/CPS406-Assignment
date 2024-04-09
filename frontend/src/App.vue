<script setup lang="ts">
import {RouterView} from 'vue-router'
import Navbar from "@/components/organisms/navbar.vue";
import {onMounted} from "vue";
import {useAuthStore} from "@/stores/authStore";
onMounted(() => {
    let authStore = useAuthStore();
    authStore.initAuth();
});
</script>

<template>
    <div class="flex flex-col h-full">
        <Navbar />
        <RouterView/>
    </div>
</template>

<style scoped>
</style>
