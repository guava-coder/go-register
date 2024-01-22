import Ajaj from '../request/ajaj.js'
import HttpStatusHandler from '../request/http_status_handler.js'

/**
 *
 *
 * @export
 * @return {{
 * verifyEmail:(bodyStr = '') => {}
 * }}
 */
export default function EmailController () {
  const getHandler = () => {
    const statusHandler = HttpStatusHandler()
    statusHandler.BadRequest = () => console.log('not jwt')
    return statusHandler
  }
  const serv = EmailService()
  return {
    verifyEmail: (bodyStr = '') => { return serv.post({ url: '/api/v1/email/verify', bodyStr, statusHandler: getHandler() }) }
  }
}

function EmailService () {
  const ajaj = Ajaj()
  const h = new Headers({ 'Content-Type': 'application/json' })
  return {
    post: (args = {}) => { args.headers = h; return ajaj.post(args) }
  }
}
