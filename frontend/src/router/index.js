import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import(/* webpackChunkName: "about" */ '../views/CodyHouse.vue')
  },
  {
    path: '/cody',
    name: 'Cody',
    component: () => import(/* webpackChunkName: "about" */ '../views/CodyHouse.vue')
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
