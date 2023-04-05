export const getUrl = url => {
  if (!url) {
    url = window.location.href
  }
  const urlArr = url.split('/')
  return urlArr.slice(0, 3).join('/') + '/'
}
