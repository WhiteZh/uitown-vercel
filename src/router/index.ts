import { createRouter, createWebHistory } from 'vue-router'
import BrowseView from '../views/BrowseView.vue'
import CodeView from '@/views/CodeView.vue'
import UserView from "@/views/UserView.vue";
import HomeView from "@/views/HomeView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/home',
      alias: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/browse/:category?',
      name: 'browse',
      component: BrowseView
    },
    {
      path: '/create',
      name: 'create',
      component: CodeView,
      meta: {
        mode: 'create',
      }
    },
    {
      path: '/view/:id',
      name: 'view',
      component: CodeView,
      meta: {
        mode: 'view'
      }
    },
    {
      path: '/user',
      name: 'user',
      component: UserView,
    }
  ]
})

export default router
