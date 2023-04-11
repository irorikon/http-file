/*
 * @Author: iRorikon
 * @Date: 2023-04-07 13:42:57
 * @FilePath: \http-file\frontend\src\pinia\modules\file.js
 */
import { defineStore } from 'pinia'
import { useGlobalStore } from '~/pinia/modules/global'
import NProgress from 'nprogress'
import { getFileList, getFileinfo } from '~/api/file'

export const useFileStore = defineStore('file', {
  state: () => ({
    filelist: [],
    file: {}
  }),
  actions: {
    setFilelist (filelist) {
      this.filelist = filelist
    },
    setFile (file) {
      this.file = file
    },
    async getFiles (path, storageType) {
      NProgress.start()
      const globalStore = useGlobalStore()
      // 获取文件列表
      globalStore.setLoading(true)
      globalStore.setPath(path)
      const res = await getFileList({
        storage_type: storageType,
        path: path
      })
      if (res.code === 0) {
        this.setFilelist(res.data)
        globalStore.setLoading(false)
        // 清空缓存
        this.setFile({})
        NProgress.done()
      }
    },
    async getFile (path, storageType) {
      // 获取文件列表
      const res = await getFileinfo({
        storage_type: storageType,
        path: path
      })
      if (res.code === 0) {
        this.setFile(res.data)
      }
    }
  }
})
