/*
 * @Author: iRorikon
 * @Date: 2023-04-04 21:08:12
 * @FilePath: \http-file\frontend\src\utils\request.js
 */
import { message } from 'ant-design-vue'
import axios from 'axios'

const instance = axios.create({
  // baseURL: process.env.VUE_APP_BASE_API,
  baseURL: 'http://localhost:5244/',
  timeout: 5000,
  headers: {
    'Content-Type': 'application/json;charset=utf-8'
  },
  withCredentials: false
})

instance.interceptors.request.use(
  config => {
    return config
  },
  error => {
    console.log('Error: ' + error.message)
    return Promise.reject(error)
  }
)

instance.interceptors.response.use(
  response => {
    return response
  },
  error => {
    console.log(error)
    if (!error.response || error.response.data.meta == undefined) {
      message.error('后端网络异常,请检查后端程序是否运行或检查网络连接!')
      return Promise.reject(error)
    }
    return error.response.data
  }
)

export default instance
