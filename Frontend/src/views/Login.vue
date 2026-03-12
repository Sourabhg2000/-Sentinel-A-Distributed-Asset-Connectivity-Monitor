<template>
  <div class="login-container">
    <div class="background-blob blob-1"></div>
    <div class="background-blob blob-2"></div>

    <Card class="login-card">
      <template #header>
        <div class="card-header">
          <i class="pi pi-shield user-icon"></i>
          <h2 class="welcome-text">Sentinel</h2>
          <p class="subtitle">High-Availability Monitoring</p>
        </div>
      </template>

      <template #content>
        <div class="login-form">
          <Button
            @click="handleGoogleLogin"
            class="google-btn"
            :loading="googleLoading"
          >
            <i class="pi pi-google mr-2"></i>
            <span>Sign in with Google</span>
          </Button>

          <div class="divider">
            <span>OR CONTINUE WITH EMAIL</span>
          </div>

          <form @submit.prevent="handleLogin" class="inner-form">
            <div class="input-group">
              <label class="label-text">Email</label>
              <InputText
                v-model="username"
                placeholder="admin@sentinel.com"
                class="modern-input"
                required
              />
            </div>

            <div class="input-group">
              <label class="label-text">Password</label>
              <Password
                v-model="password"
                :feedback="false"
                toggleMask
                placeholder="••••••••"
                inputClass="modern-input"
                class="password-wrapper"
                required
              />
            </div>

            <transition name="fade">
              <div v-if="store.loginError" class="error-container">
                {{ store.loginError }}
              </div>
            </transition>

            <Button
              type="submit"
              label="Login to Dashboard"
              class="login-btn"
              :loading="loading"
            />
          </form>
        </div>
      </template>
    </Card>
  </div>
</template>

<style scoped>

.login-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background:
    radial-gradient(circle at 20% 30%, rgba(6,182,212,0.15), transparent 40%),
    radial-gradient(circle at 80% 70%, rgba(59,130,246,0.15), transparent 40%),
    linear-gradient(135deg,#020617,#020617,#0f172a);
  position: relative;
  overflow: hidden;
  padding: 20px;
  font-family: Inter, system-ui, -apple-system, sans-serif;
}

/* PREMIUM BACKGROUND BLOBS */

.background-blob {
  position: absolute;
  width: 550px;
  height: 550px;
  border-radius: 50%;
  filter: blur(120px);
  opacity: 0.35;
  z-index: 0;
  animation: float 12s ease-in-out infinite;
}

.blob-1 {
  background: radial-gradient(circle, #22d3ee, #0891b2);
  top: -120px;
  right: -120px;
}

.blob-2 {
  background: radial-gradient(circle, #3b82f6, #1d4ed8);
  bottom: -120px;
  left: -120px;
}

@keyframes float {
  0% { transform: translateY(0px) }
  50% { transform: translateY(20px) }
  100% { transform: translateY(0px) }
}

/* PREMIUM CARD */

.login-card {
  width: 100%;
  max-width: 420px;
  z-index: 1;

  background: rgba(255,255,255,0.92);
  backdrop-filter: blur(20px);

  border-radius: 24px;
  padding: 10px;

  border: 1px solid rgba(255,255,255,0.25);

  box-shadow:
    0 30px 60px rgba(0,0,0,0.65),
    0 0 0 1px rgba(255,255,255,0.1) inset;
}

/* HEADER */

.card-header {
  text-align: center;
  padding: 20px 10px 10px;
}

.user-icon {
  font-size: 3.2rem;
  color: #06b6d4;
  margin-bottom: 8px;

  filter: drop-shadow(0 0 8px rgba(6,182,212,0.6));
}

.welcome-text {
  font-size: 1.7rem;
  font-weight: 700;
  letter-spacing: 0.5px;
  color: #0f172a;
}

.subtitle {
  font-size: 0.9rem;
  color: #64748b;
  margin-top: 4px;
}

/* GOOGLE BUTTON */

.google-btn {
  width: 100%;
  background: #ffffff !important;
  color: #0f172a !important;
  border: 1px solid #e2e8f0 !important;
  padding: 12px !important;
  border-radius: 12px !important;
  font-weight: 600 !important;

  display: flex;
  justify-content: center;
  align-items: center;
  gap: 10px;

  transition: all 0.25s ease;
}

.google-btn:hover {
  transform: translateY(-2px);
  box-shadow:
    0 8px 20px rgba(0,0,0,0.1);
}

/* DIVIDER */

.divider {
  display: flex;
  align-items: center;
  text-align: center;
  margin: 1.6rem 0;
  color: #94a3b8;
  font-size: 0.7rem;
  letter-spacing: 1px;
}

.divider::before,
.divider::after {
  content: '';
  flex: 1;
  border-bottom: 1px solid #e2e8f0;
}

.divider:not(:empty)::before { margin-right: .6em; }
.divider:not(:empty)::after { margin-left: .6em; }

/* FORM */

.inner-form {
  display: flex;
  flex-direction: column;
  gap: 1.3rem;
}

.label-text {
  font-size: 0.8rem;
  font-weight: 600;
  color: #334155;
}

/* INPUT */

.modern-input,
:deep(.modern-input) {
  width: 100% !important;
  padding: 13px 16px !important;

  border-radius: 12px !important;
  border: 1px solid #cbd5e1 !important;

  transition: all 0.25s ease;
}

.modern-input:focus,
:deep(.modern-input:focus) {
  border-color: #0ea5e9 !important;
  box-shadow: 0 0 0 3px rgba(14,165,233,0.15);
}

/* LOGIN BUTTON */

.login-btn {
  margin-top: 6px;

  background: linear-gradient(135deg,#06b6d4,#2563eb) !important;

  border: none !important;
  border-radius: 12px !important;

  font-weight: 600 !important;

  transition: all 0.25s ease;
}

.login-btn:hover {
  transform: translateY(-2px);

  box-shadow:
    0 10px 25px rgba(37,99,235,0.4);
}

/* ERROR */

.error-container {
  background: rgba(239,68,68,0.1);
  border: 1px solid rgba(239,68,68,0.2);
  color: #b91c1c;
  padding: 10px 12px;
  border-radius: 10px;
  font-size: 0.8rem;
}

/* FADE */

.fade-enter-active,
.fade-leave-active {
  transition: opacity .25s;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>

<script setup>
import { ref, onMounted } from 'vue';
import { useAuthStore } from '../store/login';
import { useRouter } from "vue-router";
import InputText from "primevue/inputtext";
import Password from "primevue/password";
import Button from "primevue/button";
import Card from "primevue/card";

const store = useAuthStore();
const router = useRouter();

const username = ref('');
const password = ref('');
const loading = ref(false);
const googleLoading = ref(false);

const handleLogin = async () => {
    loading.value = true;
    try {
        const success = await store.login(username.value, password.value);
        if (success) {
            router.push("/dashboard"); // Updated to our actual dashboard route
        }
    } finally {
        loading.value = false;
    }
}

// OAuth 2.0 Implementation
const handleGoogleLogin = () => {
    googleLoading.value = true;
    // Industry standard: Redirect to backend which handles the OAuth handshake
    window.location.href = "http://localhost:8080/auth/google";
}

onMounted(() => {
    store.logout();
});
</script>