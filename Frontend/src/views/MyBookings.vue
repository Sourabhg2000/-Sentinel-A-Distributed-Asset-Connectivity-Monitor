<template>
  <div class="bookings-container">
    <header class="page-header glass-panel">
      <div>
        <h1 class="text-2xl font-bold">My Scheduled Sessions</h1>
        <p class="text-slate-400">Mark your activities as complete to earn your calories!</p>
      </div>
    </header>

    <div class="bookings-grid mt-4">
      <div v-if="isLoading" class="text-center p-5 w-full">
        <i class="pi pi-spin pi-spinner text-4xl text-green-500"></i>
        <p class="mt-2 text-slate-400">Loading your schedule...</p>
      </div>

      <div v-else-if="!myBookings || myBookings.length === 0" class="glass-panel text-center p-5 w-full">
        <i class="pi pi-calendar-times text-4xl mb-3 text-slate-500"></i>
        <p>No active bookings found. Go to the Activity Page to book a session!</p>
      </div>

      <div v-else v-for="booking in myBookings" :key="booking._id" class="booking-card glass-panel">
        <div class="flex justify-content-between align-items-start">
          <Tag :value="booking.status" :severity="getStatusColor(booking.status)" />
          <span class="text-xs text-slate-500">{{ formatDate(booking.bookedAt) }}</span>
        </div>
        
        <h3 class="mt-3 mb-1">{{ booking.activityName }}</h3>
        <p class="text-sm text-slate-400 mb-3"><i class="pi pi-clock mr-1"></i> Verified Booking</p>

        <div class="actions mt-auto pt-3 border-top-1 border-white-alpha-10">
          <Button 
            v-if="booking.status !== 'Completed'"
            label="Mark as Completed" 
            icon="pi pi-check-circle" 
            class="p-button-success w-full"
            @click="confirmCompletion(booking)"
            :loading="isProcessing === booking._id"
          />
          <span v-else class="completed-text">
            <i class="pi pi-verified mr-2"></i>Activity Finished
          </span>
        </div>
      </div>
    </div>

    <Dialog v-model:visible="showCompleteModal" header="Confirm Completion" modal :style="{ width: '350px' }">
      <div class="text-center">
        <i class="pi pi-exclamation-triangle text-4xl text-yellow-500 mb-3"></i>
        <p>Did you finish your <strong>{{ selectedBooking?.activityName }}</strong> session?</p>
        <div class="flex flex-column gap-2 mt-4">
          <Button label="Yes, I did it!" class="p-button-success" @click="handleComplete" />
          <Button label="Not yet" class="p-button-text p-button-secondary" @click="showCompleteModal = false" />
        </div>
      </div>
    </Dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
// --- PRIME VUE IMPORTS ADDED HERE ---
import Tag from 'primevue/tag';
import Button from 'primevue/button';
import Dialog from 'primevue/dialog';

const myBookings = ref([]);
const isLoading = ref(true);
const isProcessing = ref(null);
const showCompleteModal = ref(false);
const selectedBooking = ref(null);

const token = sessionStorage.getItem("auth_token");

// 1. Extract userId from the userData object in SessionStorage
const getUserId = () => {
    const rawData = sessionStorage.getItem("userData");
    if (!rawData) return null;
    try {
        const parsed = JSON.parse(rawData);
        return parsed.userId; 
    } catch (e) {
        return null;
    }
};

// 2. Fetch My Bookings logic
const fetchMyBookings = async () => {
    const currentUserId = getUserId();
    if (!currentUserId) {
        console.error("User ID not found");
        isLoading.value = false;
        return;
    }

    isLoading.value = true;
    try {
        const response = await fetch("http://localhost:8000/r/getMyBookings", {
            method: "POST",
            headers: { 
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}` 
            },
            body: JSON.stringify({ userId: currentUserId }),
        });

        const data = await response.json();
        myBookings.value = Array.isArray(data) ? data : [];
    } catch (err) {
        console.error("Error:", err);
        myBookings.value = [];
    } finally {
        isLoading.value = false;
    }
};

// 3. Mark Activity as Complete
const handleComplete = async () => {
    const currentUserId = getUserId();
    if (!selectedBooking.value || !currentUserId) return;
    
    isProcessing.value = selectedBooking.value._id;
    try {
        const response = await fetch("http://localhost:8000/r/completeActivity", {
            method: "POST",
            headers: { 
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}` 
            },
            body: JSON.stringify({
                sessionId: selectedBooking.value.sessionId,
                userId: currentUserId,
                activityName: selectedBooking.value.activityName
            }),
        });

        if (response.ok) {
            showCompleteModal.value = false;
            await fetchMyBookings(); // Refresh UI
        }
    } catch (err) {
        console.error("Completion error:", err);
    } finally {
        isProcessing.value = null;
    }
};

const confirmCompletion = (booking) => {
    selectedBooking.value = booking;
    showCompleteModal.value = true;
};

const getStatusColor = (status) => {
    return status === 'Completed' ? 'success' : 'warning';
};

const formatDate = (date) => {
    return date ? new Date(date).toLocaleDateString() : 'Scheduled';
};

onMounted(fetchMyBookings);
</script>

<style scoped>
.bookings-container { padding: 2rem; color: white; background: #0f172a; min-height: 100vh; font-family: 'Inter', sans-serif; }
.glass-panel { background: rgba(255,255,255,0.03); backdrop-filter: blur(10px); border: 1px solid rgba(255,255,255,0.1); border-radius: 16px; padding: 1.5rem; }
.bookings-grid { display: grid; grid-template-columns: repeat(auto-fill, minmax(300px, 1fr)); gap: 1.5rem; }
.booking-card { display: flex; flex-direction: column; transition: 0.3s; height: 100%; border: 1px solid rgba(255,255,255,0.1); }
.booking-card:hover { border-color: #22c55e; transform: translateY(-3px); }
.completed-text { color: #22c55e; font-weight: 600; display: flex; align-items: center; justify-content: center; width: 100%; }

/* PrimeVue specific styling for the Dialog dark mode compatibility */
:deep(.p-dialog) {
    background: #1e293b;
    color: white;
    border: 1px solid rgba(255,255,255,0.1);
}
:deep(.p-dialog-header) {
    background: #1e293b;
    color: white;
}
:deep(.p-dialog-content) {
    background: #1e293b;
    color: white;
}
</style>