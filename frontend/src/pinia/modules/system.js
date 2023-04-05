/*
 * @Author: iRorikon
 * @Date: 2023-04-05 19:40:23
 * @FilePath: \http-file-web\src\pinia\modules\system.js
 */

import { defineStore } from 'pinia'
import { getSystemConfig } from '../../api/system'
import { ref, watch } from 'vue'

export const useSystemStore = defineStore('system', () => {
  const loadingInstance = ref(null)
  const systemInfo = ref({
    title: '文件管理系统',
    version: '1.0.0',
    logo: 'https://img.oez.cc/2020/12/19/0f8b57866bdb5.gif',
    footer_text: '©2023 iRorikon',
    footer_link: 'https://github.com/iRorikon',
    music_img: 'https://img.oez.cc/2020/12/19/0f8b57866bdb5.gif',
    script: '',
    autoplay: false,
    sort: 'name-descend',
    preview: [],
    expire: 720000,
    base_color: '#fff'
  })
  const token = ref(localStorage.getItem('token') || '')
  const setSystemInfo = val => {
    systemInfo.value = val
  }

  const setToken = val => {
    token.value = val
  }

  const ResetSystemInfo = (value = {}) => {
    systemInfo.value = {
      ...systemInfo.value,
      ...value
    }
  }

  /* 获取系统信息 */
  const GetSystemInfo = async () => {
    const res = await getSystemConfig()
    if (res.code === 200) {
      setSystemInfo(res.data.system_info)
    }
    return res
  }

  /* 清理数据 */
  const ClearStorage = async () => {
    token.value = ''
    sessionStorage.clear()
    localStorage.clear()
  }

  /* 基础模式 */
  const BaseColor = computed(() => {
    if (systemInfo.value.base_color === 'dark') {
      return '#fff'
    } else if (systemInfo.value.base_color === 'light') {
      return '#191a23'
    } else {
      return systemInfo.value.base_color
    }
  })

  watch(
    () => token.value,
    () => {
      window.localStorage.setItem('token', token.value)
    }
  )
  return {
    systemInfo,
    token,
    setToken,
    ResetSystemInfo,
    GetSystemInfo,
    ClearStorage,
    BaseColor
  }
})
