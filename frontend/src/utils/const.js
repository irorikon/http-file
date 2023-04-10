/*
 * @Author: iRorikon
 * @Date: 2023-04-04 21:24:44
 * @FilePath: \http-file\frontend\src\utils\const.js
 */
import { getUrl } from '~/utils/common'

export const fileTypeList = {
  exe: 'windows',
  md: 'file-markdown',
  pdf: 'file-pdf',
  epub: 'file-pdf',
  ppt: 'file-ppt',
  pptx: 'file-ppt',
  odp: 'file-ppt',
  doc: 'file-word',
  docx: 'file-word',
  odt: 'file-word',
  xls: 'file-excel',
  xlsx: 'file-excel',
  ods: 'file-excel',
  jpg: 'file-jpg',
  jp2: 'file-jpg',
  png: 'file-jpg',
  gif: 'file-jpg',
  webp: 'file-jpg',
  cr2: 'file-jpg',
  tif: 'file-jpg',
  bmp: 'file-jpg',
  jxr: 'file-jpg',
  psd: 'file-jpg',
  ico: 'file-jpg',
  heif: 'file-jpg',
  dwg: 'file-jpg',
  exr: 'file-jpg',
  avif: 'file-jpg',
  zip: 'file-zip',
  tar: 'file-zip',
  rar: 'file-zip',
  gz: 'file-zip',
  bz2: 'file-zip',
  '7z': 'file-zip',
  xz: 'file-zip',
  zst: 'file-zip',
  iso: 'file-zip',
  apk: 'android',
  dmg: 'apple',
  ipa: 'apple',
  m3u8: 'youtube',
  mp4: 'youtube',
  m4v: 'youtube',
  mkv: 'youtube',
  webm: 'youtube',
  mov: 'youtube',
  avi: 'youtube',
  wmv: 'youtube',
  mpg: 'youtube',
  flv: 'youtube',
  '3gp': 'youtube',
  mid: 'customer-service',
  mp3: 'customer-service',
  m4a: 'customer-service',
  ogg: 'customer-service',
  flac: 'customer-service',
  wav: 'customer-service',
  amr: 'customer-service',
  aac: 'customer-service',
  aiff: 'customer-service',
  rtf: 'font',
  eot: 'font',
  woff: 'font',
  woff2: 'font',
  ttf: 'font',
  otf: 'font'
}

export const categorys = {
  wasm: 'application',
  dex: 'application',
  dey: 'application',
  zip: 'archive',
  tar: 'archive',
  rar: 'archive',
  gz: 'archive',
  bz2: 'archive',
  '7z': 'archive',
  xz: 'archive',
  zst: 'archive',
  exe: 'archive',
  swf: 'archive',
  rtf: 'archive',
  eot: 'archive',
  ps: 'archive',
  sqlite: 'archive',
  nes: 'archive',
  crx: 'archive',
  cab: 'archive',
  deb: 'archive',
  ar: 'archive',
  Z: 'archive',
  lz: 'archive',
  rpm: 'archive',
  elf: 'archive',
  dcm: 'archive',
  iso: 'archive',
  macho: 'archive',
  mid: 'audio',
  mp3: 'audio',
  m4a: 'audio',
  ogg: 'audio',
  flac: 'audio',
  wav: 'audio',
  amr: 'audio',
  aac: 'audio',
  aiff: 'audio',
  doc: 'doc',
  docx: 'doc',
  xls: 'doc',
  xlsx: 'doc',
  ppt: 'doc',
  pptx: 'doc',
  odp: 'doc',
  ods: 'doc',
  odt: 'doc',
  epub: 'doc',
  pdf: 'doc',
  woff: 'font',
  woff2: 'font',
  ttf: 'font',
  otf: 'font',
  jpg: 'image',
  jp2: 'image',
  png: 'image',
  gif: 'image',
  webp: 'image',
  cr2: 'image',
  tif: 'image',
  bmp: 'image',
  jxr: 'image',
  psd: 'image',
  ico: 'image',
  heif: 'image',
  dwg: 'image',
  exr: 'image',
  avif: 'image',
  mp4: 'video',
  m4v: 'video',
  mkv: 'video',
  webm: 'video',
  mov: 'video',
  avi: 'video',
  wmv: 'video',
  mpg: 'video',
  flv: 'video',
  '3gp': 'video'
}

export const doc = ['pdf', 'doc', 'docx', 'ppt', 'pptx', 'xls', 'xlsx']

export const backendUrl =
  import.meta.env.VITE_BASE_API_URL != '/'
    ? import.meta.env.VITE_BASE_API_URL
    : getUrl('')
