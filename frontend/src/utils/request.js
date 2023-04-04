import axios from 'axios'
import { Message } from 'vue-devui'
import { emitter } from '@/utils/bus.js'

const service = axios.create({
  baseURL: 'http://localhost:5000',
  timeout: 5000
})

let activeAxios = 0
let timer
const showLoading = () => {
  activeAxios++
  if (timer) {
    clearTimeout(timer)
  }
  timer = setTimeout(() => {
    if (activeAxios > 0) {
      emitter.emit('show-loading')
    }
  }, 400)
}

const closeLoading = () => {
  activeAxios--
  if (activeAxios <= 0) {
    clearTimeout(timer)
    emitter.emit('close-loading')
  }
}

// http request 拦截器
service.interceptors.request.use(
  config => {
    if (!config.donNotShowLoading) {
      showLoading()
    }
    config.headers = {
      'Content-Type': 'application/json',
      ...config.headers
    }
    return config
  },
  error => {
    if (!error.config.donNotShowLoading) {
      closeLoading()
    }
    Message({
      message: error,
      type: 'error',
      bordered: false
    })
    return error
  }
)

// http response 拦截器
service.interceptors.response.use(
  response => {
    if (!response.config.donNotShowLoading) {
      closeLoading()
    }
    if (response.data.code === 0 || response.headers.success === 'true') {
      if (response.headers.msg) {
        response.data.msg = decodeURI(response.headers.msg)
      }
      return response.data
    } else {
      Message({
        message: response.data.msg || decodeURI(response.headers.msg),
        type: 'error',
        bordered: false
      })
      return response.data.msg ? response.data : response
    }
  },
  error => {
    if (!error.config.donNotShowLoading) {
      closeLoading()
    }
    if (!error.response) {
      Message({
        message: error,
        type: 'error',
        bordered: false
      })
      return
    }
    switch (error.response.status) {
      case 500:
        Message({
          message: '服务器错误',
          type: 'error',
          bordered: false
        })
      case 404:
        Message({
          message: '资源不存在',
          type: 'error',
          bordered: false
        })
        break
    }
    return error
  }
)
export default service
