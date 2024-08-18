const MainLayout = () => import("layouts/MainLayout.vue");
const IndexPage = () => import("pages/IndexPage.vue");
const UserListPage = () => import("pages/user/UserListPage.vue");
const UserDetailPage = () => import("pages/user/UserDetailPage.vue");
const UserFormPage = () => import("pages/UserFormPage.vue");
const ErrorNotFound = () => import("pages/error/ErrorNotFound.vue");

const routes = [
  {
    path: "/",
    component: MainLayout,
    children: [
      { path: "", component: IndexPage },
      { path: "users", component: UserListPage },
      { path: "users/create", component: UserFormPage },
      { path: "users/:id", component: UserDetailPage },
      { path: "users/:id/edit", component: UserFormPage },
    ],
  },
  {
    path: "/:catchAll(.*)*",
    component: ErrorNotFound,
  },
];

export default routes;
