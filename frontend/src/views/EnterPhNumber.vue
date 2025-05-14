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

      <h1 class="title" style="color: black; text-align: left; align-self: flex-start; width: 100%; margin-top: 0;">What's your phone number?</h1>

      <!-- Phone input section -->
      <div class="phone-input-container">
        <div class="country-code">
          <div class="flag">🇮🇳</div>
          <div class="code" style="color: black;">+91</div>
          <div class="arrow">▼</div>
        </div>
        <div class="divider"></div>
        <input class="phone-input" type="number" placeholder="" v-model="phoneNumber" style="background-color: transparent; color: black; -webkit-appearance: none; -moz-appearance: textfield; appearance: textfield; margin: 0;" />
      </div>

      <p class="disclaimer">Hinge will send you a text with a verification code. Message and data rates may apply.</p>

      <!-- Help link -->
      <div class="help-link">
        <span>What if my number changes?</span>
      </div>
      
      <button 
        class="arrow-circle" 
        :disabled="!isPhoneNumberValid"
        :style="{
          'align-self': 'flex-end', 
          'width': '50px', 
          'height': '50px', 
          'margin-top': '30px',
          'background-color': isPhoneNumberValid ? 'purple' : '#f5f5f5',
          'color': isPhoneNumberValid ? 'white' : '#888',
          'border': 'none',
          'border-radius': '50%',
          'cursor': isPhoneNumberValid ? 'pointer' : 'not-allowed',
          'display': 'flex',
          'align-items': 'center',
          'justify-content': 'center',
          'outline': 'none',
          'opacity': isPhoneNumberValid ? '1' : '0.7'
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
  name: 'EnterPhNumber',
  data() {
    return {
      phoneNumber: ''
    }
  },
  computed: {
    isPhoneNumberValid() {
      // Ensure phoneNumber is treated as a string and check if it's exactly 10 digits
      return String(this.phoneNumber).length === 10;
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

.phone-input-container {
  display: flex;
  width: 100%;
  border-bottom: 1px solid #e0e0e0;
  margin-bottom: 1rem;
}

.country-code {
  display: flex;
  align-items: center;
  padding: 0.75rem 0;
}

.flag {
  margin-right: 0.5rem;
}

.code {
  font-size: 1.25rem;
  font-weight: bold;
  margin-right: 0.25rem;
}

.arrow {
  font-size: 0.75rem;
  color: #888;
}

.divider {
  width: 1px;
  background-color: #e0e0e0;
  margin: 0 1rem;
}

.phone-input {
  flex: 1;
  border: none;
  font-size: 1.25rem;
  padding: 0.75rem 0;
  outline: none;
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
