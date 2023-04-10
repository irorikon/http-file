/*
 * @Author: iRorikon
 * @Date: 2023-04-06 20:56:01
 * @FilePath: \http-file\frontend\src\router\route.exception.js
 */
const exceptionRoutes = [
  {
    path: '/401',
    name: '401',
    component: () => import('~/views/error/401.vue')
  },
  {
    path: '/403',
    name: '403',
    component: () => import('~/views/error/403.vue')
  },
  {
    path: '/404',
    name: '404',
    component: () => import('~/views/error/404.vue')
  },
  {
    path: '/:pathMatch(.*)',
    redirect: '/404'
  }
]

export default exceptionRoutes
