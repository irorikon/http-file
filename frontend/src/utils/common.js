/*
 * @Author: iRorikon
 * @Date: 2023-04-07 21:32:28
 * @FilePath: \http-file\frontend\src\utils\common.js
 */
import { categorys, fileTypeList, backendUrl } from './const'
import { encodePath } from '~/utils/base64'

export const isEmpty = val => {
  if (typeof val === 'string') {
    return val.trim() === ''
  } else if (typeof val === 'number') {
    return val === 0
  } else if (typeof val === 'boolean') {
    return val === false
  } else if (Array.isArray(val)) {
    return val.length === 0
  } else if (typeof val === 'object') {
    return Object.keys(val).length === 0
  } else if (typeof val === 'function') {
    return false
  } else {
    return true
  }
}

export const isNotEmpty = val => {
  return !isEmpty(val)
}

export const getIcon = record => {
  if (record.is_dir) {
    return 'folder'
  }
  if (Object.prototype.hasOwnProperty.call(fileTypeList, record.type)) {
    return fileTypeList[record.type]
  }
  if (Object.prototype.hasOwnProperty.call(fileTypeList, record.ext)) {
    return fileTypeList[record.ext]
  }
  return 'file'
}

export const getCategory = record => {
  if (record.is_dir) {
    return 'folder'
  }
  if (Object.prototype.hasOwnProperty.call(categorys, record.type)) {
    return categorys[record.type]
  }
  if (Object.prototype.hasOwnProperty.call(categorys, record.ext)) {
    return categorys[record.ext]
  }
}

export const getFileSize = size => {
  if (!size) return ''

  const num = 1024.0

  if (size < num) {
    return size + 'B'
  } else if (size < Math.pow(num, 2)) {
    return (size / num).toFixed(2) + 'KB'
  } else if (size < Math.pow(num, 3)) {
    return (size / Math.pow(num, 2)).toFixed(2) + 'MB'
  } else if (size < Math.pow(num, 4)) {
    return (size / Math.pow(num, 3)).toFixed(2) + 'GB'
  } else {
    return (size / Math.pow(num, 4)).toFixed(2) + 'TB'
  }
}

export const getDownloadLink = file => {
  if (isNotEmpty(file)) {
    return backendUrl + 'file/d/' + encodePath(file)
  } else {
    setTimeout(() => {
      return backendUrl + 'file/d/' + encodePath(file)
    }, 200)
  }
}

export const copyToClipboard = text => {
  navigator.clipboard
    .writeText(text)
    .then(() => {
      console.log('Copied to clipboard')
    })
    .catch(error => {
      console.error('Error copying to clipboard:', error)
    })
}

export const byteToString = bytes => {
  if (typeof bytes === 'string') {
    return bytes
  }
  var str = ''
  var _bytes = bytes
  for (var i = 0; i < _bytes.length; i++) {
    var one = _bytes[i].toString(2)
    var v = one.match(/^1+?(?=0)/)
    if (v && one.length === 8) {
      var bytesLength = v[0].length
      var store = _bytes[i].toString(2).slice(7 - bytesLength)
      for (var st = 1; st < bytesLength; st++) {
        store += _bytes[st + i].toString(2).slice(2)
      }
      str += String.fromCharCode(parseInt(store, 2))
      i += bytesLength - 1
    } else {
      str += String.fromCharCode(_bytes[i])
    }
  }
  return str
}

export const getUrl = url => {
  if (!url) {
    url = window.location.href
  }
  const urlArr = url.split('/')
  return urlArr.slice(0, 3).join('/') + '/'
}

export const getFileLinkEncode = file => {
  if (isNotEmpty(file)) {
    return backendUrl + 'file/d/' + encodePath(file)
  }
}

export const getFileLink = file => {
  if (isNotEmpty(file)) {
    return backendUrl + 'file/d' + file.path
  }
}
