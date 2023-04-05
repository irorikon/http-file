/*
 * @Author: iRorikon
 * @Date: 2023-04-05 19:53:05
 * @FilePath: \http-file-web\src\api\system.js
 */
import request from '../utils/request'

export const getSystemConfig = () => {
  return request({
    url: '/system/info',
    method: 'get'
  })
}

export const setSystemConfig = data => {
  return request({
    url: '/system/set',
    method: 'post',
    data
  })
}

export const getSystemState = () => {
  return request({
    url: '/system/state',
    method: 'post'
  })
}
