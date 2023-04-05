/*
 * @Author: iRorikon
 * @Date: 2023-04-04 19:39:23
 * @FilePath: \http-file\frontend\vite.config.js
 */
import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'

import * as path from 'path'
// https://vitejs.dev/config/
// export default defineConfig({
//   plugins: [vue()]
//   require('dotenv').config({ path: `./.env.${mode}` });
//   // resolve: {
//   //   alias: {
//   //     '@': path.resolve(__dirname, 'src'),
//   //     '#': path.resolve(__dirname, 'node_modules')
//   //   }
//   // }
// })
const envDir = path.resolve(__dirname, '.env')
const loadEnvVariables = mode => {
  Object.assign(process.env, loadEnv(mode, envDir, ''))
}

export default defineConfig(init => {
  loadEnvVariables(init.mode)
  return {
    plugins: [vue()],
    define: {
      'process.env': {}
    }
  }
})
