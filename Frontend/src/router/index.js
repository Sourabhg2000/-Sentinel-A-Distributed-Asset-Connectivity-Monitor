import { createRouter, createWebHistory } from 'vue-router';
import { useAuthStore } from '../store/login';
import App from '../views/AppLayout.vue'; // Ensure this file exists

const routes = [
  {
    path: '/',
    name: 'Login',
    component: () => import('../views/Login.vue')
  },
  {
    path: '/signUp',
    name: 'SignUp',
    component: () => import('../views/SignUp.vue')
  },
  {
    path: '/app', // 🔹 Changed from '/' to '/app' for better structure
    component: App,
    children: [
      {
        path: 'home',
        name: 'Home',
        component: () => import('../views/Home.vue'),
        // meta: { requiresAuth: true }
      },
      {
        path: 'about',
        name: 'About',
        component: () => import('../views/aboutPage.vue'),
        meta: { requiresAuth: true }
      },
      {
        path: 'admin',
        name: 'Admin',
        component: () => import('../views/adminDashboard.vue'),
        meta: { requiresAuth: true }
      },
      {
        path: 'leaderBoard',
        name: 'leaderBoard',
        component: () => import('../views/DynamicTable.vue'),
        meta: { requiresAuth: true }
      },
      {
        path: 'myBookings',
        name: 'MyBookings',
        component: () => import('../views/MyBookings.vue'),
        meta: { requiresAuth: true }
      }
    ]
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

// 🔹 Protect routes that require authentication
router.beforeEach((to, from, next) => {
  const store = useAuthStore();
  const isAuthenticated = store.token || sessionStorage.getItem("auth_token");

  if (to.meta.requiresAuth && !isAuthenticated) {
    next("/");
  } else {
    next();
  }
});

export default router;
