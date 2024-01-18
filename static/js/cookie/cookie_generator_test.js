import CookieGenerator from './cookie_generator.js'
import UnitTest from '../util/unit_test.js'

function CookieGeneratorTest () {
  const id = 'test'
  const cookieGen = CookieGenerator(id)
  return {
    testSet: (name = '') => {
      const u = UnitTest(name)
      cookieGen.set('CookieGenerator unit test')
      u.assertNotTrue(cookieGen.get() === undefined)
    },
    testGet: (name = '') => {
      const u = UnitTest(name)
      u.assertNotTrue(cookieGen.get() === undefined)
      console.log(cookieGen.get())
    },
    testGetAll: (name = '') => {
      const u = UnitTest(name)
      u.assertNotTrue(cookieGen.getAll() === undefined)
      console.log(cookieGen.getAll())
    },
    testDelete: (name = '') => {
      const u = UnitTest(name)
      cookieGen.delete()
      u.assertTrue(cookieGen.get() === undefined)
    }
  }
}

CookieGeneratorTest().testSet('TestSet')
CookieGeneratorTest().testGet('TestGet')
CookieGeneratorTest().testGetAll('TestGetAll')
CookieGeneratorTest().testDelete('TestDelete')
