<template>
  <div class="profile-container">
    <!-- Phone status bar -->
    <div class="status-bar">
      <span class="time">9:41</span>
      <div class="status-icons">
        <span class="signal">●●●●</span>
        <span class="wifi">📶</span>
      </div>
    </div>

    <!-- App header -->
    <div class="app-header">
      <h1 class="app-title">Hinge</h1>
      <div class="header-controls">
        <button class="filter-button">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4" />
          </svg>
        </button>
        <button class="settings-button">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
          </svg>
        </button>
      </div>
    </div>

    <!-- Profile image and name -->
    <div class="profile-header">
      <div class="profile-image-container">
        <div class="profile-image">
          <img src="https://randomuser.me/api/portraits/women/44.jpg" alt="Profile picture" />
          <button class="edit-profile" style="position: absolute; top: 0; right: 0; bottom: auto;">
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7l4-4m0 0l4 4m-4-4v18" transform="rotate(135, 12, 12)" />
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 20h16" />
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 4l-4 4" />
            </svg>
          </button>
        </div>
        <h2 class="profile-name">Julia <span class="verified-badge">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="#1DA1F2">
            <path d="M9 16.17L4.83 12l-1.42 1.41L9 19 21 7l-1.41-1.41L9 16.17z"/>
          </svg>
        </span></h2>
      </div>
    </div>

    <!-- Navigation tabs -->
    <div class="navigation-tabs">
      <button class="tab active">Get More</button>
      <button class="tab">Safety</button>
      <button class="tab">My Hinge</button>
    </div>

    <!-- Subscription cards -->
    <div class="subscription-cards">
      <!-- HingeX Card -->
      <div class="premium-card">
        <div class="premium-card-content">
          <h3 class="premium-title">HingeX</h3>
          <p class="premium-description">Get seen sooner and go on 3x as many dates</p>
          <button class="premium-action">Check it out</button>
        </div>
      </div>

      <!-- Boost Card -->
      <div class="boost-card">
        <div class="boost-icon">
          <span class="boost-counter">6</span>
        </div>
        <div class="boost-content">
          <h3>Boost</h3>
          <p>Get seen by 11X more people</p>
        </div>
      </div>

      <!-- Roses Card -->
      <div class="roses-card">
        <div class="roses-icon">
          <span class="roses-counter">4</span>
        </div>
        <div class="roses-content">
          <h3>Roses</h3>
          <p>2x as likely to lead to a date</p>
        </div>
      </div>
    </div>

    <!-- Bottom Navigation -->
    <div class="bottom-navigation">
      <button class="nav-button">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
        </svg>
      </button>
      <button class="nav-button">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z" />
        </svg>
      </button>
      <button class="nav-button">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
        </svg>
        <span class="notification-badge">3</span>
      </button>
      <button class="nav-button">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
        </svg>
        <span class="notification-badge">1</span>
      </button>
      <button class="nav-button active">
        <img src="https://via.placeholder.com/24" alt="Profile" class="profile-nav-icon" />
      </button>
    </div>
  </div>
</template>

<script setup>
// Component logic
</script>

<style scoped>
.profile-container {
  max-width: 430px;
  margin: 0 auto;
  position: relative;
  background-color: white;
  border-radius: 40px;
  overflow: hidden;
  height: 100vh;
  color: #000;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
}

.status-bar {
  display: flex;
  justify-content: space-between;
  padding: 12px 24px 0;
  height: 44px;
}

.time {
  font-weight: 600;
}

.app-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 16px;
}

.app-title {
  font-size: 24px;
  font-weight: bold;
  margin: 0;
}

.header-controls {
  display: flex;
  gap: 16px;
}

.filter-button, .settings-button {
  background: none;
  border: none;
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 50%;
  background-color: #f0f0f0;
  cursor: pointer;
}

.profile-header {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 16px 0;
}

.profile-image-container {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.profile-image {
  position: relative;
  width: 120px;
  height: 120px;
  border-radius: 50%;
  border: 4px solid #9a40b0;
  overflow: hidden;
}

.profile-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.edit-profile {
  position: absolute;
  bottom: 0;
  right: 0;
  z-index: 1000;
  background-color: white;
  border: 2px solid #9a40b0;
  border-radius: 50%;
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0;
}

.profile-name {
  font-size: 24px;
  font-weight: 600;
  margin: 8px 0;
  display: flex;
  align-items: center;
  gap: 8px;
}

.verified-badge {
  background-color: #f0f0f0;
  border-radius: 50%;
  width: 24px;
  height: 24px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.navigation-tabs {
  display: flex;
  justify-content: space-around;
  border-bottom: 1px solid #e0e0e0;
  margin-bottom: 16px;
}

.tab {
  flex: 1;
  padding: 12px 0;
  text-align: center;
  background: none;
  border: none;
  font-size: 16px;
  color: #888;
  cursor: pointer;
}

.tab.active {
  color: #000;
  border-bottom: 2px solid #000;
  font-weight: 500;
}

.subscription-cards {
  padding: 0 16px;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.premium-card {
  background-image: linear-gradient(rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)), url('https://via.placeholder.com/400x150');
  background-size: cover;
  color: white;
  border-radius: 12px;
  padding: 24px;
  height: 150px;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.premium-card-content {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.premium-title {
  font-size: 24px;
  font-weight: 600;
  margin: 0 0 8px;
}

.premium-description {
  font-size: 16px;
  margin: 0 0 16px;
  text-align: center;
}

.premium-action {
  background-color: white;
  color: black;
  font-weight: 500;
  border: none;
  border-radius: 20px;
  padding: 8px 16px;
  cursor: pointer;
}

.boost-card, .roses-card {
  display: flex;
  align-items: center;
  padding: 16px;
  border-radius: 12px;
  border: 1px solid #e0e0e0;
}

.boost-icon, .roses-icon {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 16px;
  position: relative;
}

.boost-icon {
  background-color: #2e8a7d;
  color: white;
}

.roses-icon {
  background-color: #c85a86;
  color: white;
}

.boost-counter, .roses-counter {
  position: absolute;
  top: -4px;
  right: -4px;
  background-color: white;
  border: 1px solid #e0e0e0;
  border-radius: 50%;
  width: 20px;
  height: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  font-weight: 600;
  color: black;
}

.boost-content, .roses-content {
  flex: 1;
}

.boost-content h3, .roses-content h3 {
  margin: 0 0 4px;
  font-size: 16px;
}

.boost-content p, .roses-content p {
  margin: 0;
  font-size: 14px;
  color: #666;
}

.bottom-navigation {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  max-width: 430px;
  margin: 0 auto;
  background-color: white;
  display: flex;
  justify-content: space-around;
  padding: 12px 0;
  border-top: 1px solid #e0e0e0;
}

.nav-button {
  background: none;
  border: none;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  position: relative;
  color: #888;
  padding: 8px;
}

.nav-button.active {
  color: #000;
}

.notification-badge {
  position: absolute;
  top: 0;
  right: 0;
  background-color: #ff3b30;
  color: white;
  border-radius: 50%;
  width: 20px;
  height: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 12px;
  font-weight: 600;
}

.profile-nav-icon {
  width: 24px;
  height: 24px;
  border-radius: 50%;
}

/* Home indicator for modern iPhones */
.profile-container::after {
  content: "";
  display: block;
  height: 5px;
  width: 30%;
  background-color: #e0e0e0;
  border-radius: 3px;
  margin: 8px auto;
  position: absolute;
  bottom: 8px;
  left: 0;
  right: 0;
}
</style>
