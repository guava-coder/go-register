import HttpStatusHandler from '../request/http_status_handler.js'
import ResponseHandler from '../request/response_handler.js'

/**
 * Advance fetch
 * @return {{
 * get: (args=RequestArgs()) => { },
 * post: (args=RequestArgs()) => {},
 * put: (args=RequestArgs()) => {},
 * delete: (args=RequestArgs()) => {}
 * }}
 */
export default function Ajaj () {
  const ajax = async (met = '', arg = RequestArgs()) => {
    return fetch(arg.url, {
      method: met,
      body: arg.bodyStr,
      headers: arg.headers
    }).then(res => ResponseHandler().run(res, arg.statusHandler))
  }

  return {
    get: (args = RequestArgs()) => { return ajax('GET', args) },
    post: (args = RequestArgs()) => { return ajax('POST', args) },
    put: (args = RequestArgs()) => { return ajax('PUT', args) },
    delete: (args = RequestArgs()) => { return ajax('DELETE', args) }
  }
}

/**
 *
 *
 * @return {{
 * url: string,
 * bodyStr: string,
 * statusHandler: HttpStatusHandler(),
 * headers: string
 * }}
 */
function RequestArgs () {
  return {
    url: '',
    bodyStr: '',
    statusHandler: HttpStatusHandler(),
    headers: ''
  }
}
