<script setup lang="ts">

import SignUp from "@/components/organisms/sign_up.vue";
//if user is already logged in
import {useAuthStore} from "@/stores/authStore";
import {useRouter} from "vue-router";
    const router = useRouter();
const authStore = useAuthStore();
if (authStore.isAuthenticated) {
    if (authStore.isCoach) {
        router.push({name: 'dashboard-coach'});
    } else {
        router.push({name: 'dashboard-user'});
    }
}
</script>

<template>
    <div class="flex justify-center items-center">
        <SignUp/>
    </div>
</template>

<style scoped>

</style>