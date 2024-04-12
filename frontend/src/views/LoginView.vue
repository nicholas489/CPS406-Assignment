<script setup lang="ts">

import LoginForm from "@/components/organisms/loginForm.vue";
import Toast from "primevue/toast"
//if user is already logged in
import {useAuthStore} from "@/stores/authStore";
import {useRouter} from "vue-router";
import {onMounted} from "vue";
import {useToast} from "primevue/usetoast";

const toast = useToast();
const router = useRouter();
const authStore = useAuthStore();
if (authStore.isAuthenticated) {
    if (authStore.isCoach) {
        router.push({name: 'dashboard-coach'});
    } else {
        router.push({name: 'dashboard-user'});
    }
}
onMounted(async () => {
    for (let e of authStore.toasts) {
        toast.add({ severity: e[0], summary: e[1], detail: e[2], life: 3000 });
    }
    authStore.clearToast()
})
</script>

<template>
    <div class="flex justify-center items-center">
        <LoginForm/>
    </div>
    <Toast/>
</template>

<style scoped>

</style>