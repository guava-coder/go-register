import EmailController from './email_controller.js'
import UnitTest from '../util/unit_test.js'

function Test () {
  const controller = EmailController()
  return {
    testVerifyEmail: (u = UnitTest()) => {
      controller.verifyEmail('{"Email":"###@mail.com"}')
        .then(data => {
          u.assertNotTrue(data === undefined)
          console.log(data)
        })
    },
    testSendVerificationMail: (u = UnitTest()) => {
      controller.sendVerificationMail('{"Email":"###@mail.com"}')
        .then(data => {
          u.assertNotTrue(data === undefined)
          console.log(data)
        })
    }
  }
}

Test().testSendVerificationMail(UnitTest('testSendVerificationMail'))
Test().testVerifyEmail(UnitTest('testVerifyEmail'))
