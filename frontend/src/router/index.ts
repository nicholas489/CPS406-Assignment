import { createRouter, createWebHistory } from 'vue-router'
import LoginView from '../views/LoginView.vue'
import DashboardView from '../views/UserDashboardView.vue'
import PaymentView from '@/views/PaymentView.vue'
import CoachDashboardView from "@/views/CoachDashboardView.vue";
import StartSessionView from "@/views/StartSessionView.vue";


const router = createRouter({
  history: createWebHistory("/"),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: LoginView
    },
    {
      path: '/user/dashboard',
      name: 'dashboard-user',
      component: DashboardView
    },
    {
      path: '/coach/dashboard',
      name: 'dashboard-coach',
      component: CoachDashboardView
    },
    {
      path: '/payment',
      name: 'payment',
      component: PaymentView
    },
    {
      path: '/create-session',
      name: 'create-session',
      component: StartSessionView
    }
  ]
},

)

export default router
