import Ajaj from '../request/ajaj.js'
import HttpStatusHandler from '../request/http_status_handler.js'

/**
 * Generates a handler function that alerts 'user not exist' on a Bad Request status.
 *
 * @return {HttpStatusHandler} The status handler object.
 */
function getHandler () {
  const statusHandler = HttpStatusHandler()
  statusHandler.BadRequest = () => alert('user not exist')
  return statusHandler
}

/**
 * Sends a verification email with the provided body string.
 *
 * @param {string} bodyStr - The body content of the email.
 * @return {Promise} A Promise that resolves when the email is successfully sent.
 */
export function sendVerificationMail (bodyStr = '') {
  return post({ url: '/api/v1/email/send/verification', bodyStr, statusHandler: getHandler() })
}

/**
 * Verify the email by sending a POST request to '/api/v1/email/verify'.
 *
 * @param {string} bodyStr - the string containing the email body
 * @return {Promise}
 */
export function verifyEmail (bodyStr = '') {
  return post({ url: '/api/v1/email/verify', bodyStr, statusHandler: getHandler() })
}

const ajaj = Ajaj()
const h = new Headers({ 'Content-Type': 'application/json' })
function post (args = {}) {
  args.headers = h; return ajaj.post(args)
}
