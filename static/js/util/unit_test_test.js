import UnitTest from './unit_test.js'

function UnitTestTest () {
  return {
    testAssertTrue: (u = UnitTest()) => u.assertTrue(true),
    testAssertNotTrue: (u = UnitTest()) => u.assertNotTrue(false)
  }
}

const test = UnitTestTest()

test.testAssertTrue(UnitTest('testAssertTrue'))
test.testAssertNotTrue(UnitTest('testAssertNotTrue'))
