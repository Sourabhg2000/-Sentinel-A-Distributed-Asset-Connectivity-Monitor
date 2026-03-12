<template>
  <div class="leaderboard-wrapper glass-panel">
    <div v-if="isLoading" class="text-center p-5">
       <i class="pi pi-spin pi-spinner text-4xl text-green-500"></i>
       <p class="mt-2 text-slate-400">Calculating rankings...</p>
    </div>

    <template v-else>
      <div class="leaderboard-header">
        <div class="title-group">
          <h2 class="gradient-text">Fitness Elite</h2>
          <p>Ranked by Sessions Completed</p>
        </div>
        <div class="crown-icon" @click="fetchLeaderboard" style="cursor: pointer">
          <i class="pi pi-sync" :class="{'pi-spin': isLoading}" style="font-size: 1.5rem; color: #64748b"></i>
        </div>
      </div>

      <div class="podium-container" v-if="leaderboardData.length > 0">
        <div class="podium-spot rank-2">
          <div class="avatar-wrapper">
            <img :src="`https://api.dicebear.com/7.x/avataaars/svg?seed=${leaderboardData[1]?.firstName}`" alt="user" />
            <div class="badge">2</div>
          </div>
          <span class="podium-name">{{ leaderboardData[1]?.firstName || '---' }}</span>
          <span class="podium-score">{{ leaderboardData[1]?.bookingCount || 0 }} <small>sessions</small></span>
        </div>

        <div class="podium-spot rank-1">
          <div class="avatar-wrapper">
            <img :src="`https://api.dicebear.com/7.x/avataaars/svg?seed=${leaderboardData[0]?.firstName}`" alt="user" />
            <div class="badge gold-crown"><i class="pi pi-star-fill"></i></div>
          </div>
          <span class="podium-name">{{ leaderboardData[0]?.firstName || '---' }}</span>
          <span class="podium-score">{{ leaderboardData[0]?.bookingCount || 0 }} <small>sessions</small></span>
        </div>

        <div class="podium-spot rank-3">
          <div class="avatar-wrapper">
            <img :src="`https://api.dicebear.com/7.x/avataaars/svg?seed=${leaderboardData[2]?.firstName}`" alt="user" />
            <div class="badge">3</div>
          </div>
          <span class="podium-name">{{ leaderboardData[2]?.firstName || '---' }}</span>
          <span class="podium-score">{{ leaderboardData[2]?.bookingCount || 0 }} <small>sessions</small></span>
        </div>
      </div>

      <div class="rank-list-container">
        <div v-for="(user, index) in leaderboardData.slice(3)" :key="user.userId" 
             :class="['list-item', { 'is-me': user.userId === currentUserId }]">
          <span class="rank-idx">#{{ index + 4 }}</span>
          <Avatar :image="`https://api.dicebear.com/7.x/initials/svg?seed=${user.firstName}`" shape="circle" size="medium" />
          <div class="user-meta">
            <span class="name">{{ user.firstName }} {{ user.lastName }}</span>
            <span class="role">{{ user.bookingCount }} Activities</span>
          </div>
          <div class="score-pill">
            {{ user.bookingCount }} <span>runs</span>
          </div>
        </div>
      </div>

      <div class="leaderboard-footer">
        <div class="my-rank-info">
          <span class="text-slate-400">Your Current Rank</span>
          <span class="highlight-text">#{{ currentUserRank }}</span>
        </div>
        <Button icon="pi pi-share-alt" label="Share" class="p-button-rounded p-button-success p-button-text" />
      </div>
    </template>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { apiCallpost } from '../plugins/testAxios';
import Avatar from 'primevue/avatar';
import Button from 'primevue/button';

const leaderboardData = ref([]);
const isLoading = ref(true);
const currentUserId = ref("");
const currentUserRank = ref("N/A");

// Get userId safely from sessionStorage
const getUserIdFromStorage = () => {
    const rawData = sessionStorage.getItem("userData");
    if (!rawData) return null;
    try {
        const parsed = JSON.parse(rawData);
        return parsed.userId;
    } catch (e) {
        return null;
    }
};

const fetchLeaderboard = async () => {
    isLoading.value = true;
    const uid = getUserIdFromStorage();
    currentUserId.value = uid;

    const token = sessionStorage.getItem("auth_token");

    try {
        const response = await fetch("http://localhost:8000/r/leaderboard", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            },
            // Sending empty object as requested, or you can send { userId: uid }
            body: JSON.stringify({}) 
        });

        if (!response.ok) throw new Error("Failed to fetch rankings");

        const data = await response.json();

        if (data && Array.isArray(data)) {
            // Data is already sorted by bookingCount from the Go backend
            leaderboardData.value = data;
            
            // Local rank calculation for the footer
            const myIndex = leaderboardData.value.findIndex(user => user.userId === uid);
            currentUserRank.value = myIndex !== -1 ? myIndex + 1 : "N/A";
        }
    } catch (error) {
        console.error("Leaderboard UI Error:", error);
    } finally {
        isLoading.value = false;
    }
};

onMounted(() => {
    fetchLeaderboard();
});
</script>

<style scoped>
/* Keep your existing CSS - it works great with this structure */
/* Added a small scrollbar tweak for the list */
.rank-list-container::-webkit-scrollbar {
  width: 4px;
}
.rank-list-container::-webkit-scrollbar-thumb {
  background: rgba(34, 197, 94, 0.3);
  border-radius: 10px;
}
</style>