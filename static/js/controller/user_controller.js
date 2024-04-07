import Ajaj from '../request/ajaj.js'
import AuthHeaders from '../request/auth_headers.js'
import HttpStatusHandler from '../request/http_status_handler.js'

/**
 * Function to find user data.
 *
 * @return {Promise} User data promise object
 */
export function findUserData () {
  return getUserData({ url: '/api/v1/user/query', bodyStr: '', statusHandler: HttpStatusHandler() })
}

/**
 * Adds a user using the provided body string.
 *
 * @param {string} bodyStr - the body string for the user to be added
 * @return {Promise} a Promise that resolves with the result of the POST request
 */
export function addUser (bodyStr = '') {
  return post({ url: '/api/v1/user/add', bodyStr, statusHandler: HttpStatusHandler() })
}

/**
 * Updates user authentication.
 *
 * @param {string} bodyStr - the body string for the update
 * @return {Promise} the response from the PUT request
 */
export function updateUserAuth (bodyStr = '') {
  const updateAuthHandler = () => {
    const handler = HttpStatusHandler()
    handler.BadRequest = () => alert('Token incorrect')
    return handler
  }

  return put({ url: '/api/v1/user/auth', bodyStr, statusHandler: updateAuthHandler() })
}

/**
 * Update user information with the provided body string.
 *
 * @param {string} bodyStr - the string containing user information to update
 * @return {Promise} the response from the put request
 */
export function updateUserInfo (bodyStr = '') {
  return put({ url: '/api/v1/user/update', bodyStr, statusHandler: HttpStatusHandler() })
}

/**
 * Function that checks the password by calling the API endpoint.
 *
 * @param {string} bodyStr - the body string for the API request
 * @return {Promise} the response from the API call
 */
export function checkPassword (bodyStr = '') {
  return getUserData({ url: '/api/v1/user/check/password', bodyStr, statusHandler: HttpStatusHandler() })
}

/**
 * Update user password.
 *
 * @param {string} bodyStr - the string containing the new password
 * @return {Promise} the result of the PUT request
 */
export function updatePassword (bodyStr = '') {
  return put({ url: '/api/v1/user/password', bodyStr, statusHandler: HttpStatusHandler() })
}

const ajaj = Ajaj()
const h = AuthHeaders().get()
/**
 * Get user data based on the arguments provided.
 *
 * @param {object} args - The arguments object containing headers and other data.
 * @return {object} The response data from the API.
 */
function getUserData (args = {}) {
  args.headers = h
  return (h.get('authorization').includes('undefined')) ? ajaj.post() : ajaj.post(args)
}
/**
 * Perform a POST request with the given arguments.
 *
 * @param {Object} args - the arguments for the POST request
 * @return {Promise} a Promise representing the POST request
 */
function post (args = {}) {
  args.headers = h; return ajaj.post(args)
}
/**
 * A function that sends a PUT request with the provided arguments.
 *
 * @param {object} args - The arguments for the PUT request.
 * @return {Promise} A Promise representing the PUT request.
 */
function put (args = {}) {
  args.headers = h; return ajaj.put(args)
}
