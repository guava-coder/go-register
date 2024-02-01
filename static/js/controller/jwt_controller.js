import GotoVerifyPage from '../../view/verification/go_to_verify_page.js'
import HttpStatusHandler from '../request/http_status_handler.js'
import ResponseHandler from '../request/response_handler.js'
import UserToken from '../cookie/user_token.js'

export default function JwtController () {
  return {
    login: async (bodyStr = '') => {
      const statusHandler = HttpStatusHandler()

      statusHandler.BadRequest = () => alert('email address incorrect')
      statusHandler.Forbidden = () => alert('password incorrect')
      statusHandler.Unauthorized = () => {
        const authUser = confirm('The user doesn\'t authorize, do you want to authorize now?')
        if (authUser) {
          GotoVerifyPage()
        }
      }

      return fetch('/api/v1/jwt/login', {
        method: 'POST',
        body: bodyStr,
        headers: new Headers({
          'Content-Type': 'application/json'
        })
      }).then(res => ResponseHandler().run(res, statusHandler))
        .catch(err => console.log(err))
        .then(data => UserToken().set(data.Token))
    }
  }
}
