<template>
  <div class="drink-input-screen">
    <!-- Progress indicator -->
    <div class="progress-indicator">
      <div class="step-dot"></div>
      <div class="step-dot"></div>
      <div class="active-step">
        <div class="step-icon">
          <span class="icon-content">🚭</span>
        </div>
      </div>
      <div class="step-dot"></div>
      <div class="step-dot"></div>
    </div>
    
    <!-- Drink question -->
    <div class="input-form">
      <h1>Do you smoke?</h1>
      
      <div class="options-container">
        <button class="option-button" 
                @click="selectOption('Yes')"
                :class="{ selected: selectedOption === 'Yes' }">Yes</button>
        <button class="option-button" 
                @click="selectOption('Sometimes')"
                :class="{ selected: selectedOption === 'Sometimes' }">Sometimes</button>
        <button class="option-button" 
                @click="selectOption('No')"
                :class="{ selected: selectedOption === 'No' }">No</button>
        <button class="option-button wide-button" 
                @click="selectOption('Prefer Not to Say')"
                :class="{ selected: selectedOption === 'Prefer Not to Say' }">Prefer Not to Say</button>
      </div>
    </div>
    
    <!-- Visibility toggle -->
    <div class="visibility-toggle">
      <span class="toggle-label">Visible</span>
      <label class="toggle-switch">
        <input type="checkbox" v-model="isVisible">
        <span class="slider round"></span>
      </label>
    </div>

    {{ selectedOption }}

    <!-- Next button -->
    <div class="next-button" 
         @click="selectedOption ? goToNextStep() : null"
         :class="{ 'enabled': selectedOption }">
      <span class="next-icon">›</span>
    </div>

    <!-- Bottom bar -->
    <div class="bottom-bar"></div>
  </div>
</template>

<script>
export default {
  name: 'InputSmokeView',
  data() {
    return {
      selectedOption: null,
      isVisible: true
    }
  },
  methods: {
    selectOption(option) {
      this.selectedOption = option;
    },
    goToNextStep() {
      // Navigate to next step
      this.$emit('drink-preference-submitted', { 
        preference: this.selectedOption, 
        visible: this.isVisible 
      });
      // Or use router: this.$router.push('/next-step');
    }
  }
}
</script>

<style scoped>
.drink-input-screen {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  padding: 20px;
  position: relative;
  background-color: white;
  color: black;
}

.progress-indicator {
  display: flex;
  align-items: center;
  justify-content: flex-start;
  margin-bottom: 30px;
  gap: 12px;
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
  background-color: #5d3fd3;
  color: white;
}

.step-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background-color: #ddd;
}

.input-form {
  margin-top: 10px;
  text-align: left;
}

h1 {
  font-size: 32px;
  font-weight: bold;
  margin-bottom: 30px;
  color: black;
  text-align: left;
}

.options-container {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  margin-top: 20px;
}

.option-button {
  background-color: #f0f0f0;
  border: 2px solid transparent;
  border-radius: 25px;
  padding: 12px 24px;
  font-size: 16px;
  cursor: pointer;
  color: black;
  transition: all 0.2s ease;
}

.option-button.selected {
  background-color: #e8e4f7;
  border-color: #5d3fd3;
  color: #5d3fd3;
  font-weight: bold;
}

.wide-button {
  min-width: 170px;
}

.visibility-toggle {
  margin-top: 40px;
  display: flex;
  align-items: center;
  gap: 10px;
}

.toggle-label {
  font-weight: bold;
  color: #5d3fd3;
}

.toggle-switch {
  position: relative;
  display: inline-block;
  width: 60px;
  height: 30px;
}

.toggle-switch input {
  opacity: 0;
  width: 0;
  height: 0;
}

.slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #ccc;
  transition: .4s;
}

.slider.round {
  border-radius: 34px;
}

.slider.round:before {
  border-radius: 50%;
}

.slider:before {
  position: absolute;
  content: "";
  height: 22px;
  width: 22px;
  left: 4px;
  bottom: 4px;
  background-color: white;
  transition: .4s;
}

input:checked + .slider {
  background-color: #5d3fd3;
}

input:checked + .slider:before {
  transform: translateX(30px);
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
  cursor: default;
  transition: all 0.3s ease;
  opacity: 0.5;
}

.next-button.enabled {
  background-color: #5d3fd3;
  opacity: 1;
  cursor: pointer;
}

.next-icon {
  font-size: 30px;
  color: black;
  transition: color 0.3s ease;
}

.next-button.enabled .next-icon {
  color: white;
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
