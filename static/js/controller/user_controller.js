import Ajaj from '../request/ajaj.js'
import AuthHeaders from '../request/auth_headers.js'
import HttpStatusHandler from '../request/http_status_handler.js'

export default function UserController () {
  const serv = UserService()

  const jwtHeaderHandler = () => {
    const statusHandler = HttpStatusHandler()
    statusHandler.BadRequest = () => console.log('not jwt')
    statusHandler.Unauthorized = () => console.log('jwt verify failed')
    return statusHandler
  }

  return {
    findUserData: () => serv.post({ url: '/api/v1/user/query', bodyStr: '', statusHandler: jwtHeaderHandler() })
  }
}

function UserService () {
  const ajaj = Ajaj()
  const h = AuthHeaders().get()
  return {
    post: (args = { }) => { args.headers = h; return ajaj.post(args) }
  }
}
