/*
 * @Author: iRorikon
 * @Date: 2023-04-07 12:55:56
 * @FilePath: \http-file\frontend\src\pinia\modules\global.js
 */
import { defineStore } from 'pinia'

export const useGlobalStore = defineStore('global', {
  // type of file(folder or file) default folder
  state: () => ({
    loading: false,
    showPassword: false,
    storage_type: '',
    path: '',
    paths: {}
  }),
  actions: {
    setLoading (val) {
      this.loading = val
    },
    setType (val) {
      this.type = val
    },
    setShowPassword (val) {
      this.showPassword = val
    },
    setStorageType (val) {
      this.storage_type = val
    },
    setPath (val) {
      this.path = val
    },
    setPaths (val) {
      let list = []
      val.split('/').forEach(item => {
        if (item === '') {
          console.debug('empty')
        } else {
          list.push(item)
        }
      })
      this.paths = list
    }
  }
})
