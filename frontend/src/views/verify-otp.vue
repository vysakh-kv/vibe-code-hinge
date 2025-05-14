<script setup>
import { useRouter } from 'vue-router';

const router = useRouter();
</script>

<template>
  <div class="phone-screen">
    <!-- Header with close icon -->
    <div class="header">
      <div class="close-icon">✕</div>
    </div>

    <!-- Phone icon and title -->
    <div class="content">
      <div class="phone-icon" style="align-self: flex-start; margin-bottom: 0.5rem; margin-top: 1.5rem;">
        <div class="icon-circle" style="width: 50px; height: 50px;">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"></path>
          </svg>
        </div>
      </div>

      <h1 class="title" style="color: black; text-align: left; align-self: flex-start; width: 100%; margin-top: 0;">
        Enter the code sent to your number
      </h1>

      <!-- OTP input section -->
      <div class="otp-input-container">
        <input 
          v-for="(digit, index) in 6" 
          :key="index"
          type="text"
          maxlength="1"
          class="otp-digit"
          v-model="otpDigits[index]"
          @input="focusNextInput(index)"
          @keydown.delete="handleBackspace(index, $event)"
          ref="otpInputs"
        />
      </div>

      <!-- Help link -->
      <div class="help-link">
        <span>Didn't get a code?</span>
      </div>
      
      <button 
        class="arrow-circle" 
        :disabled="!isOtpComplete"
        @click="router.push('/input-name')"
        :style="{
          'align-self': 'flex-end', 
          'width': '50px', 
          'height': '50px', 
          'margin-top': '30px',
          'background-color': isOtpComplete ? 'purple' : '#f5f5f5',
          'color': isOtpComplete ? 'white' : '#888',
          'border': 'none',
          'border-radius': '50%',
          'cursor': isOtpComplete ? 'pointer' : 'not-allowed',
          'display': 'flex',
          'align-items': 'center',
          'justify-content': 'center',
          'outline': 'none',
          'opacity': isOtpComplete ? '1' : '0.7'
        }"
      >
        <span class="arrow-right" style="font-size: 2rem;">›</span>
      </button>
    </div>

    <!-- Bottom bar -->
    <div class="bottom-bar">
      <div class="home-indicator"></div>
    </div>

    <!-- Footer -->
    <div class="footer">
      <div class="branding">
        <span class="brand-logo">H</span>
        <span class="brand-name">Hinge</span>
      </div>
      <div class="powered-by">
        <span>curated by</span>
        <span class="mobbin-logo">Mobbin</span>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'VerifyOtp',
  data() {
    return {
      phoneNumber: '',
      otpDigits: Array(6).fill('')
    }
  },
  computed: {
    isPhoneNumberValid() {
      // Ensure phoneNumber is treated as a string and check if it's exactly 10 digits
      return String(this.phoneNumber).length === 10;
    },
    isOtpComplete() {
      // Check if all 6 OTP digits are filled
      return this.otpDigits.every(digit => digit !== '');
    }
  },
  methods: {
    focusNextInput(index) {
      if (index < 5) {
        this.$refs.otpInputs[index + 1].focus();
      }
    },
    handleBackspace(index, event) {
      if (event.key === "Backspace") {
        if (this.otpDigits[index] === '') {
          // If current field is empty, move to previous field and clear it
          if (index > 0) {
            this.otpDigits[index - 1] = '';
            this.$refs.otpInputs[index - 1].focus();
          }
        } else {
          // If current field has content, just clear it and keep focus here
          this.otpDigits[index] = '';
          // Keep focus on current input
          this.$refs.otpInputs[index].focus();
        }
      }
    }
  }
}
</script>

<style scoped>
.phone-screen {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  background-color: white;
  position: relative;
}

.header {
  display: flex;
  justify-content: flex-end;
  padding: 1rem;
}

.close-icon {
  font-size: 1.5rem;
  font-weight: bold;
  cursor: pointer;
}

.content {
  display: flex;
  flex-direction: column;
  align-items: center;
  flex: 1;
  padding: 0 1.5rem;
}

.phone-icon {
  margin-bottom: 1rem;
}

.icon-circle {
  width: 60px;
  height: 60px;
  border-radius: 50%;
  border: 2px solid black;
  display: flex;
  align-items: center;
  justify-content: center;
}

.icon-circle svg {
  width: 30px;
  height: 30px;
  stroke: black;
}

.title {
  font-size: 2.5rem;
  font-weight: bold;
  text-align: center;
  margin-bottom: 3rem;
}

.otp-input-container {
  display: flex;
  justify-content: space-between;
  width: 100%;
  margin-bottom: 1rem;
  gap: 8px;
}

.otp-digit {
  flex: 1;
  max-width: 45px;
  height: 50px;
  text-align: center;
  font-size: 1.25rem;
  border: none;
  border-bottom: 2px solid #e0e0e0;
  outline: none;
  background-color: transparent;
  color: black;
  margin: 0;
  padding: 0.5rem 0;
}

.otp-digit:focus {
  border-bottom: 2px solid purple;
}

.disclaimer {
  color: #888;
  text-align: center;
  font-size: 0.9rem;
  margin-bottom: 2rem;
}

.help-link {
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
  color: purple;
  font-weight: 500;
  padding: 0.5rem 0;
  cursor: pointer;
}

.arrow-circle {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background-color: #f5f5f5;
  display: flex;
  align-items: center;
  justify-content: center;
}

.arrow-right {
  font-size: 1.5rem;
  color: #888;
}

.number-pad {
  margin-top: auto;
  padding: 1rem;
  background-color: #f0f0f0;
}

.number-row {
  display: flex;
  justify-content: space-between;
  margin-bottom: 0.5rem;
}

.number-key, .empty-key, .delete-key {
  width: 32%;
  height: 60px;
  background-color: white;
  border-radius: 8px;
  border: none;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  cursor: pointer;
}

.delete-key svg {
  width: 24px;
  height: 24px;
  stroke: #444;
}

.number {
  font-size: 1.5rem;
  font-weight: medium;
}

.letters {
  font-size: 0.7rem;
  color: #666;
  letter-spacing: 1px;
}

.bottom-bar {
  padding: 0.5rem 0;
  display: flex;
  justify-content: center;
}

.home-indicator {
  width: 30%;
  height: 5px;
  background-color: #ccc;
  border-radius: 3px;
}

.footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.5rem 1rem;
  border-top: 1px solid #eee;
}

.branding {
  display: flex;
  align-items: center;
}

.brand-logo {
  background-color: black;
  color: white;
  width: 24px;
  height: 24px;
  border-radius: 5px;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 0.5rem;
  font-weight: bold;
}

.powered-by {
  display: flex;
  align-items: center;
  color: #666;
  font-size: 0.8rem;
}

.mobbin-logo {
  font-weight: bold;
  margin-left: 0.3rem;
}
</style>
