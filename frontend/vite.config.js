/*
 * @Author: iRorikon
 * @Date: 2023-04-04 19:39:23
 * @FilePath: \http-file\frontend\vite.config.js
 */
import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'

import * as path from 'path'

export default defineConfig(env => {
  // env 环境变量
  const viteEnv = loadEnv(env.mode, process.cwd())

  return {
    base: viteEnv.VITE_BASE,
    // 插件
    plugins: [vue()],
    define: {
      'process.env': {}
    },
    // 别名设置
    resolve: {
      alias: {
        '~': path.resolve(__dirname, './src') // 把 @ 指向到 src 目录去
      }
    },
    server: {
      host: true, // host设置为true才可以使用network的形式，以ip访问项目
      port: process.env.VITE_BASE_PORT, // 端口号
      open: true, // 自动打开浏览器
      cors: true, // 跨域设置允许
      strictPort: true // 如果端口已占用直接退出
      // 接口代理
      // proxy: {
      //   '/api': {
      //     // 本地 8000 前端代码的接口 代理到 8888 的服务端口
      //     target: 'http://localhost:5244/',
      //     changeOrigin: true, // 允许跨域
      //     rewrite: path => path.replace('/api/', '/')
      //   }
      // }
    },
    build: {
      reportCompressedSize: false,
      // 消除打包大小超过500kb警告
      chunkSizeWarningLimit: 2000,
      minify: 'esbuild',
      assetsDir: 'static/assets',
      // 静态资源打包到dist下的不同目录
      rollupOptions: {
        output: {
          chunkFileNames: 'static/js/[hash].js',
          entryFileNames: 'static/js/[hash].js',
          assetFileNames: 'static/[ext]/[hash].[ext]'
        }
      }
    }
  }
})
