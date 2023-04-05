/*
 * @Author: iRorikon
 * @Date: 2023-04-05 19:56:14
 * @FilePath: \http-file-web\src\api\storage.js
 */
import request from '../utils/request'

export const getStorageFileList = (path_name, password) => {
  return request({
    url: 'storage/list',
    method: 'post',
    data: {
      path: path_name,
      password: password
    }
  })
}

export const createStorageDir = (path_name, password) => {
  return request({
    url: 'storage/create',
    method: 'post',
    data: {
      path: path_name,
      password: password
    }
  })
}

export const updateStorageDir = (new_path_name, old_path_name, password) => {
  return request({
    url: 'storage/update',
    method: 'post',
    data: {
      new_path: new_path_name,
      old_path: old_path_name,
      password: password
    }
  })
}

export const deleteStorageFile = (path_name, dir) => {
  return request({
    url: 'storage/delete',
    method: 'post',
    data: {
      path: path_name,
      password: dir
    }
  })
}
