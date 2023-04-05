/*
 * @Author: iRorikon
 * @Date: 2023-04-04 21:51:22
 * @FilePath: \http-file\frontend\src\hooks\useRefresh.js
 */
import { useRoute } from 'vue-router'
import { useStore } from 'vuex'

const useRefresh = () => {
  const route = useRoute()
  const store = useStore()
  const refresh = () => {
    const paths = route.params.path
    if (paths) {
      store.commit('setDrive', paths[0])
    }
    store.dispatch('fetchPathOrSearch', {
      path: decodeURI(route.fullPath.substring(1)).split('?')[0],
      query: route.query['q']
    })
  }
  return {
    refresh
  }
}
export default useRefresh
