import HttpStatusHandler from './http_status_handler.js'
/**
 * Handling Response return by fetch function
 *
 * @return {{
 * run: (res = Response, statusHandler = HttpStatusHandler()) => {}
 * }}
 */
export default function ResponseHandler () {
  return {
    run: (res = Response, statusHandler = HttpStatusHandler()) => {
      const d = res.json()
      const rejectPromise = () => { return Promise.reject.bind(Promise) }

      switch (res.status) {
        case 200 :
          statusHandler.OK()
          return d
        case 400:
          statusHandler.BadRequest()
          return d.then(rejectPromise())
        case 401:
          statusHandler.Unauthorized()
          return d.then(rejectPromise())
        case 403:
          statusHandler.Forbidden()
          return d.then(rejectPromise())
        case 404:
          alert('Cannot find resources')
          return d.then(rejectPromise())
        case 500:
          alert('System error, please try again later.')
          return d.then(rejectPromise())
        default:
          alert('Unhandle error.')
          return d.then(rejectPromise())
      }
    }
  }
}
