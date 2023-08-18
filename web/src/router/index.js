import { createRouter, createWebHashHistory } from 'vue-router'

const router = createRouter({
  history: createWebHashHistory(), // hash模式：createWebHashHistory，history模式：createWebHistory
  routes: [
    {
      path: '/',
      redirect: '/dashboard'
    },
    {
      path: '/bmDashboard',
      name: 'bmDashboard',
      component: () => import(/* webpackChunkName: "introduce" */ '../views/BmDashboard.vue')
    },
    {
      path: '/dashboard',
      name: 'dashboard',
      component: () => import(/* webpackChunkName: "dashboard" */ '../views/Dashboard.vue')
    },
    {
      path: '/sophonTest',
      name: 'sophonTest',
      component: () => import(/* webpackChunkName: "dashboard" */ '../views/SophonTest.vue')
    },
  ]
})

export default router