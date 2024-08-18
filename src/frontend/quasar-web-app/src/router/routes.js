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
    path: "/:catchAll(.*)*",
    component: ErrorNotFound,
  },
];

export default routes;
