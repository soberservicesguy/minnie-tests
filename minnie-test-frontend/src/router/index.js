import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home'
import OrdersComponent from '../views/Orders'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/home',
    name: 'Home',
    component: Home,
  },
  {
    path: '/orders',
    name: 'Orders',
    component: OrdersComponent,
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
})

export default router
