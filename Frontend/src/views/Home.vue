<template>
  <div class="user-dashboard-container">
    <header class="dashboard-header glass-panel">
      <div class="welcome-text">
        <h1>Hello, {{ user.firstName }}! 👋</h1>
        <p>Ready to crush your goals today?</p>
      </div>
      <div class="user-actions">
        <Button label="Log Activity" icon="pi pi-plus" class="p-button-rounded p-button-success" @click="showLogModal = true" />
      </div>
    </header>

    <div class="stats-grid">
      <div class="stat-card glass-panel">
        <i class="pi pi-bolt"></i>
        <div class="stat-info">
          <span class="label">Total Calories</span>
          <span class="value">{{ stats.totalCalories }} kcal</span>
        </div>
      </div>
      <div class="stat-card glass-panel">
        <i class="pi pi-calendar"></i>
        <div class="stat-info">
          <span class="label">This Week</span>
          <span class="value">{{ stats.weeklyActivities }} Activities</span>
        </div>
      </div>
      <div class="stat-card glass-panel">
        <i class="pi pi-clock"></i>
        <div class="stat-info">
          <span class="label">Active Time</span>
          <span class="value">{{ stats.totalMinutes }} mins</span>
        </div>
      </div>
    </div>

    <section class="recent-section glass-panel">
      <div class="panel-header">
        <h3><i class="pi pi-history"></i> Your Recent Activities</h3>
        <Button label="View All" class="p-button-text p-button-sm" />
      </div>
      
      <DataTable :value="recentActivities" class="p-datatable-sm modern-table" responsiveLayout="scroll">
        <Column field="activityName" header="Activity"></Column>
        <Column field="duration" header="Duration (min)"></Column>
        <Column field="caloriesBurned" header="Calories"></Column>
        <Column field="date" header="Date"></Column>
        <Column header="Status">
          <template #body="slotProps">
             <Tag :value="slotProps.data.status" :severity="getStatusSeverity(slotProps.data.status)" />
          </template>
        </Column>
      </DataTable>
    </section>

    <Dialog v-model:visible="showLogModal" header="Log New Activity" :modal="true" class="glass-modal" :style="{ width: '400px' }">
        <p>Form coming soon...</p>
    </Dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';

// State
const showLogModal = ref(false);
const user = ref({ firstName: 'User' });
const stats = ref({
  totalCalories: 1250,
  weeklyActivities: 5,
  totalMinutes: 240
});

const recentActivities = ref([
  { activityName: 'Running', duration: 30, caloriesBurned: 300, date: '2026-02-17', status: 'Completed' },
  { activityName: 'Cycling', duration: 45, caloriesBurned: 400, date: '2026-02-16', status: 'Completed' },
  { activityName: 'Yoga', duration: 60, caloriesBurned: 150, date: '2026-02-15', status: 'Pending' }
]);

const getStatusSeverity = (status) => {
  switch (status) {
    case 'Completed': return 'success';
    case 'Pending': return 'warning';
    default: return 'info';
  }
};

onMounted(() => {
    // 1. Get user profile from localStorage or Pinia
    // 2. Fetch user-specific stats from Go Backend
});
</script>

<style scoped>
.user-dashboard-container {
  padding: 2rem;
  display: flex;
  flex-direction: column;
  gap: 2rem;
  background: #0f172a; /* Deep dark blue background */
  min-height: 100vh;
  color: white;
}

.glass-panel {
  background: rgba(255, 255, 255, 0.03);
  backdrop-filter: blur(12px);
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 16px;
  padding: 1.5rem;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1.5rem;
}

.stat-card {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.stat-card i {
  font-size: 2rem;
  color: #22c55e; /* Green accent */
  background: rgba(34, 197, 94, 0.1);
  padding: 1rem;
  border-radius: 12px;
}

.stat-info .label {
  display: block;
  font-size: 0.9rem;
  color: #94a3b8;
}

.stat-info .value {
  font-size: 1.5rem;
  font-weight: bold;
}

.panel-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
}
</style>