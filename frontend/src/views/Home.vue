<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();

const navigateToLogin = () => {
  router.push('/login');
};

const navigateToRegister = () => {
  router.push('/register');
};
</script>

<template>
  <div class="home-container">
    <div class="hero-section">
      <div class="hero-content">
        <h1 class="app-title">Vibe</h1>
        <p class="app-tagline">Connect with people who match your vibe</p>
        
        <div class="cta-buttons">
          <button class="btn btn-primary" @click="navigateToRegister">Sign Up</button>
          <button class="btn btn-secondary" @click="navigateToLogin">Login</button>
        </div>
      </div>
      
      <div class="hero-image">
        <div class="phone-mockup">
          <div class="mockup-content">
            <!-- App screenshot/preview would go here -->
            <div class="app-preview">
              <div class="profile-card">
                <div class="profile-image"></div>
                <div class="profile-info">
                  <h3>Sarah, 28</h3>
                  <p>Software Developer</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <div class="features-section">
      <h2>How It Works</h2>
      
      <div class="features-grid">
        <div class="feature-card">
          <div class="feature-icon">👤</div>
          <h3>Create Profile</h3>
          <p>Sign up and customize your profile to showcase your personality.</p>
        </div>
        
        <div class="feature-card">
          <div class="feature-icon">🔍</div>
          <h3>Discover Matches</h3>
          <p>Browse through potential matches based on your preferences.</p>
        </div>
        
        <div class="feature-card">
          <div class="feature-icon">💬</div>
          <h3>Connect & Chat</h3>
          <p>Match with people you like and start meaningful conversations.</p>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.home-container {
  padding-top: 2rem;
}

.hero-section {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 4rem;
}

.hero-content {
  flex: 1;
  padding-right: 2rem;
}

.app-title {
  font-size: 3.5rem;
  margin-bottom: 1rem;
  color: #fd5d8d;
}

.app-tagline {
  font-size: 1.5rem;
  color: #666;
  margin-bottom: 2rem;
}

.cta-buttons {
  display: flex;
  gap: 1rem;
}

.btn {
  padding: 0.75rem 1.5rem;
  font-size: 1rem;
  font-weight: 500;
  border-radius: 50px;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-primary {
  background-color: #fd5d8d;
  color: white;
  border: none;
}

.btn-primary:hover {
  background-color: #e54b7b;
  transform: translateY(-2px);
}

.btn-secondary {
  background-color: white;
  color: #fd5d8d;
  border: 2px solid #fd5d8d;
}

.btn-secondary:hover {
  background-color: #fff5f8;
  transform: translateY(-2px);
}

.hero-image {
  flex: 1;
  display: flex;
  justify-content: center;
}

.phone-mockup {
  width: 280px;
  height: 560px;
  background-color: #222;
  border-radius: 36px;
  padding: 12px;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.2);
  position: relative;
}

.mockup-content {
  width: 100%;
  height: 100%;
  background-color: white;
  border-radius: 28px;
  overflow: hidden;
}

.app-preview {
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  background-color: #f8f8f8;
}

.profile-card {
  width: 85%;
  height: 70%;
  background-color: white;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  overflow: hidden;
}

.profile-image {
  height: 75%;
  background-color: #ddd;
  background-image: linear-gradient(45deg, #fd5d8d, #fc8da3);
}

.profile-info {
  padding: 1rem;
}

.profile-info h3 {
  margin: 0 0 0.25rem 0;
  font-size: 1.25rem;
}

.profile-info p {
  margin: 0;
  color: #666;
}

.features-section {
  padding: 4rem 0;
  text-align: center;
}

.features-section h2 {
  font-size: 2rem;
  margin-bottom: 3rem;
  color: #333;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 2rem;
}

.feature-card {
  background-color: white;
  border-radius: 12px;
  padding: 2rem;
  box-shadow: 0 6px 12px rgba(0, 0, 0, 0.05);
  transition: transform 0.3s;
}

.feature-card:hover {
  transform: translateY(-10px);
}

.feature-icon {
  font-size: 2.5rem;
  margin-bottom: 1rem;
}

.feature-card h3 {
  margin-bottom: 1rem;
  color: #333;
}

.feature-card p {
  color: #666;
  line-height: 1.5;
}

@media (max-width: 900px) {
  .hero-section {
    flex-direction: column;
  }
  
  .hero-content {
    padding-right: 0;
    margin-bottom: 3rem;
    text-align: center;
  }
  
  .cta-buttons {
    justify-content: center;
  }
  
  .features-grid {
    grid-template-columns: 1fr;
    max-width: 400px;
    margin: 0 auto;
  }
}
</style> 