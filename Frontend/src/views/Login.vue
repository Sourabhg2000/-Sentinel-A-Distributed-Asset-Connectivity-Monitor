<template>
  <div class="dashboard-wrapper">
    <div class="mesh-bg"></div>

    <header class="stats-header">
      <div class="brand">
        <i class="pi pi-shield logo-icon"></i>
        <h1>SENTINEL<span>.OS</span></h1>
      </div>
      <div class="quick-stats">
        <div class="stat-item">
          <span class="label">OPERATIONAL</span>
          <span class="value text-cyan">{{ upCount }}</span>
        </div>
        <div class="stat-item">
          <span class="label">DEGRADED</span>
          <span class="value text-rose">{{ downCount }}</span>
        </div>
      </div>
      <div class="user-profile">
        <img :src="userAvatar" class="avatar" />
        <Button icon="pi pi-power-off" class="p-button-rounded p-button-text text-slate" @click="logout" />
      </div>
    </header>

    <main class="content-area">
      <section class="action-bar">
        <div class="input-container">
          <i class="pi pi-globe globe-icon"></i>
          <InputText 
            v-model="newUrl" 
            placeholder="Enter URL to monitor (e.g. https://mahadnyandeep.org/)" 
            class="url-input"
            @keyup.enter="addNewTarget"
          />
          <Button 
            label="Initialize Probe" 
            icon="pi pi-plus" 
            class="add-btn" 
            :loading="isAdding"
            @click="addNewTarget" 
          />
        </div>
      </section>

      <section class="monitor-grid">
        <div v-for="site in targets" :key="site.id" class="site-card">
          <div class="card-glow" :class="site.status === 'UP' ? 'glow-green' : 'glow-red'"></div>
          
          <div class="card-content">
            <div class="site-info">
              <div class="status-indicator">
                <span class="dot" :class="site.status === 'UP' ? 'bg-green' : 'bg-red'"></span>
                <span class="status-text">{{ site.status }}</span>
              </div>
              <h3 class="site-url">{{ site.url }}</h3>
            </div>

            <div class="metrics">
              <div class="metric">
                <span class="m-label">LATENCY</span>
                <span class="m-value">{{ site.latency }}ms</span>
              </div>
              <div class="metric">
                <span class="m-label">UPTIME</span>
                <span class="m-value">99.9%</span>
              </div>
            </div>

            <div class="card-actions">
              <Button icon="pi pi-chart-bar" class="p-button-text p-button-sm" />
              <Button icon="pi pi-trash" class="p-button-text p-button-danger p-button-sm" @click="removeTarget(site.id)" />
            </div>
          </div>
        </div>
      </section>
    </main>
  </div>
</template>

<style scoped>
.dashboard-wrapper {
  min-height: 100vh;
  background: #020617;
  color: #f8fafc;
  padding: 1.5rem;
  position: relative;
}

.mesh-bg {
  position: absolute;
  inset: 0;
  background: radial-gradient(circle at 50% -20%, rgba(6, 182, 212, 0.1), transparent);
  pointer-events: none;
}

/* --- Header & Stats --- */
.stats-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 2rem;
  background: rgba(15, 23, 42, 0.6);
  backdrop-filter: blur(10px);
  border-radius: 20px;
  border: 1px solid rgba(255, 255, 255, 0.05);
  margin-bottom: 2rem;
}

.brand h1 { font-size: 1.2rem; font-weight: 800; letter-spacing: 2px; }
.brand span { color: #06b6d4; }
.logo-icon { color: #06b6d4; margin-right: 10px; }

.quick-stats { display: flex; gap: 3rem; }
.stat-item { display: flex; flex-direction: column; }
.label { font-size: 0.65rem; color: #64748b; font-weight: 700; }
.value { font-size: 1.5rem; font-weight: 800; }

.avatar { width: 40px; height: 40px; border-radius: 50%; border: 2px solid #06b6d4; }

/* --- Action Bar --- */
.action-bar { margin-bottom: 3rem; display: flex; justify-content: center; }
.input-container {
  display: flex;
  align-items: center;
  background: #0f172a;
  border: 1px solid #1e293b;
  padding: 0.5rem 0.5rem 0.5rem 1.5rem;
  border-radius: 100px;
  width: 100%;
  max-width: 700px;
  box-shadow: 0 10px 30px rgba(0,0,0,0.3);
}

.url-input {
  background: transparent !important;
  border: none !important;
  color: white !important;
  flex-grow: 1;
  font-size: 1rem;
}

.add-btn {
  background: linear-gradient(135deg, #06b6d4, #2563eb) !important;
  border: none !important;
  border-radius: 100px !important;
  padding: 0.8rem 1.5rem !important;
}

/* --- Monitor Grid --- */
.monitor-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 1.5rem;
}

.site-card {
  position: relative;
  background: rgba(30, 41, 59, 0.4);
  border: 1px solid rgba(255,255,255,0.05);
  border-radius: 24px;
  padding: 1.5rem;
  overflow: hidden;
  transition: transform 0.3s ease;
}

.site-card:hover { transform: translateY(-5px); }

.card-glow {
  position: absolute;
  top: 0; left: 0; width: 100%; height: 4px;
}
.glow-green { background: #10b981; box-shadow: 0 4px 20px rgba(16, 185, 129, 0.4); }
.glow-red { background: #f43f5e; box-shadow: 0 4px 20px rgba(244, 63, 94, 0.4); }

.status-indicator { display: flex; align-items: center; gap: 8px; margin-bottom: 1rem; }
.dot { width: 8px; height: 8px; border-radius: 50%; }
.bg-green { background: #10b981; box-shadow: 0 0 10px #10b981; }
.bg-red { background: #f43f5e; box-shadow: 0 0 10px #f43f5e; }
.status-text { font-size: 0.7rem; font-weight: 800; color: #94a3b8; }

.site-url { font-size: 1.1rem; font-weight: 700; white-space: nowrap; overflow: hidden; text-overflow: ellipsis; }

.metrics { display: flex; justify-content: space-between; margin-top: 1.5rem; padding-top: 1rem; border-top: 1px solid rgba(255,255,255,0.05); }
.m-label { display: block; font-size: 0.6rem; color: #64748b; }
.m-value { font-size: 1rem; font-weight: 700; color: #cbd5e1; }
</style>

<script setup>
import { ref, computed, onMounted } from 'vue';
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import axios from 'axios';

// Mock Data for the vivid demo
const targets = ref([

]);

const newUrl = ref('');
const isAdding = ref(false);
const userAvatar = ref('https://ui-avatars.com/api/?name=Admin&background=06b6d4&color=fff');

const upCount = computed(() => targets.value.filter(t => t.status === 'UP').length);
const downCount = computed(() => targets.value.filter(t => t.status === 'DOWN').length);

onMounted(() => {
    fetchTargets();
});

const addNewTarget = async () => {
    if (!newUrl.value) return;
    
    isAdding.value = true;
    try {
        const response = await axios.post('http://localhost:8000/o/target', {
            url: newUrl.value,
            interval_seconds: 60 // Monitoring frequency
        }, { 
            withCredentials: true 
        });

        // Add the new target returned by MongoDB to your reactive list
        targets.value.unshift(response.data);
        newUrl.value = ''; 
    } catch (error) {
        console.error("Monitoring initialization failed:", error.response?.data || error.message);
    } finally {
        fetchTargets();
        isAdding.value = false;
    }
};

const fetchTargets = async () => {
    try {
        const response = await axios.post('http://localhost:8000/o/targets', { 
            withCredentials: true 
        });
        targets.value = response.data;
    } catch (error) {
        console.error("Failed to fetch targets:", error.response?.data || error.message);
    }
};

const removeTarget = (id) => {
  targets.value = targets.value.filter(t => t.id !== id);
};

const logout = () => {
  // Logic to clear JWT and redirect to login
  window.location.href = '/login';
};
</script>