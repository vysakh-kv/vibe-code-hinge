<script setup>
import { useRouter } from 'vue-router';

const router = useRouter();
</script>

<template>
  <div class="photos-input-screen">
    <!-- Progress indicator -->
    <div class="progress-indicator">
      <div class="step-dot"></div>
      <div class="active-step">
        <div class="step-icon">
          <span class="icon-content">🖼️</span>
        </div>
      </div>
      <div class="step-dot"></div>
      <div class="step-dot"></div>
      <div class="step-dot"></div>
    </div>
    
    <!-- Photos headline -->
    <div class="input-form">
      <h1>Pair your photos and videos with prompts</h1>
      
      <div class="photo-grid">
        <div v-for="(photo, index) in photos" :key="index" 
             class="photo-item" 
             @click="openPromptDialog(index)">
          <img v-if="photo.base64" :src="photo.base64" alt="Uploaded photo" />
          <div v-else class="add-photo-placeholder">
            <span class="plus-icon">+</span>
          </div>
          <div v-if="photo.prompt" class="prompt-indicator"></div>
        </div>
      </div>
      
      <div class="upload-info">
        <span class="required-count">{{ requiredCount }} required</span>
        <span class="drag-info">Drag to reorder</span>
      </div>
      
      <div class="tooltip-box">
        <div class="lightbulb-icon">💡</div>
        <p>Tap a photo to add a prompt and make your profile stand out even more.</p>
      </div>
    </div>

    <!-- Prompt dialog -->
    <div v-if="showPromptDialog" class="prompt-dialog">
      <div class="prompt-dialog-content">
        <h3>Add a prompt to your photo</h3>
        <textarea v-model="currentPrompt" placeholder="Write something about this photo..."></textarea>
        <div class="dialog-buttons">
          <button @click="cancelPrompt">Cancel</button>
          <button @click="savePrompt" class="save-button">Save</button>
        </div>
      </div>
    </div>
    
    <!-- Hidden file input -->
    <input 
      type="file" 
      ref="fileInput" 
      @change="onFileSelected" 
      accept="image/*,video/*" 
      style="display: none"
      multiple
    >

    <!-- Next button -->
    <div class="next-button" 
         @click="isReadyToProgress ? router.push('/input-prompt') : null"
         :class="{ 'enabled': isReadyToProgress }">
      <span class="next-icon">›</span>
    </div>


    <!-- Bottom bar -->
    <div class="bottom-bar"></div>
  </div>
</template>

<script>
export default {
  name: 'InputPhotosView',
  data() {
    return {
      photos: Array(6).fill().map(() => ({ base64: '', prompt: '' })),
      requiredCount: 3,
      showPromptDialog: false,
      currentPhotoIndex: null,
      currentPrompt: '',
    }
  },
  computed: {
    isReadyToProgress() {
      const filledPhotos = this.photos.filter(photo => photo.base64);
      return filledPhotos.length >= this.requiredCount;
    }
  },
  methods: {
    openPromptDialog(index) {
      if (!this.photos[index].base64) {
        // If no photo, trigger file upload
        this.currentPhotoIndex = index;
        this.$refs.fileInput.click();
      } else {
        // If photo exists, open prompt dialog
        this.currentPhotoIndex = index;
        this.currentPrompt = this.photos[index].prompt;
        this.showPromptDialog = true;
      }
    },
    
    cancelPrompt() {
      this.showPromptDialog = false;
      this.currentPrompt = '';
      this.currentPhotoIndex = null;
    },
    
    savePrompt() {
      if (this.currentPhotoIndex !== null) {
        this.photos[this.currentPhotoIndex].prompt = this.currentPrompt;
      }
      this.showPromptDialog = false;
      this.currentPrompt = '';
      this.currentPhotoIndex = null;
    },
    
    onFileSelected(event) {
      const files = event.target.files;
      if (!files.length) return;
      
      const file = files[0];
      const reader = new FileReader();
      
      reader.onload = (e) => {
        if (this.currentPhotoIndex !== null) {
          this.photos[this.currentPhotoIndex].base64 = e.target.result;
        }
      };
      
      reader.readAsDataURL(file);
      
      // Reset file input
      this.$refs.fileInput.value = null;
    },
    
    goToNextStep() {
      // Submit photos and move to next step
      this.$emit('photos-submitted', this.photos);
      // Or use router: this.$router.push('/next-step');
    }
  }
}
</script>

<style scoped>
.photos-input-screen {
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

.photo-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 15px;
  margin-top: 20px;
}

.photo-item {
  position: relative;
  aspect-ratio: 1/1;
  border-radius: 10px;
  border: 2px dashed #ccc;
  overflow: hidden;
  cursor: pointer;
  display: flex;
  justify-content: center;
  align-items: center;
}

.photo-item img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.add-photo-placeholder {
  width: 100%;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.plus-icon {
  font-size: 40px;
  color: #aaa;
}

.prompt-indicator {
  position: absolute;
  bottom: 5px;
  left: 5px;
  width: 10px;
  height: 10px;
  border-radius: 50%;
  background-color: #5d3fd3;
}

.upload-info {
  display: flex;
  justify-content: space-between;
  margin-top: 10px;
  font-size: 14px;
  color: #666;
}

.tooltip-box {
  margin-top: 30px;
  border: 1px solid #eee;
  border-radius: 10px;
  padding: 20px;
  display: flex;
  align-items: center;
  background-color: #fafafa;
}

.lightbulb-icon {
  font-size: 24px;
  margin-right: 15px;
}

.tooltip-box p {
  margin: 0;
  font-size: 16px;
  line-height: 1.4;
}

.prompt-dialog {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 100;
}

.prompt-dialog-content {
  background-color: white;
  padding: 20px;
  border-radius: 15px;
  width: 80%;
  max-width: 400px;
}

.prompt-dialog h3 {
  margin-top: 0;
  font-size: 18px;
}

.prompt-dialog textarea {
  width: 100%;
  height: 100px;
  border: 1px solid #ddd;
  border-radius: 10px;
  padding: 10px;
  resize: none;
  font-family: inherit;
  margin: 15px 0;
}

.dialog-buttons {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}

.dialog-buttons button {
  padding: 10px 20px;
  border-radius: 20px;
  border: none;
  cursor: pointer;
}

.dialog-buttons .save-button {
  background-color: #5d3fd3;
  color: white;
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
