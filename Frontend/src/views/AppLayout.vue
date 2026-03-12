<template>
  <div class="app-container">
    <div class="background-blob blob-1"></div>
    <div class="background-blob blob-2"></div>

    <aside class="sidebar">
      <div class="sidebar-header">
        <i class="pi pi-chart-line logo-icon"></i>
        <span class="logo-text">Fitness Tracker</span>
      </div>

      <nav class="nav-links">
       <router-link 
  v-if="store.user?.role === 'Admin'" 
  to="/app/admin" 
  class="nav-item"
>
    <i class="pi pi-home"></i>
    <span>Dashboard</span>
</router-link>
        
        <router-link to="/app/about" class="nav-item">
          <i class="pi pi-calendar"></i>
          <span>Activities</span>
        </router-link>

        <router-link to="/app/leaderboard" class="nav-item">
          <i class="pi pi-trophy"></i>
          <span>Leaderboard</span>
        </router-link>

        <router-link v-if="store.user?.role === 'Admin'" to="/app/admin" class="nav-item admin-link">
          <i class="pi pi-shield"></i>
          <span>Admin Panel</span>
        </router-link>
        <router-link v-if="store.user?.role === 'User'" to="/app/myBookings" class="nav-item">
  <i class="pi pi-calendar-check"></i>
  <span>My Bookings</span>
</router-link>
      </nav>

      <div class="sidebar-footer">
        <button @click="logout" class="logout-btn">
          <i class="pi pi-power-off"></i>
          <span>Logout</span>
        </button>
      </div>
    </aside>

    <main class="main-content">
      <header class="top-bar">
        <div class="user-info">
          <span class="welcome-msg">Welcome back, <strong>{{ store.user?.firstName + ' ' + store.user?.lastName || 'User' }}</strong></span>
        </div>
      </header>
      
      <div class="page-container">
        <router-view />
      </div>
    </main>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router';
import { useAuthStore } from '../store/login';

const store = useAuthStore();
const router = useRouter();

const logout = () => {
  console.log("Logging out");
  store.logout();
  router.push("/");
};
</script>

<style scoped>
/* ===============================
   Layout Structure
================================ */
.app-container {
  display: flex;
  min-height: 100vh;
  background: #0f172a; /* Deep navy from login */
  color: #f8fafc;
  position: relative;
  overflow: hidden;
}

/* Reusing background blobs from Login */
.background-blob {
  position: absolute;
  width: 600px;
  height: 600px;
  border-radius: 50%;
  filter: blur(120px);
  opacity: 0.15;
  z-index: 0;
}
.blob-1 { background: #22c55e; top: -200px; right: -100px; }
.blob-2 { background: #06b6d4; bottom: -200px; left: -100px; }

/* ===============================
   Glassmorphism Sidebar
================================ */
.sidebar {
  width: 260px;
  background: rgba(30, 41, 59, 0.7);
  backdrop-filter: blur(20px);
  border-right: 1px solid rgba(255, 255, 255, 0.1);
  display: flex;
  flex-direction: column;
  z-index: 10;
}

.sidebar-header {
  padding: 2rem;
  display: flex;
  align-items: center;
  gap: 12px;
}

.logo-icon {
  font-size: 1.5rem;
  color: #22c55e;
}

.logo-text {
  font-weight: 800;
  font-size: 1.2rem;
  letter-spacing: -0.5px;
}

/* ===============================
   Navigation Links
================================ */
.nav-links {
  flex: 1;
  padding: 1rem;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 16px;
  border-radius: 12px;
  color: #94a3b8;
  text-decoration: none;
  transition: all 0.3s ease;
}

.nav-item i { font-size: 1.1rem; }

.nav-item:hover {
  background: rgba(255, 255, 255, 0.05);
  color: #f8fafc;
}

/* Router Link Active Class */
.router-link-active {
  background: rgba(34, 197, 94, 0.15) !important;
  color: #22c55e !important;
  font-weight: 600;
}

.admin-link {
  margin-top: 20px;
  border-top: 1px solid rgba(255, 255, 255, 0.05);
  padding-top: 20px;
}

/* ===============================
   Main Content Area
================================ */
.main-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  z-index: 1;
  background: rgba(15, 23, 42, 0.6);
}

.top-bar {
  height: 70px;
  padding: 0 2rem;
  display: flex;
  align-items: center;
  justify-content: flex-end;
  border-bottom: 1px solid rgba(255, 255, 255, 0.05);
}

.user-info {
  display: flex;
  align-items: center;
  gap: 15px;
}

.role-badge {
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 0.75rem;
  font-weight: 700;
  text-transform: uppercase;
}

.role-badge.admin { background: #06b6d4; color: #000; }
.role-badge.user { background: #22c55e; color: #000; }

.page-container {
  padding: 2rem;
  flex: 1;
  overflow-y: auto;
}

/* ===============================
   Logout Button
================================ */
.sidebar-footer {
  padding: 1rem;
}

.logout-btn {
  width: 100%;
  padding: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  background: rgba(239, 68, 68, 0.1);
  color: #ef4444;
  border: 1px solid rgba(239, 68, 68, 0.2);
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.logout-btn:hover {
  background: #ef4444;
  color: white;
}
</style>