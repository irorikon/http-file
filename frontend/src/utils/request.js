/*
 * @Author: iRorikon
 * @Date: 2023-04-04 21:08:12
 * @FilePath: \http-file\frontend\src\utils\request.js
 */
import { message } from 'ant-design-vue'
import axios from 'axios'

// 创建实例
const instance = axios.create({
  // 前缀
  baseURL: import.meta.env.VITE_BASE_API_URL,
  // 超时
  timeout: 0,
  // 请求头
  headers: {
    'Content-Type': 'application/json;charset=utf-8'
  },
  withCredentials: false
})

// 请求拦截器
instance.interceptors.request.use(
  config => {
    // TODO 在这里可以加上想要在请求发送前处理的逻辑
    // TODO 比如 loading 等
    return config
  },
  error => {
    console.log('Error: ' + error.message)
    return Promise.reject(error)
  }
)

// 响应拦截器
instance.interceptors.response.use(
  response => {
    if (response.status === 200) {
      return response.data
    } else {
      message.error('Error: ' + response.status)
      return Promise.reject(response)
    }
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
