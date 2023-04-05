/*
 * @Author: iRorikon
 * @Date: 2023-04-04 21:24:44
 * @FilePath: \http-file\frontend\src\utils\const.js
 */
import { getUrl } from './get_url'

export const fileExtensions = {
  exe: 'windows',
  xls: 'file-excel',
  xlsx: 'file-excel',
  md: 'file-markdown',
  pdf: 'file-pdf',
  ppt: 'file-ppt',
  pptx: 'file-ppt',
  doc: 'file-word',
  docx: 'file-word',
  jpg: 'file-jpg',
  zip: 'file-zip',
  gz: 'file-zip',
  rar: 'file-zip',
  '7z': 'file-zip',
  tar: 'file-zip',
  jar: 'file-zip',
  xz: 'file-zip',
  apk: 'android',
  dmg: 'apple',
  ipa: 'apple',
  m3u8: 'youtube'
}

export const doc = ['pdf', 'doc', 'docx', 'ppt', 'pptx', 'xls', 'xlsx']

export const categorys = {
  image: 'file-image',
  doc: 'file-text',
  video: 'youtube',
  audio: 'customer-service'
}

// export const backendUrl =
//   process.env.VUE_APP_API_URL != '/' ? process.env.VUE_APP_API_URL : getUrl('')

export const backendUrl = 'http://localhost:5244/'
