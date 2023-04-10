/*
 * @Author: iRorikon
 * @Date: 2023-04-06 20:00:21
 * @FilePath: \http-file\frontend\src\router\index.js
 */
import { createRouter, createWebHistory } from 'vue-router'
import exceptionRoutes from './route.exception'
import NProgress from 'nprogress'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import(/* webpackChunkName: "home" */ '~/views/Home.vue')
  },
  ...exceptionRoutes
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

/**
 * @description: 全局路由前置守卫，在进入路由前触发，导航在所有守卫 resolve 完之前一直处于等待中。
 * @param {RouteLocationNormalized} to  即将要进入的目标
 * @param {RouteLocationNormalizedLoaded} from  当前导航正在离开的路由
 * @return {*}
 */
router.beforeEach((to, from) => {
  // 设置页面标题
  document.title = to.meta.title || import.meta.env.VITE_APP_TITLE
  if (!NProgress.isStarted()) {
    NProgress.start()
  }
})

router.afterEach((to, from) => {
  NProgress.done()
})

router.onError(() => {
  NProgress.remove()
})

export default router
