/**
 * This function allow you to print test results in console
 *
 * @export
 * @param {string} [name='']
 * @return {{
 * assertTrue: (boolean) => {},
 * assertNotTrue: (boolean) => {}
 * }}
 */
export default function UnitTest (name = '') {
  const printPass = () => console.log(`${name} %cpass`, 'color:green')
  const printFailed = () => console.log(`${name} %cfailed`, 'color:red')

  const assertTrue = (bool) => {
    if (bool) printPass()
    else printFailed()
  }
  const assertNotTrue = (bool) => {
    if (bool) printFailed()
    else printPass()
  }

  return {
    assertTrue: (bool = true) => assertTrue(bool),
    assertNotTrue: (bool = true) => assertNotTrue(bool)
  }
}
