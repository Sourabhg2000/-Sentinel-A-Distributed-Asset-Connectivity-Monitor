<template>
  <div class="activity-container">
    <ConfirmDialog></ConfirmDialog>
    <Toast />

    <header class="header-section glass-panel">
      <div>
        <h1 class="text-3xl font-bold">Activity Dashboard</h1>
        <p class="text-slate-400">Track your progress and book upcoming sessions</p>
      </div>
    </header>

    <div class="dashboard-grid">
      <aside class="sidebar-layout">
        <div class="glass-panel stat-card highlight-green">
          <span class="label">Total Calories</span>
          <h2 class="value">{{ totalCalories }} <small>kcal</small></h2>
          <div class="progress-container">
            <div class="progress-fill" :style="{ width: Math.min((totalCalories/2000)*100, 100) + '%' }"></div>
          </div>
          <p class="text-xs mt-2 text-slate-400">Target: 2000 kcal</p>
        </div>

        <div class="glass-panel stat-card mt-3">
          <span class="label">Sessions Completed</span>
          <h2 class="value">{{ completedSessionsCount }}</h2>
        </div>
      </aside>

      <main class="main-layout">
        <section class="glass-panel mb-4">
          <div class="flex justify-content-between align-items-center mb-4">
            <h3 class="m-0 text-xl"><i class="pi pi-calendar-plus text-green-400 mr-2"></i>Available Sessions</h3>
          </div>
          
          <div class="sessions-grid">
            <div v-for="session in availableSessions" :key="session.id" class="session-card">
              <div class="card-content">
                <Tag :value="session.category" severity="info" class="mb-2" />
                <h4 class="text-lg font-bold m-0">{{ session.activityName }}</h4>
                <p class="text-xs text-green-400 mt-1">{{ session.calories }} kcal/hr</p>
              </div>

              <Button 
                :label="session.isBooked ? 'Reserved' : 'Book Now'" 
                :icon="session.isBooked ? 'pi pi-check' : 'pi pi-bookmark'" 
                class="w-full mt-3"
                :class="session.isBooked ? 'p-button-secondary' : 'p-button-success'"
                :disabled="session.isBooked || session.slots === 0"
                @click="confirmBooking(session)" 
              />
            </div>
          </div>
        </section>

        <section class="glass-panel">
          <h3 class="mb-4"><i class="pi pi-history text-blue-400 mr-2"></i>Recent Activity Logs</h3>
          <DataTable :value="userLogs" :loading="loadingLogs" responsiveLayout="scroll" class="p-datatable-sm modern-table" :rows="5" paginator>
            <template #empty>No activities found.</template>
            <Column field="activityName" header="Activity"></Column>
            <Column field="status" header="Status">
              <template #body="slotProps">
                <Tag :value="slotProps.data.status" :severity="slotProps.data.status === 'Completed' ? 'success' : 'warning'" />
              </template>
            </Column>
            <Column field="bookedAt" header="Date">
              <template #body="slotProps">
                {{ new Date(slotProps.data.bookedAt).toLocaleDateString() }}
              </template>
            </Column>
          </DataTable>
        </section>
      </main>
    </div>

    <Dialog v-model:visible="showLogModal" header="Log Manual Activity" modal class="glass-modal" :style="{ width: '400px' }">
      <div class="flex flex-column gap-3 mt-2">
        <div class="field flex flex-column gap-2">
          <label>Select Activity</label>
          <Dropdown v-model="selectedActivity" :options="activityTypesList" optionLabel="activityName" placeholder="Choose activity..." />
        </div>
        <div class="field flex flex-column gap-2">
          <label>Duration (Minutes)</label>
          <InputNumber v-model="manualDuration" suffix=" mins" />
        </div>
        <Button label="Save Activity" icon="pi pi-save" class="p-button-success w-full mt-2" @click="saveActivityLog" :loading="isSaving" />
      </div>
    </Dialog>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from 'vue';
import { useConfirm } from "primevue/useconfirm";
import { useToast } from "primevue/usetoast";
// Import PrimeVue components if not registered globally
import Tag from 'primevue/tag';
import Button from 'primevue/button';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Dialog from 'primevue/dialog';
import Dropdown from 'primevue/dropdown';
import InputNumber from 'primevue/inputnumber';
import Toast from 'primevue/toast';
import ConfirmDialog from 'primevue/confirmdialog';

const toast = useToast();
const confirm = useConfirm();

// UI STATES
const showLogModal = ref(false);
const isSaving = ref(false);
const loadingLogs = ref(false);

// DATA STATES
const userLogs = ref([]);
const availableSessions = ref([]);
const activityTypesList = ref([]);
const selectedActivity = ref(null);
const manualDuration = ref(30);

// COMPUTE STATS
const totalCalories = computed(() => {
  // Logic to calculate calories from completed logs
  return userLogs.value
    .filter(log => log.status === 'Completed')
    .length * 250; // Mock calculation
});

const completedSessionsCount = computed(() => {
  return userLogs.value.filter(log => log.status === 'Completed').length;
});

// GET USER INFO
const getUserData = () => JSON.parse(sessionStorage.getItem("userData") || "{}");
const getAuthHeaders = () => ({
    "Content-Type": "application/json",
    "Authorization": `Bearer ${sessionStorage.getItem("auth_token")}`
});

onMounted(() => {
  fetchActivityTypes();
  fetchUserRecentLogs();
});

// 1. Fetch Logs from SG_Booking
const fetchUserRecentLogs = async () => {
    const user = getUserData();
    if (!user.userId) return;

    loadingLogs.value = true;
    try {
        const response = await fetch("http://localhost:8000/r/getMyBookings", {
            method: "POST",
            headers: getAuthHeaders(),
            body: JSON.stringify({ userId: user.userId }),
        });
        const data = await response.json();
        if (response.ok) {
            userLogs.value = Array.isArray(data) ? data : [];
        }
    } catch (err) {
        console.error("Error fetching logs:", err);
    } finally {
        loadingLogs.value = false;
    }
};

// 2. Fetch Session Types
const fetchActivityTypes = async () => {
    try {
        const response = await fetch("http://localhost:8000/r/activityTypes", {
            method: "POST",
            headers: getAuthHeaders(),
            body: JSON.stringify({}),
        });
        const data = await response.json();
        if (response.ok) {
            activityTypesList.value = data;
            availableSessions.value = data.map((item, index) => ({
                id: item.activityTypeId || index,
                activityName: item.activityName,
                category: item.activityType || 'General',
                slots: 10, 
                isBooked: false,
                calories: item.caloriesPerHr
            }));
        }
    } catch (err) {
        console.error("Network error:", err);
    }
};

// 3. Book Session and Refresh
const handleBooking = async (session) => {
    const user = getUserData();
    const payload = { 
        activityId: session.id,
        activityName: session.activityName,
        userInfo: user
    };

    try {
        const response = await fetch("http://localhost:8000/r/bookSession", {
            method: "POST",
            headers: getAuthHeaders(),
            body: JSON.stringify(payload),
        });

        if (response.ok) {
            session.isBooked = true;
            toast.add({ severity: 'success', summary: 'Booked', detail: `Reserved ${session.activityName}`, life: 3000 });
            // Refresh logs immediately to show new booking in table
            await fetchUserRecentLogs();
        }
    } catch (err) {
        console.error("Booking error:", err);
    }
};

const confirmBooking = (session) => {
    confirm.require({
        message: `Do you want to book ${session.activityName}?`,
        header: 'Confirm Booking',
        acceptClass: 'p-button-success',
        accept: () => handleBooking(session)
    });
};

// 4. Save Manual Log and Refresh
const saveActivityLog = async () => {
    if(!selectedActivity.value) return;
    
    isSaving.value = true;
    const user = getUserData();
    
    // Using the same structure as booking but marking it differently if your backend supports it
    const payload = {
        userId: user.userId,
        activityId: selectedActivity.value.activityTypeId,
        activityName: selectedActivity.value.activityName,
        duration: manualDuration.value,
        status: 'Completed' // Manual logs are usually finished activities
    };

    try {
        const response = await fetch("http://localhost:8000/r/completeActivity", { // Using the completion endpoint
            method: "POST",
            headers: getAuthHeaders(),
            body: JSON.stringify(payload),
        });

        if(response.ok) {
            toast.add({ severity: 'success', summary: 'Saved', detail: 'Activity logged successfully' });
            showLogModal.value = false;
            await fetchUserRecentLogs(); // Refresh the table
        }
    } catch (err) {
        console.error("Save error:", err);
    } finally {
        isSaving.value = false;
    }
};
</script>

<style scoped>
/* Keeping your existing styles... */
.activity-container { padding: 2rem; background: #0f172a; min-height: 100vh; color: #f8fafc; font-family: 'Inter', sans-serif; }
.glass-panel { background: rgba(255, 255, 255, 0.03); backdrop-filter: blur(12px); border: 1px solid rgba(255, 255, 255, 0.1); border-radius: 16px; padding: 1.5rem; }
.header-section { display: flex; justify-content: space-between; align-items: center; margin-bottom: 2rem; }
.dashboard-grid { display: grid; grid-template-columns: 280px 1fr; gap: 1.5rem; }
.stat-card.highlight-green { border-left: 5px solid #22c55e; }
.value { font-size: 2rem; font-weight: 800; margin: 0.5rem 0; }
.value small { font-size: 1rem; color: #64748b; }
.progress-container { height: 8px; background: rgba(255,255,255,0.1); border-radius: 4px; }
.progress-fill { height: 100%; background: #22c55e; border-radius: 4px; transition: width 0.5s ease; }
.sessions-grid { display: grid; grid-template-columns: repeat(auto-fill, minmax(220px, 1fr)); gap: 1rem; }
.session-card { background: rgba(255, 255, 255, 0.05); border: 1px solid rgba(255, 255, 255, 0.1); border-radius: 12px; padding: 1rem; display: flex; flex-direction: column; justify-content: space-between; transition: 0.3s; }
.session-card:hover { border-color: #22c55e; transform: translateY(-3px); }
:deep(.modern-table) .p-datatable-thead > tr > th { background: transparent !important; color: #64748b !important; border-bottom: 1px solid rgba(255,255,255,0.1); }
:deep(.modern-table) .p-datatable-tbody > tr { background: transparent !important; color: white !important; border-bottom: 1px solid rgba(255,255,255,0.05); }
</style>