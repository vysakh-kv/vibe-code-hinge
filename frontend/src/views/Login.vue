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

const login = async () => {
  if (!email.value || !password.value) {
    error.value = 'Please enter both email and password';
    return;
  }
  
  loading.value = true;
  error.value = '';
  
  try {
    await authStore.login(email.value, password.value);
    router.push('/discover');
  } catch (err) {
    error.value = err.response?.data?.message || 'Login failed. Please check your credentials.';
  } finally {
    loading.value = false;
  }
};

const navigateToRegister = () => {
  router.push('/register');
};

const navigateToHome = () => {
  router.push('/');
};
</script>

<template>
  <div class="login-container">
    <div class="login-card">
      <div class="login-header">
        <h1 class="app-logo" @click="navigateToHome">Vibe</h1>
        <h2 class="login-title">Log In</h2>
        <p class="login-subtitle">Welcome back! Please log in to your account.</p>
      </div>
      
      <form @submit.prevent="login" class="login-form">
        <div v-if="error" class="error-message">
          {{ error }}
        </div>
        
        <div class="form-group">
          <label for="email">Email</label>
          <input 
            id="email" 
            type="email" 
            v-model="email" 
            placeholder="Enter your email"
            autocomplete="email"
            required
          />
        </div>
        
        <div class="form-group">
          <label for="password">Password</label>
          <input 
            id="password" 
            type="password" 
            v-model="password"
            placeholder="Enter your password"
            autocomplete="current-password"
            required
          />
        </div>
        
        <button 
          type="submit" 
          class="login-button" 
          :disabled="loading"
        >
          {{ loading ? 'Logging in...' : 'Log In' }}
        </button>
      </form>
      
      <div class="login-footer">
        <p>
          Don't have an account? 
          <a href="#" @click.prevent="navigateToRegister">Sign up</a>
        </p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 80vh;
  padding: 2rem;
}

.login-card {
  width: 100%;
  max-width: 450px;
  background-color: white;
  border-radius: 12px;
  box-shadow: 0 6px 24px rgba(0, 0, 0, 0.1);
  padding: 2.5rem;
}

.login-header {
  text-align: center;
  margin-bottom: 2rem;
}

.app-logo {
  color: #fd5d8d;
  margin-bottom: 1.5rem;
  cursor: pointer;
}

.login-title {
  font-size: 1.75rem;
  margin-bottom: 0.5rem;
  color: #333;
}

.login-subtitle {
  color: #666;
}

.login-form {
  margin-bottom: 2rem;
}

.form-group {
  margin-bottom: 1.5rem;
}

.form-group label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 500;
  color: #444;
}

.form-group input {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 1rem;
  transition: border-color 0.2s;
}

.form-group input:focus {
  border-color: #fd5d8d;
  outline: none;
}

.login-button {
  width: 100%;
  padding: 0.875rem;
  background-color: #fd5d8d;
  color: white;
  border: none;
  border-radius: 6px;
  font-size: 1rem;
  font-weight: 500;
  cursor: pointer;
  transition: background-color 0.2s;
}

.login-button:hover {
  background-color: #e54b7b;
}

.login-button:disabled {
  background-color: #fda0ba;
  cursor: not-allowed;
}

.login-footer {
  text-align: center;
  color: #666;
}

.login-footer a {
  color: #fd5d8d;
  text-decoration: none;
  font-weight: 500;
}

.login-footer a:hover {
  text-decoration: underline;
}

.error-message {
  background-color: #ffebee;
  color: #d32f2f;
  padding: 0.75rem;
  border-radius: 6px;
  margin-bottom: 1.5rem;
  font-size: 0.875rem;
}
</style> 