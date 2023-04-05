import { backendUrl } from '../utils/const'
import loadJS from '../utils/load_js'
import { message } from 'ant-design-vue'
import { createStore } from 'vuex'
import { infoGet, pathPost, searchPost } from '../utils/api'
import { Md5 } from 'ts-md5/dist/md5'

const MetaProps = {
  code: 0,
  msg: ''
}

const InfoProps = {
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
}

export const FileProps = {
  dir: '',
  file_extension: '',
  file_id: '',
  name: '',
  type: '',
  updated_at: '',
  category: '',
  content_type: '',
  size: 0,
  password: '',
  sizeStr: '',
  time: '',
  icon: '',
  content_hash: ''
}

const Audio = {
  name: '',
  url: '',
  cover: ''
}

export const GlobalDataProps = {
  loading: false,
  info: InfoProps,
  password: '',
  meta: MetaProps,
  data: [],
  type: '',
  audios: [],
  drive: '',
  isImages: false,
  showImages: false,
  isMultiple: false
}

export default createStore({
  state: {
    info: {},
    loading: true,
    password: localStorage.getItem('password') || '',
    meta: {
      code: 200
    },
    data: [],
    type: 'loading',
    audios: [],
    drive: '',
    isImages: false,
    showImages: false,
    isMultiple: localStorage.getItem('isMultiple') === 'true'
  },
  mutations: {
    setIsMultiple (state, isMultiple) {
      localStorage.setItem('isMultiple', isMultiple ? 'true' : 'false')
      state.isMultiple = isMultiple
    },
    setShowImages (state, showImages) {
      state.showImages = showImages
    },
    setLoading (state, loading) {
      state.loading = loading
    },
    setPassword (state, password) {
      state.password = password
      localStorage.setItem('password', password)
    },
    setInfo (state, info) {
      state.info = info
    },
    setMeta (state, meta) {
      state.meta = meta
    },
    setDrive (state, drive) {
      state.drive = drive
    },
    setData (state, data) {
      state.isImages = false
      if (!data) {
        state.type = 'no'
        state.data = []
        return
      }
      if (data.type) {
        state.type = 'file'
      } else {
        state.type = 'folder'
        const audios = []
        const files = data
        console.log(files)
        for (const file of files) {
          if (file.category === 'audio') {
            const md5 = Md5.hashStr(state.password)
            audios.push({
              name: file.name,
              url:
                backendUrl +
                'd/' +
                encodeURI(file.dir + file.name) +
                '?pw=' +
                md5.substring(8, 24),
              cover:
                state.info.music_img ||
                'https://img.oez.cc/2020/12/19/0f8b57866bdb5.gif'
            })
          }
        }
        state.isImages =
          files.filter(item => item.category === 'image').length > 0
        state.audios = audios
      }
      state.data = data
    }
  },
  actions: {
    async fetchInfo ({ state, commit }) {
      const { data } = await infoGet()
      const infoData = data.data
      document.title = infoData.title || '文件管理系统'
      if (infoData.script) {
        await loadJS(infoData.script)
      }
      if (!infoData.logo) {
        infoData.logo = '../assets/alist.png'
      }
      commit('setInfo', infoData)
    },
    async fetchPathOrSearch ({ state, commit }, { path, query }) {
      commit('setLoading', true)
      state.type = 'loading'
      if (query) {
        const { data } = await searchPost(query)
        if (data.code !== 200) {
          message.error(data.msg)
        }
        commit('setData', data.data)
      } else {
        const { data } = await pathPost(path, state.password)
        commit('setData', data)
        if (data.code !== 200) {
          message.error(data.msg)
        }
        if (data.code === 401) {
          return
        }
        commit('setData', data.data)
      }
      commit('setLoading', false)
    }
  },
  getters: {},
  modules: {}
})
