/*
 * @Author: iRorikon
 * @Date: 2023-04-03 21:04:38
 * @FilePath: \http-file\vite.config.js
 */
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

import * as path from 'path'
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src'),
      '#': path.resolve(__dirname, 'node_modules')
    }
  }
})
