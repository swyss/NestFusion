const routes = [
  {
    path: '/login',
    name: 'UserLogin',
    component: () => import('pages/user/Login.vue')
  },
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/IndexPage.vue') },
      { path: '/account', name: 'UserAccount', component: () => import('pages/user/Account.vue') }
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
