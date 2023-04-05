/*
 * @Author: iRorikon
 * @Date: 2023-04-04 21:33:51
 * @FilePath: \http-file\frontend\src\utils\file_size.js
 */
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
