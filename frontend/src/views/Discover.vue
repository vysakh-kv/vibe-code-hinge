<template>
  <div class="discover-container">
    <!-- Status bar -->
    <div class="status-bar">
      <span class="time">9:41</span>
      <div class="status-icons">
        <span class="signal">●●●</span>
        <span class="wifi">📶</span>
        <span class="battery">100%</span>
      </div>
    </div>

    <!-- Filter tabs -->
    <div class="filter-tabs">
      <button class="filter-tab active">Compatible</button>
      <button class="filter-tab">Active today</button>
      <button class="filter-tab">Nearby</button>
    </div>

    <!-- Profile card -->
    <div class="profile-card">
      <div class="profile-header">
        <div class="profile-info">
          <h2 class="profile-name">
            {{ currentProfile.name }} 
            <span v-if="currentProfile.verified" class="verified-badge">✓</span>
          </h2>
          <p class="profile-details">{{ currentProfile.pronouns }} · {{ currentProfile.status }}</p>
        </div>
        <div class="profile-actions">
          <button class="action-button">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <circle cx="12" cy="12" r="1"></circle>
              <circle cx="19" cy="12" r="1"></circle>
              <circle cx="5" cy="12" r="1"></circle>
            </svg>
          </button>
        </div>
      </div>

      <div 
        class="profile-image-container"
        @touchstart="handleTouchStart"
        @touchend="handleTouchEnd"
      >
        <div class="image-slider" :style="{ transform: `translateX(-${currentImageIndex * 100}%)` }">
          <img 
            v-for="(image, index) in currentProfile.images" 
            :key="index" 
            :src="image" 
            :alt="`${currentProfile.name} photo ${index + 1}`" 
            class="profile-image"
          >
        </div>

        <!-- Image navigation controls -->
        <div class="image-nav-controls">
          <button v-if="currentProfile.images.length > 1" @click="prevImage" class="image-nav-button prev">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <polyline points="15 18 9 12 15 6"></polyline>
            </svg>
          </button>
          <button v-if="currentProfile.images.length > 1" @click="nextImage" class="image-nav-button next">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <polyline points="9 18 15 12 9 6"></polyline>
            </svg>
          </button>
        </div>

        <!-- Image indicators -->
        <div class="image-indicators">
          <button 
            v-for="(image, index) in currentProfile.images" 
            :key="index"
            class="image-indicator"
            :class="{ active: index === currentImageIndex }"
            @click="goToImage(index)"
          ></button>
        </div>

        <button class="like-button">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"></path>
          </svg>
        </button>
      </div>

      <!-- Bio information -->
      <div v-if="currentProfile.bio" class="profile-bio">
        <p>{{ currentProfile.bio }}</p>
      </div>

      <div class="profile-audio">
        <p class="audio-text">{{ currentProfile.audioPrompt }}</p>
        <div class="audio-player">
          <button class="play-button" @click="togglePlay">
            <svg v-if="!isPlaying" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="currentColor" stroke="none">
              <polygon points="5,3 19,12 5,21"></polygon>
            </svg>
            <svg v-else xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="currentColor" stroke="none">
              <rect x="6" y="4" width="4" height="16"></rect>
              <rect x="14" y="4" width="4" height="16"></rect>
            </svg>
          </button>
          <div class="audio-visualization">
            <div class="audio-wave">
              <div 
                v-for="(height, index) in waveHeights" 
                :key="index" 
                class="wave-bar" 
                :style="{ 
                  height: `${height}px`,
                  opacity: isPlaying ? 1 : 0.5
                }"
              ></div>
            </div>
            <div v-if="duration > 0" class="audio-progress">
              <div class="progress-bar" :style="{ width: `${(currentTime / duration) * 100}%` }"></div>
            </div>
          </div>
        </div>
        <div class="action-buttons">
          <button class="dismiss-button" @click="nextProfile">
            <svg xmlns="http://www.w3.org/2000/svg" width="30" height="30" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <line x1="18" y1="6" x2="6" y2="18"></line>
              <line x1="6" y1="6" x2="18" y2="18"></line>
            </svg>
          </button>
          <button class="heart-button">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M20.84 4.61a5.5 5.5 0 0 0-7.78 0L12 5.67l-1.06-1.06a5.5 5.5 0 0 0-7.78 7.78l1.06 1.06L12 21.23l7.78-7.78 1.06-1.06a5.5 5.5 0 0 0 0-7.78z"></path>
            </svg>
          </button>
        </div>
      </div>
    </div>

    <!-- Bottom navigation -->
    <div class="bottom-nav">
      <button class="nav-button">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"></path>
          <polyline points="9 22 9 12 15 12 15 22"></polyline>
        </svg>
      </button>
      <button class="nav-button">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M19 21l-7-5-7 5V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2z"></path>
        </svg>
      </button>
      <button class="nav-button active">
        <div class="notification-badge">1</div>
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <circle cx="12" cy="12" r="10"></circle>
          <path d="M8 14s1.5 2 4 2 4-2 4-2"></path>
          <line x1="9" y1="9" x2="9.01" y2="9"></line>
          <line x1="15" y1="9" x2="15.01" y2="9"></line>
        </svg>
      </button>
      <button class="nav-button">
        <div class="notification-badge">3</div>
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M21 11.5a8.38 8.38 0 0 1-.9 3.8 8.5 8.5 0 0 1-7.6 4.7 8.38 8.38 0 0 1-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 0 1-.9-3.8 8.5 8.5 0 0 1 4.7-7.6 8.38 8.38 0 0 1 3.8-.9h.5a8.48 8.48 0 0 1 8 8v.5z"></path>
        </svg>
      </button>
      <button class="nav-button">
        <img src="https://placehold.co/30x30/e9e9e9/a1a1a1?text=U" alt="User" class="user-avatar">
      </button>
    </div>

    <!-- Home indicator -->
    <div class="home-indicator"></div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted, computed } from 'vue';

// Audio player functionality
const isPlaying = ref(false);
const currentTime = ref(0);
const duration = ref(0);
const audioElement = ref(null);
const intervalId = ref(null);
const waveHeights = ref(Array(30).fill(0).map(() => Math.random() * 25 + 5));

// Toggle play/pause
function togglePlay() {
  if (!audioElement.value) return;
  
  if (isPlaying.value) {
    audioElement.value.pause();
  } else {
    audioElement.value.play();
  }
  
  isPlaying.value = !isPlaying.value;
}

// Update audio progress
function updateProgress() {
  if (!audioElement.value) return;
  
  currentTime.value = audioElement.value.currentTime;
  duration.value = audioElement.value.duration;
  
  // Animate the wave bars
  waveHeights.value = waveHeights.value.map(() => 
    isPlaying.value ? Math.random() * 25 + 5 : 5
  );
}

// Setup audio element and listeners
onMounted(() => {
  audioElement.value = new Audio('/sample-audio.mp3');
  
  audioElement.value.addEventListener('ended', () => {
    isPlaying.value = false;
    currentTime.value = 0;
    clearInterval(intervalId.value);
  });
  
  audioElement.value.addEventListener('loadedmetadata', () => {
    duration.value = audioElement.value.duration;
  });
  
  // Update progress every 100ms when playing
  intervalId.value = setInterval(() => {
    if (isPlaying.value) {
      updateProgress();
    }
  }, 100);
});

onUnmounted(() => {
  if (intervalId.value) {
    clearInterval(intervalId.value);
  }
  
  if (audioElement.value) {
    audioElement.value.pause();
    audioElement.value.removeEventListener('ended', () => {});
    audioElement.value.removeEventListener('loadedmetadata', () => {});
  }
});

// User profiles data for discovery
const currentProfileIndex = ref(0);
const currentImageIndex = ref(0);

const profiles = ref([
  {
    id: 1,
    name: 'Amelia',
    pronouns: 'she/they',
    status: 'Active today',
    verified: true,
    images: [
      'https://placehold.co/400x500/e9e9e9/a1a1a1?text=Amelia+1',
      'https://placehold.co/400x500/d9d9d9/919191?text=Amelia+2',
      'https://placehold.co/400x500/c9c9c9/818181?text=Amelia+3',
      'https://placehold.co/400x500/b9b9b9/717171?text=Amelia+4'
    ],
    bio: 'Coffee enthusiast, book lover, and hiking addict.',
    audioPrompt: 'I won\'t shut up about'
  },
  {
    id: 2,
    name: 'Alex',
    pronouns: 'he/him',
    status: 'Active today',
    verified: false,
    images: [
      'https://placehold.co/400x500/e0e0e0/909090?text=Alex+1',
      'https://placehold.co/400x500/d0d0d0/808080?text=Alex+2',
      'https://placehold.co/400x500/c0c0c0/707070?text=Alex+3'
    ],
    bio: 'Filmmaker by day, chef by night. Looking for someone to share meals with.',
    audioPrompt: 'My secret talent is'
  },
  {
    id: 3,
    name: 'Jordan',
    pronouns: 'they/them',
    status: 'Active yesterday',
    verified: true,
    images: [
      'https://placehold.co/400x500/f0f0f0/808080?text=Jordan+1',
      'https://placehold.co/400x500/e0e0e0/707070?text=Jordan+2',
      'https://placehold.co/400x500/d0d0d0/606060?text=Jordan+3',
      'https://placehold.co/400x500/c0c0c0/505050?text=Jordan+4',
      'https://placehold.co/400x500/b0b0b0/404040?text=Jordan+5'
    ],
    bio: 'Art curator with a passion for indie music and sustainable fashion.',
    audioPrompt: 'My favorite conversation starter is'
  }
]);

// Current profile being viewed
const currentProfile = computed(() => profiles.value[currentProfileIndex.value]);

// Navigation functions
function nextProfile() {
  if (isPlaying.value) {
    togglePlay(); // Stop audio when moving to next profile
  }
  
  currentProfileIndex.value = (currentProfileIndex.value + 1) % profiles.value.length;
  currentImageIndex.value = 0; // Reset to first image when switching profiles
  currentTime.value = 0;
}

function prevProfile() {
  if (isPlaying.value) {
    togglePlay(); // Stop audio when moving to previous profile
  }
  
  currentProfileIndex.value = (currentProfileIndex.value - 1 + profiles.value.length) % profiles.value.length;
  currentImageIndex.value = 0; // Reset to first image when switching profiles
  currentTime.value = 0;
}

// Image navigation
function nextImage() {
  currentImageIndex.value = (currentImageIndex.value + 1) % currentProfile.value.images.length;
}

function prevImage() {
  currentImageIndex.value = (currentImageIndex.value - 1 + currentProfile.value.images.length) % currentProfile.value.images.length;
}

function goToImage(index) {
  currentImageIndex.value = index;
}

// Touch handling for image swipe
let touchStartX = 0;
let touchEndX = 0;

function handleTouchStart(e) {
  touchStartX = e.changedTouches[0].screenX;
}

function handleTouchEnd(e) {
  touchEndX = e.changedTouches[0].screenX;
  handleSwipe();
}

function handleSwipe() {
  const swipeThreshold = 50;
  if (touchEndX - touchStartX > swipeThreshold) {
    // Swipe right
    prevImage();
  } else if (touchStartX - touchEndX > swipeThreshold) {
    // Swipe left
    nextImage();
  }
}
</script>

<style scoped>
.discover-container {
  max-width: 430px;
  margin: 0 auto;
  height: 100vh;
  background-color: white;
  color: #1f1f1f;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
  display: flex;
  flex-direction: column;
  position: relative;
  overflow: hidden;
}

.status-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 16px;
}

.time {
  font-weight: 600;
}

.status-icons {
  display: flex;
  gap: 5px;
}

.filter-tabs {
  display: flex;
  padding: 10px 16px;
  gap: 15px;
  overflow-x: auto;
}

.filter-tab {
  background: none;
  border: none;
  border-radius: 20px;
  padding: 8px 16px;
  font-size: 14px;
  font-weight: 500;
  white-space: nowrap;
  cursor: pointer;
}

.filter-tab.active {
  background-color: #1f1f1f;
  color: white;
}

.profile-card {
  flex: 1;
  margin: 10px 16px;
  display: flex;
  flex-direction: column;
  background-color: white;
  border-radius: 12px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  overflow: hidden;
}

.profile-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  z-index: 2;
}

.profile-info {
  display: flex;
  flex-direction: column;
}

.profile-name {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 20px;
  font-weight: 600;
  margin: 0;
}

.verified-badge {
  background-color: #8050c7;
  color: white;
  border-radius: 50%;
  width: 18px;
  height: 18px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
}

.profile-details {
  color: #666;
  font-size: 14px;
  margin: 4px 0 0;
}

.action-button {
  background: none;
  border: none;
  cursor: pointer;
  color: #666;
}

.profile-image-container {
  position: relative;
  height: 400px;
  overflow: hidden;
  background-color: #f8f8f8;
  touch-action: pan-x;
}

.image-slider {
  display: flex;
  transition: transform 0.3s ease;
  height: 100%;
  width: 100%;
}

.profile-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  flex-shrink: 0;
}

.image-nav-controls {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 10px;
  pointer-events: none;
}

.image-nav-button {
  background-color: rgba(255, 255, 255, 0.7);
  border: none;
  border-radius: 50%;
  width: 36px;
  height: 36px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  pointer-events: auto;
  transition: background-color 0.2s ease;
}

.image-nav-button:hover {
  background-color: rgba(255, 255, 255, 0.9);
}

.image-indicators {
  position: absolute;
  bottom: 15px;
  left: 0;
  right: 0;
  display: flex;
  justify-content: center;
  gap: 8px;
  z-index: 2;
}

.image-indicator {
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background-color: rgba(255, 255, 255, 0.5);
  border: none;
  padding: 0;
  cursor: pointer;
  transition: all 0.2s ease;
}

.image-indicator.active {
  background-color: white;
  width: 20px;
  border-radius: 3px;
}

.like-button {
  position: absolute;
  bottom: 15px;
  right: 15px;
  background-color: white;
  border: none;
  border-radius: 50%;
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.2);
  cursor: pointer;
  z-index: 2;
}

.profile-bio {
  padding: 12px 16px;
  background-color: white;
  border-bottom: 1px solid #f0f0f0;
}

.profile-bio p {
  margin: 0;
  font-size: 14px;
  line-height: 1.5;
  color: #333;
}

.profile-audio {
  padding: 16px;
  background-color: white;
  position: relative;
}

.audio-text {
  font-size: 16px;
  margin: 0 0 10px;
}

.audio-player {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 15px;
}

.play-button {
  background-color: #8050c7;
  border: none;
  border-radius: 50%;
  width: 36px;
  height: 36px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
  cursor: pointer;
  transition: all 0.2s ease;
}

.play-button:hover {
  background-color: #6b40b1;
}

.audio-visualization {
  flex: 1;
  height: 40px;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.audio-wave {
  display: flex;
  align-items: center;
  gap: 2px;
  width: 100%;
  height: 30px;
}

.wave-bar {
  flex: 1;
  background-color: #8050c7;
  min-width: 2px;
  border-radius: 2px;
  transition: height 0.1s ease-in-out, opacity 0.2s ease;
}

.audio-progress {
  width: 100%;
  height: 2px;
  background-color: #e0e0e0;
  margin-top: 4px;
  border-radius: 1px;
  overflow: hidden;
}

.progress-bar {
  height: 100%;
  background-color: #8050c7;
  transition: width 0.1s linear;
}

.action-buttons {
  display: flex;
  justify-content: space-between;
  padding: 0 20px;
}

.dismiss-button, .heart-button {
  background: none;
  border: none;
  cursor: pointer;
  padding: 0;
  transition: transform 0.2s ease, color 0.2s ease;
}

.dismiss-button {
  color: #666;
}

.dismiss-button:hover {
  color: #ff3b5c;
  transform: scale(1.1);
}

.heart-button {
  color: #666;
}

.heart-button:hover {
  color: #ff3b5c;
  transform: scale(1.1);
}

.bottom-nav {
  display: flex;
  justify-content: space-around;
  padding: 15px 0;
  border-top: 1px solid #eee;
  background-color: white;
}

.nav-button {
  background: none;
  border: none;
  cursor: pointer;
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #666;
  transition: color 0.2s ease;
}

.nav-button:hover {
  color: #8050c7;
}

.nav-button.active {
  color: #8050c7;
}

.notification-badge {
  position: absolute;
  top: -5px;
  right: -5px;
  background-color: #ff3b5c;
  color: white;
  border-radius: 50%;
  width: 18px;
  height: 18px;
  font-size: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.user-avatar {
  width: 24px;
  height: 24px;
  border-radius: 50%;
  object-fit: cover;
}

.home-indicator {
  width: 30%;
  height: 5px;
  background-color: #ddd;
  border-radius: 2.5px;
  margin: 8px auto;
}
</style>
