import UnitTest from '../util/unit_test.js'
import JwtController from './jwt_controller.js'

function JwtControllerTest () {
  const jwtController = JwtController()
  return {
    testLogin: (u = UnitTest()) => {
      jwtController.login('{"Email":"mark@mail.com","Password":"123"}')
        .then(() => u.assertTrue(true))
    }
  }
}

JwtControllerTest().testLogin(UnitTest('testLogin'))
