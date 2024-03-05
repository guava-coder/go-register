import UnitTest from '../util/unit_test.js'
import AuthHeaders from './auth_headers.js'

function AuthHeadersTest () {
  return {
    testGet: (u = UnitTest()) => {
      const data = AuthHeaders().get()
      u.assertNotTrue(data === undefined)
      console.log(data)
      console.log(data.get('authorization'))
      console.log(data.get('content-type'))
    }
  }
}

AuthHeadersTest().testGet(UnitTest('TestGet'))
