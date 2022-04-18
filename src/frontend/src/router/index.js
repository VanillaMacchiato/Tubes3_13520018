import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/input',
    name: 'Input',
    component: () => import('../views/Input.vue')
  },
  {
    path: '/prediksi',
    name: 'Prediksi',
    component: () => import('../views/Prediksi.vue')
  },
  {
    path: '/hasil',
    name: 'Hasil',
    component: () => import('../views/Hasil.vue')
  },
  {
    path: '/tentang',
    name: 'Tentang',
    component: () => import('../views/Tentang.vue')
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router
