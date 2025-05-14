<script setup>
import { useRouter } from 'vue-router';

const router = useRouter();
</script>

<template>
  <div class="name-input-screen">
    <!-- Header notice -->
    <div class="header-notice">
      <p>NO BACKGROUND CHECKS ARE CONDUCTED</p>
    </div>
    
    <!-- Progress indicator -->
    <div class="progress-indicator">
      <div class="active-step">
        <div class="step-icon">
          <span class="icon-content">≡</span>
        </div>
      </div>
      <div class="step-dot"></div>
      <div class="step-dot"></div>
      <div class="step-dot"></div>
    </div>
    
    <!-- Name input form -->
    <div class="input-form">
      <h1>What's your name?</h1>
      
      <div class="input-group">
        <input 
          type="text" 
          v-model="firstName" 
          placeholder="First name (required)" 
          required
          class="text-input"
        />
        <div class="input-line"></div>
      </div>
      
      <div class="input-group">
        <input 
          type="text" 
          v-model="lastName" 
          placeholder="Last name" 
          class="text-input"
        />
        <div class="input-line"></div>
        <div class="last-name-info">
          Last name is optional, and only shared with matches. 
          <span class="why-link" @click="showWhyLastNameOptional">Why?</span>
        </div>
      </div>
    </div>
    
    <!-- Next button -->
    <div class="next-button" @click="router.push('/input-email')" :class="{ 'button-disabled': firstName.length <= 3 }">
      <span class="next-icon">›</span>
    </div>

    <!-- Bottom bar -->
    <div class="bottom-bar"></div>
  </div>
</template>

<script>
export default {
  name: 'InputNameView',
  data() {
    return {
      firstName: '',
      lastName: '',
    }
  },
  methods: {
    goToNextStep() {
      router.push('/input-email');
      if (this.firstName.trim().length > 3) {
        // Navigate to next step
        this.$emit('name-submitted', { firstName: this.firstName, lastName: this.lastName });
        // Or use router: this.$router.push('/next-step');
      }
    },
    showWhyLastNameOptional() {
      // Show modal or tooltip explaining why last name is optional
      alert('Last names are optional to protect your privacy until you match with someone.');
    }
  }
}
</script>

<style scoped>
.name-input-screen {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  padding: 20px;
  position: relative;
  background-color: white;
  color: black;
}

.header-notice {
  text-align: center;
  margin-top: 20px;
  margin-bottom: 40px;
  font-weight: bold;
}

.progress-indicator {
  display: flex;
  align-items: center;
  justify-content: flex-start;
  margin-bottom: 30px;
  gap: 5px;
}

.active-step {
  display: flex;
  justify-content: center;
  align-items: center;
}

.step-icon {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  border: 2px solid #5d3fd3;
  display: flex;
  justify-content: center;
  align-items: center;
  color: #5d3fd3;
  font-size: 18px;
}

.step-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background-color: #ddd;
  margin: 0 5px;
}

.input-form {
  margin-top: 10px;
  text-align: left;
}

h1 {
  font-size: 26px;
  font-weight: bold;
  margin-bottom: 30px;
  color: black;
  text-align: left;
}

.input-group {
  margin-bottom: 25px;
}

.text-input {
  width: 100%;
  padding: 10px 0;
  border: none;
  background: transparent;
  font-size: 18px;
  color: black;
}

.text-input::placeholder {
  color: #ccc;
  font-style: italic;
}

.input-line {
  height: 1px;
  width: 100%;
  background-color: #ddd;
}

.last-name-info {
  margin-top: 8px;
  font-size: 14px;
  color: black;
}

.why-link {
  color: #5d3fd3;
  font-weight: bold;
  cursor: pointer;
}

.next-button {
  position: absolute;
  bottom: 80px;
  right: 20px;
  width: 50px;
  height: 50px;
  border-radius: 50%;
  background-color: #f0f0f0;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
  transition: opacity 0.3s, background-color 0.3s;
}

.next-button:not(.button-disabled) {
  background-color: #5d3fd3;
}

.next-button:not(.button-disabled) .next-icon {
  color: white;
}

.button-disabled {
  opacity: 0.5;
  cursor: not-allowed;
  pointer-events: none;
  background-color: #f0f0f0;
}

.next-icon {
  font-size: 30px;
  color: black;
  transition: color 0.3s;
}

.bottom-bar {
  position: absolute;
  bottom: 20px;
  left: 50%;
  transform: translateX(-50%);
  height: 5px;
  width: 40%;
  background-color: #ddd;
  border-radius: 3px;
}
</style>
