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
 * updateUserInfo: (bodyStr='')=>{},
 * checkPassword: (bodyStr = '') =>{},
 * updatePassword: (bodyStr = '') =>{}
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
    handler.BadRequest = () => console.log('User exist')
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

  const checkPwHandler = () => {
    const handler = HttpStatusHandler()
    return handler
  }

  const updatePasswordHandler = () => {
    const handler = HttpStatusHandler()
    return handler
  }

  return {
    findUserData: () =>
      serv.getUserData({ url: '/api/v1/user/query', bodyStr: '', statusHandler: jwtHeaderHandler() }),
    addUser: (bodyStr = '') =>
      serv.post({ url: '/api/v1/user/add', bodyStr, statusHandler: addUserHandler() }),
    updateUserAuth: (bodyStr = '') =>
      serv.put({ url: '/api/v1/user/auth', bodyStr, statusHandler: updateAuthHandler() }),
    updateUserInfo: (bodyStr = '') =>
      serv.put({ url: '/api/v1/user/update', bodyStr, statusHandler: updateInfoHandler() }),
    checkPassword: (bodyStr = '') =>
      serv.getUserData({ url: '/api/v1/user/check/password', bodyStr, statusHandler: checkPwHandler() }),
    updatePassword: (bodyStr = '') =>
      serv.put({ url: '/api/v1/user/password', bodyStr, statusHandler: updatePasswordHandler() }),
  }
}

function UserService () {
  const ajaj = Ajaj()
  const h = AuthHeaders().get()
  return {
    getUserData: (args = {}) => {
      args.headers = h
      return (h.get('authorization').includes('undefined')) ? ajaj.post() : ajaj.post(args)
    },
    post: (args = {}) => { args.headers = h; return ajaj.post(args) },
    put: (args = {}) => { args.headers = h; return ajaj.put(args) },
  }
}
