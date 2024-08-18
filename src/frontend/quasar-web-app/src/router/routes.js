const MainLayout = () => import("layouts/MainLayout.vue");
const IndexPage = () => import("pages/IndexPage.vue");
const ErrorNotFound = () => import("pages/ErrorNotFound.vue");

const routes = [
  {
    path: "/",
    component: MainLayout,
    children: [{ path: "", component: IndexPage }],
  },
  {
    path: '/users',
    component: () => import('pages/UserList.vue'),
  },
  {
    path: '/users/:id',
    component: () => import('pages/UserDetail.vue'),
  },
  {
    path: '/users/create',
    component: () => import('pages/UserCreate.vue'),
  },
  {
    path: "/:catchAll(.*)*",
    component: ErrorNotFound,
  },
];

export default routes;
