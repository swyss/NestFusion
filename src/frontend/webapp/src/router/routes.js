const routes = [
  {
    path: '/',
    component: () => import('layouts/main/MainLayout.vue'),
    children: [
      { path: '', name: 'home',component: () => import('pages/IndexPage.vue') },
      { path: 'account', name: 'account', component: () => import('pages/user/Account.vue')},
      { path: 'login', name: 'login', component: () => import('pages/user/Login.vue') },
      { path: "components", name:"components", component: () => import("pages/examples/NotificationComponent.vue"),
      },
    ],

  },
  {
    path: '/user',
    component: () => import('layouts/user/UserLayout.vue'),
    children: [
      { path: '', name: 'UserHome',component: () => import('pages/IndexPage.vue') },
      { path: 'account', name: 'UserAccount', component: () => import('pages/user/Account.vue')},
      { path: 'login', name: 'UserLogin', component: () => import('pages/user/Login.vue') }
    ],

  },
  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue')
  }
]


export default routes
