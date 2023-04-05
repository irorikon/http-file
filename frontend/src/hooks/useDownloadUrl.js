/*
 * @Author: iRorikon
 * @Date: 2023-04-04 21:45:07
 * @FilePath: \http-file\frontend\src\hooks\useDownloadUrl.js
 */
import { backendUrl } from '../utils/const'
import { copyToClip } from '../utils/copy_clip'
import { message } from 'ant-design-vue'
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import { useStore } from 'vuex'
import { Md5 } from 'ts-md5/dist/md5'

export const useDownloadUrl = () => {
  const store = useStore()
  const route = useRoute()
  const downloadUrl = computed(() => {
    let url = backendUrl + 'd' + encodeURI(decodeURI(route.path))
    const file = store.state.data
    if (file.password === 'y' && store.state.password) {
      const md5 = Md5.hashStr(store.state.password)
      url += '?pw=' + md5.substring(8, 24)
    }
    return url
  })
  const copyFileLink = () => {
    copyToClip(downloadUrl.value)
    message.success('链接已复制到剪贴板.')
  }
  return {
    downloadUrl,
    copyFileLink
  }
}

export const useDownloadFile = () => {
  const store = useStore()
  const getFileDownLink = file => {
    let url = backendUrl + 'd/' + encodeURI(file.dir + file.name)
    if (file.password === 'y' && store.state.password) {
      const md5 = Md5.hashStr(store.state.password)
      url += '?pw=' + md5.substring(8, 24)
    }
    return url
  }

  const copyFileLink = file => {
    copyToClip(getFileDownLink(file))
    message.success('链接已复制到剪贴板.')
  }
  return {
    getFileDownLink,
    copyFileLink
  }
}

export default useDownloadFile
