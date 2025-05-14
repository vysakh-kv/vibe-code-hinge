import { createRouter, createWebHistory } from 'vue-router';

// Lazy loading routes for better performance
const Home = () => import('../views/Home.vue');
const Login = () => import('../views/Login.vue');
const Register = () => import('../views/Register.vue');
const EnterPhNumber = () => import('../views/EnterPhNumber.vue');
const VerifyOtp = () => import('../views/verify-otp.vue');
const Welcome = () => import('../views/welcome.vue');
const InputName = () => import('../views/input-name.vue');
const InputEmail = () => import('../views/input-email.vue');
const InputHometown = () => import('../views/input-hometown.vue');
const InputWork = () => import('../views/input-work.vue');
const InputDrink = () => import('../views/input-drink.vue');
const InputSmoke = () => import('../views/input-smoke.vue');
const InputPhotos = () => import('../views/input-photos.vue');
const InputPrompt = () => import('../views/input-prompt.vue');
// const Register = () => import('../views/Register.vue');
// const Onboarding = () => import('../views/Onboarding.vue');
// const Discover = () => import('../views/Discover.vue');
// const Matches = () => import('../views/Matches.vue');
// const Messages = () => import('../views/Messages.vue');
const Profile = () => import('../views/Profile.vue');
// const Settings = () => import('../views/Settings.vue');

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    meta: { requiresAuth: false }
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
    meta: { requiresAuth: false }
  },
  {
    path: '/register',
    name: 'Register',
    component: Register,
    meta: { requiresAuth: false }
  },
  {
    path: '/enter-number',
    name: 'EnterNumber',
    component: EnterPhNumber,
    meta: { requiresAuth: false }
  },
  {
    path: '/verify-otp',
    name: 'VerifyOtp',
    component: VerifyOtp,
    meta: { requiresAuth: false }
  },
  {
    path: '/welcome',
    name: 'Welcome',
    component: Welcome,
    meta: { requiresAuth: false }
  },
  {
    path: '/input-name',
    name: 'InputName',
    component: InputName,
    meta: { requiresAuth: false }
  },
  {
    path: '/input-email',
    name: 'InputEmail',
    component: InputEmail,
    meta: { requiresAuth: false }
  },
  {
    path: '/input-hometown',
    name: 'InputHometown',
    component: InputHometown,
    meta: { requiresAuth: false }
  },
  {
    path: '/input-work',
    name: 'InputWork',
    component: InputWork,
    meta: { requiresAuth: false }
  },
  {
    path: '/input-drink',
    name: 'InputDrink',
    component: InputDrink,
    meta: { requiresAuth: false }
  },
  {
    path: '/input-smoke',
    name: 'InputSmoke',
    component: InputSmoke,
    meta: { requiresAuth: false }
  },
  {
    path: '/input-photos',
    name: 'InputPhotos',
    component: InputPhotos,
    meta: { requiresAuth: false }
  },
  {
    path: '/input-prompt',
    name: 'InputPrompt',
    component: InputPrompt,
    meta: { requiresAuth: false }
  },
  // {
  //   path: '/onboarding',
  //   name: 'Onboarding',
  //   component: Onboarding,
  //   meta: { requiresAuth: true }
  // },
  // {
  //   path: '/discover',
  //   name: 'Discover',
  //   component: Discover,
  //   meta: { requiresAuth: true }
  // },
  // {
  //   path: '/matches',
  //   name: 'Matches',
  //   component: Matches,
  //   meta: { requiresAuth: true }
  // },
  // {
  //   path: '/messages/:matchId',
  //   name: 'Messages',
  //   component: Messages,
  //   meta: { requiresAuth: true }
  // },
  {
    path: '/profile',
    name: 'Profile',
    component: Profile,
    meta: { requiresAuth: false }
  },
  // {
  //   path: '/settings',
  //   name: 'Settings',
  //   component: Settings,
  //   meta: { requiresAuth: true }
  // }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

// Navigation guard for auth
router.beforeEach((to, from, next) => {
  const isAuthenticated = !!localStorage.getItem('user-token');
  
  if (to.meta.requiresAuth && !isAuthenticated) {
    next({ name: 'Login' });
  } else if (
    (to.name === 'Login' || to.name === 'Register') && 
    isAuthenticated
  ) {
    next({ name: 'Discover' });
  } else {
    next();
  }
});

export default router; 
