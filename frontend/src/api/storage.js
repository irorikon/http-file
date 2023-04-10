/*
 * @Author: iRorikon
 * @Date: 2023-04-05 19:56:14
 * @FilePath: \http-file\frontend\src\api\storage.js
 */
import request from '../utils/request'

export const getStorageConfig = data => {
  return request({
    url: 'storage/get',
    method: 'post',
    data
  })
}

export const updateStorageConfig = data => {
  return request({
    url: 'storage/update',
    method: 'post',
    data
  })
}
