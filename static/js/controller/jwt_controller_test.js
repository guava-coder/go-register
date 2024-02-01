import UnitTest from '../util/unit_test.js'
import JwtController from './jwt_controller.js'

function JwtControllerTest () {
  const jwtController = JwtController()
  return {
    testLogin: (u = UnitTest()) => {
      jwtController.login('{"Email":"lisa@mail.com","Password":"123"}')
        .catch(err => console.log(err))
        .then(data => u.assertNotTrue(data === undefined))
    }
  }
}

JwtControllerTest().testLogin(UnitTest('testLogin'))
