/*
 * @Author: iRorikon
 * @Date: 2023-04-05 19:40:23
 * @FilePath: \http-file\frontend\src\pinia\modules\site.js
 */

import { defineStore } from 'pinia'
import { getSiteConfig } from '~/api/system'
import loadJS from '~/utils/load_js'

export const useSiteStore = defineStore('site', {
  state: () => ({
    siteInfo: {
      title: '',
      logo: '',
      footer_text: '',
      footer_link: '',
      script: '',
      sort: 'size-ascend'
    }
  }),
  actions: {
    setSiteInfo (val) {
      this.siteInfo = val
    },
    /* 获取系统信息 */
    async getSiteInfo () {
      const res = await getSiteConfig()
      if (res.code === 0) {
        document.title = res.data.title || '文件管理系统'
        if (res.data.script) {
          await loadJS(res.data.script)
        }
        if (!res.data.logo) {
          res.data.logo = './static/alist.png'
        }
        this.setSiteInfo(res.data)
      }
      return
    }
  }
})
