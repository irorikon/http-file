/*
 * @Author: iRorikon
 * @Date: 2023-04-07 13:42:57
 * @FilePath: \http-file\frontend\src\pinia\modules\file.js
 */
import { defineStore } from 'pinia'
import { useGlobalStore } from '~/pinia/modules/global'
import { useSiteStore } from '~/pinia/modules/site'
import { getFileList, getFileinfo } from '~/api/file'

export const useFileStore = defineStore('file', {
  state: () => ({
    filelist: [],
    file: {},
    audio: {
      name: '',
      url: '',
      cover: ''
    }
  }),
  actions: {
    setFilelist (filelist) {
      this.filelist = filelist
    },
    setFile (file) {
      this.file = file
    },
    setAudio (audio) {
      this.audio = {
        name: audio.name,
        url: backendUrl + 'v/' + encodeURI(file.dir + file.name),
        cover:
          useSiteStore.siteInfo.music_img ||
          'https://img.oez.cc/2020/12/19/0f8b57866bdb5.gif'
      }
    },
    async getFiles (path, storageType) {
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
