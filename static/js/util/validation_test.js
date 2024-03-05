import UnitTest from './unit_test.js'
import * as validation from './validation.js'

function ValidationTest () {
  return {
    testIsPasswordValid: (u = UnitTest()) => {
      u.assertNotTrue(validation.isPasswordInvalid('erisafEA1'))
    },
    testIsEmailValid: (u = UnitTest()) => {
      u.assertNotTrue(validation.isEmailInvalid('erisafEA1@mail.com'))
    },
    testIsPasswordConfirm: (u = UnitTest()) => {
      u.assertNotTrue(validation.isPasswordNotConfirmed('erisafEA1', 'erisafEA1'))
    }
  }
}

ValidationTest().testIsPasswordValid(UnitTest('testIsPasswordValid'))
ValidationTest().testIsEmailValid(UnitTest('testIsEmailValid'))
ValidationTest().testIsPasswordConfirm(UnitTest('testIsPasswordConfirm'))
