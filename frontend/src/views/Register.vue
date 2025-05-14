<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '../stores/auth';

const router = useRouter();
const authStore = useAuthStore();

const email = ref('');
const password = ref('');
const error = ref('');
const loading = ref(false);

const register = async () => {
  if (!email.value || !password.value) {
    error.value = 'Please fill out all required fields';
    return;
  }
  
  loading.value = true;
  error.value = '';
  
  try {
    await authStore.register(email.value, password.value);
    router.push('/onboarding');
  } catch (err) {
    error.value = err.response?.data?.message || 'Registration failed. Please try again.';
  } finally {
    loading.value = false;
  }
};

const navigateToLogin = () => {
  router.push('/login');
};
</script>

<template>
  <div class="register-page">
    <div class="background-image"></div>
    <div class="register-container">
      <div class="register-header">
        <div class="app-logo">
          <img src="https://upload.wikimedia.org/wikipedia/commons/8/87/Hinge_logo.svg" alt="Hinge" style="filter: brightness(0) saturate(100%) invert(100%) sepia(7%) saturate(352%) hue-rotate(214deg) brightness(114%) contrast(90%); -webkit-filter: brightness(0) saturate(100%) invert(100%) sepia(7%) saturate(352%) hue-rotate(214deg) brightness(114%) contrast(90%);" />
        </div>
        <p class="app-tagline">Designed to be deleted.</p>
      </div>
      
      <div class="terms-container">
        <p class="terms-text">
          By tapping 'Sign in' / 'Create account', you agree to our
          <a href="#" class="terms-link">Terms of Service</a>. Learn how we process your data in
          our <a href="#" class="terms-link">Privacy Policy</a> and <a href="#" class="terms-link">Cookies Policy</a>.
        </p>
      </div>
      
      <button 
        class="create-account-button" 
        @click="router.push('/register-form')"
        :disabled="loading"
      >
        Create account
      </button>
      
      <div class="sign-in-link">
        <button type="button" class="sign-in-button" @click="navigateToLogin">
          Sign in
        </button>
      </div>
      
      <div class="footer">
        <div class="app-icon">
          <img src="https://logos-download.com/wp-content/uploads/2019/11/HInge_Logo_full.png" alt="Hinge" class="footer-logo" />
        </div>
        <div class="curated-by">
          <span>curated by</span>
          <span class="mobbin">Mobbin</span>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.register-page {
  min-height: 100vh;
  color: white;
  display: flex;
  flex-direction: column;
  position: relative;
  overflow: hidden;
  max-width: 100vw;
}

.background-image {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: #000;
  background-image: url('https://images.unsplash.com/photo-1575390130709-7b5fee2919e4?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D');
  background-size: cover;
  background-position: center;
  opacity: 0.7;
  z-index: -1;
  animation: zoomInEffect 8s ease forwards;
  transform-origin: center;
}

@keyframes zoomInEffect {
  0% {
    transform: scale(1);
  }
  100% {
    transform: scale(1.15);
  }
}

.register-container {
  flex: 1;
  display: flex;
  flex-direction: column;
  padding: 2rem 1.5rem;
  width: 100%;
  box-sizing: border-box;
  position: relative;
  z-index: 1;
  height: 100vh;
}

.register-header {
  margin-top: 30%;
  text-align: center;
}

.app-logo {
  margin-bottom: 0.25rem;
  text-align: center;
}

.app-logo img {
  width: 140px;
  height: auto;
  max-width: 100%;
}

.app-tagline {
  font-size: 1.25rem;
  opacity: 0.9;
  margin-top: 0;
}

.terms-container {
  margin-top: auto;
  margin-bottom: 1rem;
}

.terms-text {
  text-align: center;
  font-size: 0.75rem;
  line-height: 1.4;
  color: rgba(255, 255, 255, 0.8);
  padding: 0 0.5rem;
}

.terms-link {
  color: white;
  text-decoration: underline;
  font-weight: 500;
}

.create-account-button {
  background-color: #9c2988;
  color: white;
  border: none;
  padding: 0.9rem;
  border-radius: 2rem;
  font-size: 1.1rem;
  font-weight: 500;
  cursor: pointer;
  margin-bottom: 0.75rem;
  width: 100%;
  text-align: center;
}

.create-account-button:hover {
  background-color: #8c2379;
}

.create-account-button:disabled {
  background-color: #63175a;
  cursor: not-allowed;
}

.sign-in-link {
  text-align: center;
  margin-bottom: 3rem;
}

.sign-in-button {
  background: transparent;
  border: none;
  color: white;
  font-size: 1rem;
  cursor: pointer;
  padding: 0.5rem 1rem;
}

.sign-in-button:hover {
  text-decoration: underline;
}

.error-message {
  background-color: rgba(255, 59, 48, 0.15);
  color: #ff3b30;
  padding: 0.75rem;
  border-radius: 0.5rem;
  font-size: 0.875rem;
}

.footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  position: absolute;
  bottom: 2rem;
  left: 0;
  right: 0;
  font-size: 0.75rem;
  padding: 0 1.5rem;
  box-sizing: border-box;
}

.app-icon {
  display: flex;
  align-items: center;
}

.footer-logo {
  width: 60px;
  height: auto;
}

.curated-by {
  display: flex;
  align-items: center;
  gap: 0.25rem;
  opacity: 0.8;
}

.mobbin {
  font-weight: 600;
}
</style>
