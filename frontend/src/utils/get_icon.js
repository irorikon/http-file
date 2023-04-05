/*
 * @Author: iRorikon
 * @Date: 2023-04-04 21:36:55
 * @FilePath: \http-file\frontend\src\utils\get_icon.js
 */
import { FileProps } from '../store'
import { categorys, fileExtensions } from './const'

export const getIcon = record => {
  if (record.type === 'folder') {
    return 'folder'
  }
  if (Object.prototype.hasOwnProperty.call(categorys, record.category)) {
    return categorys[record.category]
  }
  if (
    Object.prototype.hasOwnProperty.call(fileExtensions, record.file_extension)
  ) {
    return fileExtensions[record.file_extension]
  }
  return 'file'
}
