import { defineStore } from 'pinia';
import axios from 'axios';

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api/v1';

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: JSON.parse(localStorage.getItem('user')) || null,
    token: localStorage.getItem('user-token') || null,
    loading: false,
    error: null
  }),
  
  getters: {
    isAuthenticated: (state) => !!state.token,
    currentUser: (state) => state.user
  },
  
  actions: {
    async login(email, password) {
      this.loading = true;
      this.error = null;
      
      try {
        const response = await axios.post(`${API_URL}/auth/login`, {
          email,
          password
        });
        
        const { user, token } = response.data;
        
        this.user = user;
        this.token = token;
        
        localStorage.setItem('user', JSON.stringify(user));
        localStorage.setItem('user-token', token);
        
        return Promise.resolve(user);
      } catch (error) {
        this.error = error.response?.data?.message || 'Login failed';
        return Promise.reject(error);
      } finally {
        this.loading = false;
      }
    },
    
    async register(email, password, firstName, lastName) {
      this.loading = true;
      this.error = null;
      
      try {
        const response = await axios.post(`${API_URL}/auth/register`, {
          email,
          password,
          firstName,
          lastName
        });
        
        const { user, token } = response.data;
        
        this.user = user;
        this.token = token;
        
        localStorage.setItem('user', JSON.stringify(user));
        localStorage.setItem('user-token', token);
        
        return Promise.resolve(user);
      } catch (error) {
        this.error = error.response?.data?.message || 'Registration failed';
        return Promise.reject(error);
      } finally {
        this.loading = false;
      }
    },
    
    logout() {
      this.user = null;
      this.token = null;
      localStorage.removeItem('user');
      localStorage.removeItem('user-token');
    }
  }
}); 