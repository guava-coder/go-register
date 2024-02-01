import HttpStatusHandler from '../request/http_status_handler.js'
import ResponseHandler from '../request/response_handler.js'
import UserToken from '../cookie/user_token.js'

/**
 *
 *
 * @export
 * @return {{
 * login: (bodyStr = '') => {}
 * }}
 */
export default function JwtController () {
  return {
    login: async (bodyStr = '') => {
      const statusHandler = HttpStatusHandler()

      statusHandler.BadRequest = () => alert('email address incorrect')
      statusHandler.Forbidden = () => alert('password incorrect')

      return fetch('/api/v1/jwt/login', {
        method: 'POST',
        body: bodyStr,
        headers: new Headers({
          'Content-Type': 'application/json'
        })
      }).then(res => ResponseHandler().run(res, statusHandler))
        .then(data => UserToken().set(data.Token))
    }
  }
}
