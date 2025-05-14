<template>
  <div class="welcome-container">
    <div class="content" style="text-align: left; padding-left: 20px;">
      <h1 class="title" style="color: black; text-align: left; padding-left: 10px;">Welcome to Hinge.</h1>
      <p class="subtitle" style="color: black; text-align: left; padding-left: 10px;">
        The app <br />
        designed to be <br />
        deleted.<span class="cursor">|</span>
      </p>
      
      <div class="illustration" style="text-align: left; padding-left: 10px;">
        <img src="https://via.placeholder.com/300x200" alt="Placeholder illustration" style="max-width: 100%; height: auto;" />
      </div>
    </div>
    
    <div class="bottom-button">
      <button class="continue-button">Continue</button>
      <div class="home-indicator"></div>
    </div>
  </div>
</template>

<script setup>
// Import the image
</script>

<style scoped>
.welcome-container {
  height: 100vh;
  width: 100%;
  background-color: white;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  padding: 2rem;
  font-family: 'Serif', serif; /* Replace with actual font if known */
}

.content {
  margin-top: 3rem;
}

.title {
  font-size: 2.5rem;
  font-weight: bold;
  margin-bottom: 1.5rem;
}

.subtitle {
  font-size: 2rem;
  font-weight: bold;
  line-height: 1.2;
  margin-bottom: 2rem;
}

.cursor {
  color: #7456e8; /* Purple cursor color */
  animation: blink 1s infinite;
}

@keyframes blink {
  0%, 100% { opacity: 1; }
  50% { opacity: 0; }
}

.illustration {
  margin-top: 2rem;
  /* Note: You'll need to create or source the illustration shown in the screenshot */
}

.bottom-button {
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 1rem;
}

.continue-button {
  width: 100%;
  background-color: #7456e8; /* Purple button color */
  color: white;
  padding: 1rem;
  border: none;
  border-radius: 0;
  font-size: 1rem;
  font-weight: bold;
  cursor: pointer;
}

.home-indicator {
  width: 30%;
  height: 5px;
  background-color: #e0e0e0;
  border-radius: 3px;
  margin-top: 0.5rem;
}
</style>
