import playjs from '../../../dependencies/playjs/playjs.js'
import AuthHeaders from '../../../js/request/auth_headers.js'

((u = playjs()) => {
  const data = AuthHeaders().get()
  u.assertNotTrue(data === undefined)
  console.log(data)
  console.log(data.get('authorization'))
  console.log(data.get('content-type'))
})(playjs('testGet'))
