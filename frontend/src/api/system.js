/*
 * @Author: iRorikon
 * @Date: 2023-04-05 19:53:05
 * @FilePath: \http-file\frontend\src\api\system.js
 */
import request from '../utils/request'

export const getSiteConfig = () => {
  return request({
    url: '/site/get',
    method: 'get'
  })
}

export const updateSiteConfig = data => {
  return request({
    url: '/site/update',
    method: 'post',
    data
  })
}

export const getSystemConfig = () => {
  return request({
    url: '/system/get',
    method: 'get'
  })
}

export const updateSystemConfig = data => {
  return request({
    url: '/system/update',
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
