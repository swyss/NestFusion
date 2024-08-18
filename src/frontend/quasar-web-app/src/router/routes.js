import UserPage from "pages/UserPage.vue";

const MainLayout = () => import("layouts/MainLayout.vue");
const IndexPage = () => import("pages/IndexPage.vue");
const ErrorNotFound = () => import("pages/ErrorNotFound.vue");
const UserList = () => import("components/UserList.vue");
const UserDetails = () => import("components/UserDetails.vue");
const UserForm = () => import("components/UserForm.vue");

const routes = [
  {
    path: "/",
    component: MainLayout,
    children: [{path: "", component: IndexPage}],
  },
  {
    path: "/user",
    component: UserPage,
    children: [
      {path: "", component: UserList}, // List all users
      {path: ":id", component: UserDetails}, // View user details
      {path: ":id/edit", component: UserForm}, // Edit user
      {path: "new", component: UserForm}, // Create new user
    ],
  },
  {
    path: "/:catchAll(.*)*",
    component: ErrorNotFound,
  },
];

export default routes;
