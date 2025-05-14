<template>
  <div class="name-input-screen">
    <!-- Header notice -->
    <div class="header-notice">
      <p>NO BACKGROUND CHECKS ARE CONDUCTED</p>
    </div>
    
    <!-- Progress indicator -->
    <div class="progress-indicator">
      <div class="step-dot"></div>
      <div class="active-step">
        <div class="step-icon">
          <span class="icon-content">≡</span>
        </div>
      </div>
      <div class="step-dot"></div>
      <div class="step-dot"></div>
    </div>
    
    <!-- Prompts section -->
    <div class="input-form">
      <h1>Add some prompts</h1>
      <p class="instruction-text">Select and answer 3 prompts to show off your personality</p>
      
      <div v-for="(prompt, index) in selectedPrompts" :key="index" class="prompt-container">
        <div class="prompt-header">
          <div class="prompt-selector" @click="openPromptSelector(index)">
            <span>{{ prompt.question || 'Select a prompt' }}</span>
            <span class="dropdown-icon">▼</span>
          </div>
        </div>
        
        <div class="input-group">
          <textarea 
            v-model="prompt.answer" 
            placeholder="Your answer here..." 
            rows="3"
            class="text-input prompt-answer"
            :disabled="!prompt.question"
          ></textarea>
          <div class="input-line"></div>
          <div class="character-count" :class="{ 'limit-reached': prompt.answer.length > 150 }">
            {{ prompt.answer.length }}/150
          </div>
        </div>
      </div>

      <!-- Prompt selector modal -->
      <div v-if="showPromptModal" class="prompt-modal">
        <div class="prompt-modal-content">
          <div class="modal-header">
            <h2>Select a prompt</h2>
            <span class="close-modal" @click="showPromptModal = false">×</span>
          </div>
          <div class="prompt-list">
            <div 
              v-for="(promptOption, idx) in availablePrompts" 
              :key="idx" 
              class="prompt-option"
              @click="selectPrompt(promptOption)"
            >
              {{ promptOption }}
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- Next button -->
    <div class="next-button" @click="goToNextStep" :class="{ 'button-disabled': !allPromptsCompleted }">
      <span class="next-icon">›</span>
    </div>

    <!-- Bottom bar -->
    <div class="bottom-bar"></div>
  </div>
</template>

<script>
export default {
  name: 'InputPromptView',
  data() {
    return {
      selectedPrompts: [
        { question: '', answer: '' },
        { question: '', answer: '' },
        { question: '', answer: '' }
      ],
      availablePrompts: [
        'Two truths and a lie...',
        'I geek out on...',
        'My most irrational fear is...',
        'The one thing I want to know about you is...',
        'My most controversial opinion is...',
        'Together, we could...',
        "I'll pick the topic if you'll pick the place...",
        'A life goal of mine is...',
        'My simple pleasures are...',
        'The key to my heart is...'
      ],
      showPromptModal: false,
      currentPromptIndex: 0
    }
  },
  computed: {
    allPromptsCompleted() {
      return this.selectedPrompts.every(prompt => 
        prompt.question && 
        prompt.answer.trim().length > 0 && 
        prompt.answer.length <= 150
      );
    }
  },
  methods: {
    openPromptSelector(index) {
      this.currentPromptIndex = index;
      this.showPromptModal = true;
    },
    selectPrompt(promptText) {
      // Check if this prompt is already selected in another slot
      const isDuplicate = this.selectedPrompts.some((p, idx) => 
        idx !== this.currentPromptIndex && p.question === promptText
      );
      
      if (!isDuplicate) {
        this.selectedPrompts[this.currentPromptIndex].question = promptText;
        this.showPromptModal = false;
      } else {
        alert('You already selected this prompt. Please choose another one.');
      }
    },
    goToNextStep() {
      if (this.allPromptsCompleted) {
        // Navigate to next step, passing the prompts data
        this.$emit('prompts-submitted', this.selectedPrompts);
        // Or use router: this.$router.push('/next-step');
      }
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
  margin-bottom: 10px;
  color: black;
  text-align: left;
}

.instruction-text {
  margin-bottom: 30px;
  color: #666;
}

.prompt-container {
  margin-bottom: 25px;
  border: 1px solid #eee;
  border-radius: 10px;
  padding: 15px;
  background-color: #f9f9f9;
}

.prompt-header {
  margin-bottom: 15px;
}

.prompt-selector {
  display: flex;
  justify-content: space-between;
  font-weight: bold;
  cursor: pointer;
  color: #5d3fd3;
}

.dropdown-icon {
  font-size: 12px;
}

.input-group {
  margin-bottom: 10px;
  position: relative;
}

.text-input {
  width: 100%;
  padding: 10px 0;
  border: none;
  background: transparent;
  font-size: 16px;
  color: black;
}

.prompt-answer {
  resize: none;
  min-height: 80px;
  padding: 10px;
  border: 1px solid #eee;
  border-radius: 5px;
  background-color: white;
}

.text-input::placeholder {
  color: #ccc;
  font-style: italic;
}

.input-line {
  height: 1px;
  width: 100%;
  background-color: #ddd;
  margin-top: 5px;
}

.character-count {
  text-align: right;
  font-size: 12px;
  color: #999;
  margin-top: 5px;
}

.limit-reached {
  color: red;
}

.prompt-modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.prompt-modal-content {
  width: 90%;
  max-width: 500px;
  max-height: 80vh;
  background-color: white;
  border-radius: 10px;
  overflow: hidden;
  display: flex;
  flex-direction: column;
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 15px;
  border-bottom: 1px solid #eee;
}

.close-modal {
  font-size: 24px;
  cursor: pointer;
}

.prompt-list {
  overflow-y: auto;
  padding: 10px;
}

.prompt-option {
  padding: 15px;
  border-bottom: 1px solid #eee;
  cursor: pointer;
}

.prompt-option:hover {
  background-color: #f5f5f5;
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
