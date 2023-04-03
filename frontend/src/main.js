/*
 * @Author: iRorikon
 * @Date: 2023-04-03 21:04:38
 * @FilePath: \http-file\src\main.js
 */
import { createApp } from 'vue'
import 'vue-devui/style.css'
import '@devui-design/icons/icomoon/devui-icon.css'
import App from './App.vue'
import DevUI from 'vue-devui'
import { ThemeServiceInit, infinityTheme } from 'devui-theme'

ThemeServiceInit({ infinityTheme }, 'infinityTheme')

createApp(App).use(DevUI).mount('#app')
