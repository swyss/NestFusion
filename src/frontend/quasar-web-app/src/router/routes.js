// Dynamically imported components
const MainLayout = () => import("layouts/MainLayout.vue");
const ErrorNotFound = () => import("pages/error/ErrorNotFound.vue");
const IndexPage = () => import('pages/IndexPage.vue');
const UserLogin = () => import('pages/user/UserLogin.vue');
const UserRegister = () => import('pages/user/UserRegister.vue');
const AdminDashboard = () => import('pages/user/AdminDashboard.vue');
const UserProfile = () => import('pages/user/UserProfile.vue');

// Extracted child routes of MainLayout
const mainLayoutChildren = [
  {path: '', component: IndexPage},
  {path: 'login', component: UserLogin},
  {path: 'register', component: UserRegister},
  {path: 'admin', component: AdminDashboard, meta: {requiresAdmin: true}},
  {path: 'profile', component: UserProfile}
];

// Structured and organized routes array
const routes = [
  {
    path: "/",
    component: MainLayout,
    children: mainLayoutChildren
  },
  {
    path: "/:catchAll(.*)*",
    component: ErrorNotFound,
  },
];

export default routes;
