import UnitTest from '../util/unit_test.js'
import UserController from './user_controller.js'

function UserControllerTest () {
  const controller = UserController()
  return {
    testFindUserData: (name = '') => {
      const u = UnitTest(name)
      controller.findUserData()
        .then(data => {
          u.assertNotTrue(data === undefined)
          console.log(data)
        })
    }
  }
}

UserControllerTest().testFindUserData('testFindUserData')
