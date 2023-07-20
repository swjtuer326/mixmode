import { createRouter, createWebHashHistory } from "vue-router";

const router = createRouter({
    history: createWebHashHistory(), // hash模式：createWebHashHistory，history模式：createWebHistory
    routes: [
      {
        path: '/',
        redirect: '/dashboard'
      },
      {
        path: '/introduce',
        name: 'introduce',
        component: () => import(/* webpackChunkName: "introduce" */ '@/views/about.vue')
      },
      {
        path: '/dashboard',
        name: 'dashboard',
        component: () => import(/* webpackChunkName: "dashboard" */ '@/views/dashboard.vue')
      },
    ]
  })
  
  export default router