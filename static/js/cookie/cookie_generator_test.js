import CookieGenerator from './cookie_generator.js'
import UnitTest from '../util/unit_test.js'

function CookieGeneratorTest () {
  const id = 'test'
  const cookieGen = CookieGenerator(id)
  return {
    TestSet: (name = '') => {
      const u = UnitTest(name)
      cookieGen.set('CookieGenerator unit test')
      u.assertNotTrue(cookieGen.get() === undefined)
    },
    TestGet: (name = '') => {
      const u = UnitTest(name)
      u.assertNotTrue(cookieGen.get() === undefined)
      console.log(cookieGen.get())
    },
    TestGetAll: (name = '') => {
      const u = UnitTest(name)
      u.assertNotTrue(cookieGen.getAll() === undefined)
      console.log(cookieGen.getAll())
    },
    TestDelete: (name = '') => {
      const u = UnitTest(name)
      cookieGen.delete()
      u.assertTrue(cookieGen.get() === undefined)
    }
  }
}
CookieGeneratorTest().TestSet('TestSet')
CookieGeneratorTest().TestGet('TestGet')
CookieGeneratorTest().TestGetAll('TestGetAll')
CookieGeneratorTest().TestDelete('TestDelete')
