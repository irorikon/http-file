import request from '../utils/request'

export const getFileinfo = (filename, password) => {
  return request({
    url: 'file/info',
    method: 'post',
    data: {
      filename: filename,
      password: password
    }
  })
}

export const updateFile = (new_filename, old_filename, password) => {
  return request({
    url: 'file/update',
    method: 'post',
    data: {
      new_filename: new_filename,
      old_filename: old_filename,
      password: password
    }
  })
}

export const deleteFile = (password, filename) => {
  return request({
    url: 'file/delete',
    method: 'post',
    data: {
      filename: filename,
      password: password
    }
  })
}
