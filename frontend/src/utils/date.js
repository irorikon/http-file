/*
 * @Author: iRorikon
 * @Date: 2023-04-04 21:30:40
 * @FilePath: \http-file\frontend\src\utils\date.js
 */
/**
 * full
 * 用于月日时分秒小于10补0
 * @param {*} p
 */

export const full = p => {
  return p < 10 ? `0${p}` : p
}
/**
 * formatDate
 * 将timestamp日期格式化为 yyyy-mm-dd hh-mm-ss
 * @param {*} date
 */
export const formatDate = dateStr => {
  const date = new Date(dateStr)
  const year = date.getFullYear()
  const month = date.getMonth() + 1
  const day = date.getDate()
  const hour = date.getHours()
  const minute = date.getMinutes()
  const second = date.getSeconds()
  return `${year}-${full(month)}-${full(day)} ${full(hour)}:${full(
    minute
  )}:${full(second)}`
}
