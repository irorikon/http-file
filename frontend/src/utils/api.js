/*
 * @Author: iRorikon
 * @Date: 2023-04-04 16:22:08
 * @FilePath: \http-file\frontend\src\utils\api.js
 */
import service from '@/utils/request'

export const getFileList = params => {
  return service({
    url: '/file/list',
    method: 'get',
    params
  })
}

export const getFileInfo = params => {
  return service({
    url: '/file/info',
    method: 'get',
    params
  })
}

export const uploadFile = params => {
  return service({
    url: '/file/upload',
    method: 'post',
    params
  })
}

export const downloadFile = params => {
  return service({
    url: '/file/download',
    method: 'get',
    params
  })
}

export const deleteFile = params => {
  return service({
    url: '/file/delete',
    method: 'delete',
    params
  })
}
