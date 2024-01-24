import UnitTest from '../util/unit_test.js'
import UserController from './user_controller.js'

function UserControllerTest () {
  const controller = UserController()
  return {
    testFindUserData: (u = UnitTest()) => {
      controller.findUserData()
        .catch(err => console.log(err))
        .then(data => {
          u.assertNotTrue(data === undefined)
          console.log(data)
        })
    },
    testAddUser: (u = UnitTest()) => {
      controller.addUser(`{
    "Id": "",
    "Name": "eric",
    "Email": "ericwangcatch@gmail.com",
    "Bio":"test",
    "Password": "123",
    "Auth": "QQQQ"}`)
        .catch(err => console.log(err))
        .then(data => {
          u.assertNotTrue(data === undefined)
          console.log(data)
        })
    },
    testUpdateUserAuth: (u = UnitTest()) => {
      controller.updateUserAuth(`{
    "Id": "e2c18694-a181-42f8-8860-9209b9e5a40c",
    "Auth": "9S77BV"
  }`)
        .catch(err => console.log(err))
        .then(data => {
          u.assertNotTrue(data === undefined)
          console.log(data)
        })
    }
  }
}

UserControllerTest().testFindUserData(UnitTest('testFindUserData'))
UserControllerTest().testAddUser(UnitTest('testAddUser'))
UserControllerTest().testUpdateUserAuth(UnitTest('testUpdateUserAuth'))
