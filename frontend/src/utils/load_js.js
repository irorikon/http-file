/*
 * @Author: iRorikon
 * @Date: 2023-04-04 21:41:18
 * @FilePath: \http-file\frontend\src\utils\load_js.js
 */
const loadJS = content => {
  return new Promise((resolve, reject) => {
    const script = document.createElement('script')
    script.type = 'text/javascript'
    script.onload = () => {
      resolve()
    }
    script.onerror = () => {
      reject()
    }
    if (/^(http|https):\/\/([\w.]+\/?)\S*/.test(content)) {
      script.src = content
    } else {
      script.text = content
    }
    document.querySelector('body').appendChild(script)
  })
}

export default loadJS
