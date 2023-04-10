/*
 * @Author: iRorikon
 * @Date: 2023-04-04 19:39:23
 * @FilePath: \http-file\frontend\src\main.js
 */
import { createApp } from 'vue'

import {
  Button,
  Divider,
  Tag,
  Card,
  Breadcrumb,
  Table,
  Modal,
  Input,
  Result,
  Spin,
  Popover,
  Space,
  Switch,
  ConfigProvider
} from 'ant-design-vue'

import {
  LockOutlined,
  UserOutlined,
  HomeOutlined,
  WindowsOutlined,
  FileExcelOutlined,
  FileMarkdownOutlined,
  FilePdfOutlined,
  FilePptOutlined,
  FileWordOutlined,
  FileZipOutlined,
  FileJpgOutlined,
  AndroidOutlined,
  AppleOutlined,
  FileImageOutlined,
  FileTextOutlined,
  YoutubeOutlined,
  CustomerServiceOutlined,
  FileOutlined,
  FolderFilled,
  CopyOutlined,
  DownloadOutlined,
  RetweetOutlined
} from '@ant-design/icons-vue'

import 'ant-design-vue/dist/antd.css'
import './assets/global.css'
import App from './App.vue'
import router from './router'
import store from './pinia'

import contextmenu from 'vue3-contextmenu'
import 'vue3-contextmenu/dist/vue3-contextmenu.css'

const app = createApp(App)

app.use(contextmenu)
app.use(Button)
app.use(Divider)
app.use(Tag)
app.use(Card)
app.use(Breadcrumb)
app.use(Table)
app.use(Modal)
app.use(Input)
app.use(Result)
app.use(Spin)
app.use(Popover)
app.use(Space)
app.use(Switch)
app.use(ConfigProvider)
app.component('lock-outlined', LockOutlined)
app.component('user-outlined', UserOutlined)
app.component('home', HomeOutlined)
app.component('windows', WindowsOutlined)
app.component('file-excel', FileExcelOutlined)
app.component('file-markdown', FileMarkdownOutlined)
app.component('file-pdf', FilePdfOutlined)
app.component('file-ppt', FilePptOutlined)
app.component('file-word', FileWordOutlined)
app.component('file-zip', FileZipOutlined)
app.component('android', AndroidOutlined)
app.component('apple', AppleOutlined)
app.component('file-image', FileImageOutlined)
app.component('file-jpg', FileJpgOutlined)
app.component('file-text', FileTextOutlined)
app.component('youtube', YoutubeOutlined)
app.component('customer-service', CustomerServiceOutlined)
app.component('file', FileOutlined)
app.component('folder', FolderFilled)
app.component('copy', CopyOutlined)
app.component('download', DownloadOutlined)
app.component('retweet', RetweetOutlined)

/**
 * @description 导入加载进度条，防止首屏加载时间过长，用户等待
 *
 * */
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'
NProgress.configure({ showSpinner: false, ease: 'ease', speed: 500 })
NProgress.start()
window.addEventListener('load', () => {
  NProgress.done()
})

app.use(NProgress).use(store).use(router).mount('#app')

export default app
