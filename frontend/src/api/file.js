/*
 * @Author: iRorikon
 * @Date: 2023-04-06 20:00:21
 * @FilePath: \http-file\frontend\src\api\file.js
 */
import request from '~/utils/request'

export const getFileinfo = data => {
  return request({
    url: 'file/info',
    method: 'post',
    data
  })
}

export const getFileList = data => {
  return request({
    url: 'file/list',
    method: 'post',
    data
  })
}

export const renameFile = data => {
  return request({
    url: 'file/copy',
    method: 'post',
    data
  })
}

export const removeFile = data => {
  return request({
    url: 'file/remove',
    method: 'post',
    data
  })
}

export const copyFile = data => {
  return request({
    url: 'file/copy',
    method: 'post',
    data
  })
}

export const moveFile = data => {
  return request({
    url: 'file/move',
    method: 'post',
    data
  })
}

export const searchFile = data => {
  return request({
    url: 'file/search',
    method: 'post',
    data
  })
}
