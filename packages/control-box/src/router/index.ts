import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/trackview',
    name: 'trackview',
    component: () => import('@/views/TrackView.vue'),
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

export default router
