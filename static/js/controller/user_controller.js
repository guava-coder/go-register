/* eslint-disable comma-dangle */
import Ajaj from '../request/ajaj.js'
import AuthHeaders from '../request/auth_headers.js'
import HttpStatusHandler from '../request/http_status_handler.js'

/**
 *
 *
 * @export
 * @return {{
 * findUserData:()=>{},
 * addUser:(bodyStr='')=>{},
 * updateUserAuth: (bodyStr = '') =>{},
 * updateUserInfo: (bodyStr='')=>{}
 * }}
 */
export default function UserController () {
  const serv = UserService()

  const jwtHeaderHandler = () => {
    const handler = HttpStatusHandler()
    handler.BadRequest = () => console.log('no jwt')
    handler.Forbidden = () => {
      console.log('jwt verify failed, please login again.')
    }
    return handler
  }

  const addUserHandler = () => {
    const handler = HttpStatusHandler()
    handler.Unauthorized = () => console.log('Email address invaild')
    return handler
  }

  const updateAuthHandler = () => {
    const handler = HttpStatusHandler()
    handler.BadRequest = () => alert('Token incorrect')
    return handler
  }

  const updateInfoHandler = () => {
    const handler = HttpStatusHandler()
    handler.BadRequest = () => console.log('Not a user')
    return handler
  }

  return {
    findUserData: () =>
      serv.post({ url: '/api/v1/user/query', bodyStr: '', statusHandler: jwtHeaderHandler() }),
    addUser: (bodyStr = '') =>
      serv.post({ url: '/api/v1/user/add', bodyStr, statusHandler: addUserHandler() }),
    updateUserAuth: (bodyStr = '') =>
      serv.put({ url: '/api/v1/user/auth', bodyStr, statusHandler: updateAuthHandler() }),
    updateUserInfo: (bodyStr = '') =>
      serv.put({ url: '/api/v1/user/update', bodyStr, statusHandler: updateInfoHandler() })
  }
}

function UserService () {
  const ajaj = Ajaj()
  const h = AuthHeaders().get()
  return {
    post: (args = { }) => {
      args.headers = h
      return (h.get('authorization').includes('undefined')) ? ajaj.post() : ajaj.post(args)
    },
    put: (args = { }) => { args.headers = h; return ajaj.put(args) },
  }
}
