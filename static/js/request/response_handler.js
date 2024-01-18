import HttpStatusHandler from './http_status_handler.js'
/**
 * Handling Response return by fetch function
 *
 * @return {{
 *    run: (Promise(), HttpStatusHandler()) => {}
 * }}
 */
export default function ResponseHandler () {
  const handleResponse = (res = Promise(), statusHandler = HttpStatusHandler()) => {
    const d = res.json()
    if (res.status === 200) {
      statusHandler.OK()
      return d
    } else if (res.status === 400) {
      statusHandler.BadRequest()
      return d.then(Promise.reject.bind(Promise))
    } else if (res.status === 401) {
      statusHandler.Unauthorized()
      return d.then(Promise.reject.bind(Promise))
    } else if (res.status === 404) {
      alert('找不到資料')
      return d.then(Promise.reject.bind(Promise))
    } else {
      alert('系統錯誤，請重新操作')
      return d.then(Promise.reject.bind(Promise))
    }
  }

  return {
    run: (res = Promise(), statusHandler = HttpStatusHandler()) => handleResponse(res, statusHandler)
  }
}
