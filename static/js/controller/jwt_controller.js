import HttpStatusHandler from '../request/http_status_handler.js'
import ResponseHandler from '../request/response_handler.js'
import UserToken from '../cookie/user_token.js'

/**
 * Async function to handle user login.
 *
 * @param {string} bodyStr - optional string parameter for the body
 * @return {Promise} a Promise that resolves after handling the login request
 */
export async function login (bodyStr = '') {
  const statusHandler = HttpStatusHandler()
  statusHandler.Unauthorized = () => console.log('User unauthorized')
  statusHandler.BadRequest = () => alert('User not exist')
  statusHandler.Forbidden = () => alert('Password incorrect')

  return fetch('/api/v1/jwt/login', {
    method: 'POST',
    body: bodyStr,
    headers: new Headers({
      'Content-Type': 'application/json'
    })
  }).then(res => ResponseHandler().run(res, statusHandler))
    .then(data => {
      if (data !== undefined) {
        UserToken().set(data.Token)
        location.reload()
      }
    })
}
