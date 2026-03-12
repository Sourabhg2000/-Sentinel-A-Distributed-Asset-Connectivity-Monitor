import { defineStore } from 'pinia';
import { ref, computed } from 'vue';

export const useAuthStore = defineStore('auth', () => {
    // State variables
    const user = ref(JSON.parse(sessionStorage.getItem("userData")) || null);
    const token = ref(sessionStorage.getItem("auth_token") || null);
    const roleId = ref(sessionStorage.getItem("user_role") || null);
    const loginError = ref(null);

    // Login function
    const login = async (emailId, password) => {
        const payload = { emailId, password };
        loginError.value = null; // Clear previous errors

        try {
            const response = await fetch("http://localhost:8000/o/login", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(payload),
            });

            const data = await response.json();

            if (response.ok) {
                // ✅ UPDATE PINIA STATE (This makes your UI update!)
                user.value = data.user;
                token.value = data.token;
                roleId.value = data.user.role; // Assuming role is inside user object

                // ✅ UPDATE SESSION STORAGE (For persistence on refresh)
                sessionStorage.setItem("auth_token", data.token);
                sessionStorage.setItem("userData", JSON.stringify(data.user));
                sessionStorage.setItem("user_role", data.user.role);

                console.log("User logged in:", data.user);
                return data.user;
            } else {
                loginError.value = data.error || "Login failed";
                console.error("Login failed:", data.error);
                return null;
            }
        } catch (err) {
            loginError.value = "Network error. Please try again.";
            console.error("Network error:", err);
            return null;
        }
    };

    // Logout function
    const logout = () => {
        console.log("Logging out...");
        // Clear Pinia State
        user.value = null;
        token.value = null;
        roleId.value = null;
        loginError.value = null;

        // Clear Storage
        sessionStorage.clear(); 
    };

    // ✅ Computed property to check if the user is authenticated
    const isAuthenticated = computed(() => !!token.value);

    return { user, token, roleId, loginError, login, logout, isAuthenticated };
}, {
    persist: true 
});