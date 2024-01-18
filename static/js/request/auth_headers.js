import UserToken from '../cookie/user_token.js'

export default function AuthHeaders () {
  return {
    get: () => {
      const token = (UserToken().get() === null)
        ? ''
        : `Bearer ${UserToken().get()}`

      const headers = new Headers({
        'Content-Type': 'application/json',
        Authorization: token
      })
      return headers
    }
  }
}
