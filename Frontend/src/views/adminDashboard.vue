<template>
  <div class="admin-dashboard">
    <h1 class="page-title">Admin Command Center</h1>

    <div class="stats-grid">
      <div class="stat-card">
        <i class="pi pi-users"></i>
        <div>
          <p class="stat-label">Total Users</p>
          <p class="stat-value">1,284</p>
        </div>
      </div>
      <div class="stat-card">
        <i class="pi pi-bolt"></i>
        <div>
          <p class="stat-label">Daily Activities</p>
          <p class="stat-value">450</p>
        </div>
      </div>
      <div class="stat-card warning">
        <i class="pi pi-exclamation-triangle"></i>
        <div>
          <p class="stat-label">Pending Moderation</p>
          <p class="stat-value">12</p>
        </div>
      </div>
    </div>

    <div class="admin-content-grid">
<section class="glass-panel">
  <div class="panel-header">
    <h3><i class="pi pi-cog"></i> Manage Activity Types</h3>
<Button 
  label="Add Type" 
  icon="pi pi-plus" 
  class="p-button-sm p-button-success" 
  @click="showActivityModal = true" 
/>

<Dialog 
  v-model:visible="showActivityModal" 
  header="Add New Activity Type" 
  :modal="true" 
  class="glass-modal"
  :style="{ width: '400px' }"
>
  <form @submit.prevent="handleAddActivity" class="modal-form">
    <div class="flex flex-column gap-3">
      <div class="field flex flex-column gap-2">
        <label>Activity Name</label>
        <InputText v-model="newActivity.activityName" placeholder="e.g. Boxing" required />
      </div>
      
      <div class="field flex flex-column gap-2">
        <label>Calories Per Hour</label>
        <InputText v-model="newActivity.caloriesPerHr" placeholder="600" :useGrouping="false" required />
      </div>

      <div class="field flex flex-column gap-2">
        <label>Category</label>
        <Dropdown 
          v-model="newActivity.activityType" 
          :options="['Cardio', 'Strength', 'Flexibility', 'Sports']" 
          placeholder="Select Category" 
          required
        />
      </div>
    </div>

    <div class="form-actions mt-4 flex justify-content-end gap-2">
      <Button label="Cancel" class="p-button-text p-button-secondary" @click="showActivityModal = false" />
      <Button type="submit" label="Save Activity" class="p-button-success" :loading="isSavingActivity" />
    </div>
  </form>
</Dialog>  </div>

  <DataTable 
    :value="activityTypes" 
    :paginator="true" 
    :rows="4" 
    paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink"
    responsiveLayout="scroll"
    class="p-datatable-sm modern-table"
  >
    <Column field="activityName" header="Activity Name"></Column>
    <Column field="caloriesPerHr" header="Avg Cal/Hr"></Column>
    <Column field="activityType" header="Type"></Column>
    
  </DataTable>
</section>
    </div>

<section class="glass-panel full-width">
  <div class="panel-header">
    <h3><i class="pi pi-users"></i> User Management</h3>
    <Button 
      label="Register New User" 
      icon="pi pi-user-plus" 
      class="p-button-success" 
      @click="showRegisterModal = true" 
    />
  </div>

  <Dialog 
    v-model:visible="showRegisterModal" 
    header="Create New User Account" 
    :modal="true" 
    class="glass-modal"
    :style="{ width: '50vw' }"
  >
    <form @submit.prevent="handleRegister" class="modal-form">
      <div class="form-grid">
        <div class="field">
          <label>First Name</label>
          <InputText v-model="newUser.firstName" placeholder="e.g. Sourabh" />
        </div>
        <div class="field">
          <label>Last Name</label>
          <InputText v-model="newUser.lastName" placeholder="e.g. Gaikwad" />
        </div>
        <div class="field">
          <label>Full Name</label>
          <InputText v-model="fullNameComputed" readonly placeholder="Auto-generated" />
        </div>
        <div class="field">
  <label>Gender</label>
  <Dropdown 
    v-model="newUser.gender" 
    :options="['Male', 'Female', 'Other']" 
    placeholder="Select Gender" 
  />
</div>
        <div class="field">
          <label>Email ID</label>
          <InputText v-model="newUser.emailId" placeholder="user@mkcl.org" />
        </div>
        <div class="field">
          <label>MobileNo</label>
          <InputText v-model="newUser.phone" placeholder="e.g. 9876543210" />
        </div>
        <div class="field">
          <label>Date of Birth</label>
          <Calendar v-model="newUser.dob" dateFormat="dd-mm-yy" placeholder="Select Date" />
        </div>
        <div class="field">
          <label>Role</label>
          <Dropdown v-model="newUser.role" :options="['User']" placeholder="Select Role" />
        </div>
      </div>
      
      <div class="form-actions">
        <Button label="Cancel" class="p-button-text p-button-secondary" @click="showRegisterModal = false" />
        <Button type="submit" label="Register User" class="p-button-success" :loading="isRegistering" />
      </div>
    </form>
  </Dialog>
</section>
  </div>
</template>

<script setup>
import { onMounted, ref ,computed} from 'vue';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Button from 'primevue/button';
import InputText from 'primevue/inputtext';
import Dropdown from 'primevue/dropdown';
import Dialog from 'primevue/dialog';
import Calendar from 'primevue/calendar';
const showRegisterModal = ref(false);
const showActivityModal = ref(false);
const isSavingActivity = ref(false);

const newActivity = ref({
    activityName: '',
    caloriesPerHr: null,
    activityType: '',
    icon: 'pi-tag',
    isActive: true
});
const isRegistering = ref(false);

const newUser = ref({
    firstName: '',
    lastName: '',
    emailId: '',
    phone: '',
    gender: '',
    dob: null,
    role: 'User'
});

const fullNameComputed = computed(() => {
    return `${newUser.value.firstName} ${newUser.value.lastName}`.trim();
});

const activityTypes = ref([
]);


const handleAddActivity = async () => {
    isSavingActivity.value = true;
    const token = sessionStorage.getItem("auth_token");
    const payload = {
        activityName: newActivity.value.activityName,
      caloriesPerHr: Number(newActivity.value.caloriesPerHr),
        activityType: newActivity.value.activityType,
        icon: newActivity.value.icon,
        isActive: newActivity.value.isActive
    };

    try {
        // NOTE: You'll need to create this endpoint in Go!
        const response = await fetch("http://localhost:8000/r/addActivityType", {
            method: "POST",
            headers: { 
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}` 
            },
            body: JSON.stringify(payload),
        });

        if (response.ok) {
            showActivityModal.value = false;
            // Reset form
            newActivity.value = { activityName: '', caloriesPerHr: null, activityType: '', icon: 'pi-tag', isActive: true };
            // Refresh the table list
            await fetchActivityTypes(); 
        } else {
            const err = await response.json();
            console.error("Save failed:", err.error);
        }
    } catch (error) {
        console.error("Network error:", error);
    } finally {
        isSavingActivity.value = false;
    }
};

const handleRegister = async () => {
    isRegistering.value = true;
   const token = sessionStorage.getItem("auth_token");
    const payload = {
        firstName: newUser.value.firstName,
        lastName: newUser.value.lastName,
        emailId: newUser.value.emailId,
        phone: newUser.value.phone,
        gender: newUser.value.gender,
        dob: newUser.value.dob ? newUser.value.dob.toISOString().split('T')[0] : null, // Format as YYYY-MM-DD
        role: newUser.value.role
    };

        try {
        const response = await fetch("http://localhost:8000/r/register", {
            method: "POST",
            headers: { 
                "Content-Type": "application/json",
                // 3. Add the Authorization Header
                "Authorization": `Bearer ${token}` 
            },
            body: JSON.stringify(payload),
        });

        const data = await response.json();

        if (response.ok) {
   newUser.value = { firstName: '', lastName: '', emailId: '', phone: '', dob: null, role: 'User' };
        } else {
            // Handle unauthorized (401) specifically if token is expired
            if (response.status === 401) {
                console.error("Session expired. Please login again.",data.error);
            }
            return null;
        }
    } catch (err) {
        console.error("Network error:", err);
        return null;
    }finally {
        isRegistering.value = false; 
        showRegisterModal.value = false;
    }
};  

onMounted(() => {
  // Fetch activity types and recent entries from API if needed
  fetchActivityTypes();
  console.log("Admin Dashboard Mounted - API calls can be made here activity Types and recent entries",activityTypes.value);
});

const fetchActivityTypes = async () => {
    // 1. Grab the token from Pinia state or SessionStorage
    const token = sessionStorage.getItem("auth_token");
    
    // 2. Define empty payload
    const payload = {}; 

    try {
        const response = await fetch("http://localhost:8000/r/activityTypes", {
            method: "POST",
            headers: { 
                "Content-Type": "application/json",
                // 3. Add the Authorization Header
                "Authorization": `Bearer ${token}` 
            },
            body: JSON.stringify(payload),
        });

        const data = await response.json();

        if (response.ok) {
          console.log("Activity Types:", data);
          activityTypes.value = data;
            return data;
        } else {
            // Handle unauthorized (401) specifically if token is expired
            if (response.status === 401) {
                console.error("Session expired. Please login again.");
            }
            return null;
        }
    } catch (err) {
        console.error("Network error:", err);
        return null;
    }
};
</script>

<style scoped>
/* Glass Effect for the Dialog */
:deep(.glass-modal.p-dialog) {
    background: rgba(15, 23, 42, 0.9);
    backdrop-filter: blur(20px);
    border: 1px solid rgba(255, 255, 255, 0.1);
    color: white;
}

:deep(.glass-modal .p-dialog-header),
:deep(.glass-modal .p-dialog-content) {
    background: transparent;
    color: white;
}

.modal-form {
    padding-top: 1rem;
}

.form-grid {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 1.5rem;
}

.field {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
}

.field label {
    font-size: 0.85rem;
    color: #94a3b8;
}

.form-actions {
    margin-top: 2rem;
    display: flex;
    justify-content: flex-end;
    gap: 1rem;
}

/* Ensure inputs in modal fit the theme */
:deep(.p-inputtext), :deep(.p-dropdown), :deep(.p-calendar) {
    background: rgba(255, 255, 255, 0.05);
    border: 1px solid rgba(255, 255, 255, 0.1);
    color: white;
}
.admin-dashboard {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.page-title {
  font-size: 2rem;
  font-weight: 800;
  color: #fff;
}

/* Stats Cards */
.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  gap: 1.5rem;
}

.stat-card {
  background: rgba(255, 255, 255, 0.05);
  backdrop-filter: blur(10px);
  border: 1px solid rgba(255, 255, 255, 0.1);
  padding: 1.5rem;
  border-radius: 20px;
  display: flex;
  align-items: center;
  gap: 20px;
}

.stat-card i {
  font-size: 2.5rem;
  color: #22c55e;
}

.stat-card.warning i { color: #f59e0b; }

.stat-label { color: #94a3b8; font-size: 0.9rem; margin: 0; }
.stat-value { font-size: 1.8rem; font-weight: 800; margin: 0; }

/* Content Panels */
.admin-content-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(400px, 1fr));
  gap: 1.5rem;
}

.glass-panel {
  background: rgba(30, 41, 59, 0.6);
  backdrop-filter: blur(15px);
  border: 1px solid rgba(255, 255, 255, 0.05);
  padding: 1.5rem;
  border-radius: 20px;
}

.full-width { grid-column: 1 / -1; }

.panel-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
}

/* Forms & Tables */
.admin-reg-form {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1rem;
}

.status-tag {
  padding: 4px 8px;
  border-radius: 6px;
  font-size: 0.75rem;
  font-weight: 700;
}

.status-tag.pending { background: rgba(245, 158, 11, 0.2); color: #f59e0b; }
.status-tag.flagged { background: rgba(239, 68, 68, 0.2); color: #ef4444; }

:deep(.modern-table .p-datatable-thead > tr > th) {
  background: transparent;
  color: #94a3b8;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

:deep(.modern-table .p-datatable-tbody > tr) {
  background: transparent;
  color: #f8fafc;
}
</style>