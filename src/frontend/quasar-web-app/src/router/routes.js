// Dynamically imported components
const MainLayout = () => import("layouts/MainLayout.vue");
const ErrorNotFound = () => import("pages/error/ErrorNotFound.vue");
const IndexPage = () => import('pages/IndexPage.vue');
const UserPage = () => import('pages/user/UserPage.vue');
const AdminDashboard = () => import('pages/user/AdminDashboard.vue');
const UserProfile = () => import('pages/user/UserProfile.vue');

// Extracted child routes of MainLayout
const mainLayoutChildren = [
  {path: '', component: IndexPage},
];

const userLayoutChildren = [
  {path: '', component: UserPage},
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
    path: "/user",
    component: MainLayout,
    children: userLayoutChildren
  },
  {
    path: "/:catchAll(.*)*",
    component: ErrorNotFound,
  },
];

export default routes;
