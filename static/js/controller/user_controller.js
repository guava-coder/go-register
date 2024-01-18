import AuthHeaders from '../request/auth_headers.js'
import HttpStatusHandler from '../request/http_status_handler.js'
import ResponseHandler from '../request/response_handler.js'

export default function UserController () {
  const serv = UserService()

  const jwtHeaderHandler = (success) => {
    const statusHandler = HttpStatusHandler()
    statusHandler.OK = () => success()
    statusHandler.BadRequest = () => console.log('not jwt')
    statusHandler.Unauthorized = () => console.log('jwt verify failed')
    return statusHandler
  }

  return {
    findUserData: (success = () => {}) => serv.postAjax({ url: '/api/v1/user/query', bodyStr: '', statusHandler: jwtHeaderHandler(success) })
  }
}

function UserService () {
  const RequestArgs = () => {
    return {
      url: '',
      bodyStr: '',
      statusHandler: RequestArgs()
    }
  }

  const ajax = async (met, arg = RequestArgs()) => {
    console.log(AuthHeaders().get())
    return fetch(arg.url, {
      method: met,
      body: arg.bodyStr,
      headers: AuthHeaders().get()
    }).then(res => ResponseHandler().run(res, arg.statusHandler))
      .catch(err => console.log(err))
  }

  return {
    deleteAjax: (args = RequestArgs()) => { return ajax('DELETE', args) },
    putAjax: (args = RequestArgs()) => { return ajax('PUT', args) },
    postAjax: (args = RequestArgs()) => { return ajax('POST', args) }
  }
}
