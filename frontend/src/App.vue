<script setup>
import { ref, computed } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { useAuthStore } from './stores/auth';

const router = useRouter();
const route = useRoute();
const authStore = useAuthStore();

const isAuthenticated = computed(() => authStore.isAuthenticated);

const isPublicPage = computed(() => {
  return ['Home', 'Login', 'Register'].includes(route.name);
});

const logout = () => {
  authStore.logout();
  router.push('/login');
};
</script>

<template>
  <div class="app-container">
    <!-- Nav bar for authenticated users -->
    <header v-if="isAuthenticated && !isPublicPage" class="app-header">
      <nav class="main-nav">
        <div class="nav-brand">Vibe</div>
        <div class="nav-links">
          <router-link to="/discover" class="nav-link">Discover</router-link>
          <router-link to="/matches" class="nav-link">Matches</router-link>
          <router-link to="/profile" class="nav-link">Profile</router-link>
          <router-link to="/settings" class="nav-link">Settings</router-link>
          <button @click="logout" class="nav-link logout-btn">Logout</button>
        </div>
      </nav>
    </header>

    <!-- Main content area -->
    <main class="app-content">
      <router-view></router-view>
    </main>
  </div>
</template>

<style scoped>
.app-container {
  height: 100vh;
  display: flex;
  flex-direction: column;
}

.app-header {
  background-color: #fff;
  border-bottom: 1px solid #e0e0e0;
  padding: 0.5rem 1rem;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
}

.main-nav {
  display: flex;
  justify-content: space-between;
  align-items: center;
  max-width: 1200px;
  margin: 0 auto;
}

.nav-brand {
  font-size: 1.5rem;
  font-weight: bold;
  color: #fd5d8d; /* Hinge-like color */
}

.nav-links {
  display: flex;
  gap: 1.5rem;
}

.nav-link {
  color: #333;
  text-decoration: none;
  font-weight: 500;
  padding: 0.5rem 0;
  position: relative;
}

.nav-link::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  width: 0;
  height: 2px;
  background-color: #fd5d8d;
  transition: width 0.3s;
}

.nav-link:hover::after,
.router-link-active::after {
  width: 100%;
}

.logout-btn {
  background: none;
  border: none;
  cursor: pointer;
  padding: 0.5rem 0;
  font-size: 1rem;
  font-weight: 500;
}

.app-content {
  flex: 1;
  padding: 1rem;
  max-width: 1200px;
  margin: 0 auto;
  width: 100%;
  overflow-y: auto;
}

@media (max-width: 768px) {
  .nav-links {
    gap: 1rem;
  }
}
</style>
