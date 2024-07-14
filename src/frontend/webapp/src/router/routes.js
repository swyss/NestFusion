const routes = [
  {
    path: '/',
    component: () => import('layouts/main/MainLayout.vue'),
    children: [
      { path: '', name: 'home',component: () => import('pages/IndexPage.vue') },
      { path: 'account', name: 'account', component: () => import('pages/user/Account.vue')},
      { path: 'login', name: 'login', component: () => import('pages/user/Login.vue') }
    ],

  },
  {
    path: '/user',
    component: () => import('layouts/user/UserLayout.vue'),
    children: [
      { path: '', name: 'UserHome',component: () => import('pages/IndexPage.vue') },
      { path: 'account', name: 'UserAccount', component: () => import('pages/user/Account.vue')},
      { path: 'login', name: 'UserLogin', component: () => import('pages/user/Login.vue') },
      { path: 'news', component: () => import('pages/modules/NewsPage.vue') },
      { path: 'weather', component: () => import('pages/modules/WeatherPage.vue')}
    ],

  },
  {
    path: '/error404',
    component: () => import('pages/error/Error404.vue'),
  },
  {
    path: '/error500',
    component: () => import('pages/error/Error500.vue'),
  },
  {
    path: '/error403',
    component: () => import('pages/error/Error403.vue'),
  },
  {
    path: '/error401',
    component: () => import('pages/error/Error401.vue'),
  },
  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/error/ErrorNotFound.vue')
  }
]


export default routes
