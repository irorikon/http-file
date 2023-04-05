import { backendUrl } from '../../utils/const'
import loadJS from '../../utils/load_js'
import { message } from 'ant-design-vue'
import { infoGet, pathPost, searchPost } from '../../utils/api'
import { Md5 } from 'ts-md5/dist/md5'
import { defineStore } from 'pinia'
import { ref, computed, watch } from 'vue'

export const siteStore = defineStore('site', () => {
  const loading = ref(null)

  const siteInfo = ref({
    title: '',
    roots: [],
    logo: '',
    footer_text: '',
    footer_url: '',
    music_img: '',
    check_update: '',
    script: '',
    autoplay: false,
    sort: '',
    preview: {
      url: '',
      pre_process: [],
      extensions: [],
      text: [],
      max_size: 0
    }
  })
  const token = ref(window.localStorage.getItem('token') || '')
  const setInfo = val => {
    siteInfo.value = val
  }
  const setToken = val => {
    token.value = val
  }
  const GetSiteInfo = async () => {
    try {
      const res = await infoGet(backendUrl)
      if (res.code === 0) {
        setInfo(res.data)
      } else {
        message.error(res.msg)
      }
    } catch (error) {
      console.log(error)
    }
  }
  return {
    loading,
    siteInfo,
    token,
    setInfo,
    setToken,
    GetSiteInfo
  }
})
