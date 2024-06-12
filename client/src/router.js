import { createRouter, createWebHistory } from 'vue-router'
import LiveTracking from './views/LiveTracking.vue'
import OrderHistory from "@/views/OrderHistory.vue";
import RoutePath from "@/views/RoutePath.vue";

export const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: '/iframe/live-location', component: LiveTracking },
        { path: '/iframe/v2/live-location', component: RoutePath },
        { path: '/iframe/order-history', component: OrderHistory }
    ],
})
