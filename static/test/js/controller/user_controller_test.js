import playjs from '../../../dependencies/playjs/playjs.js'
import * as UserController from '../../../js/controller/user_controller.js'

((u = playjs()) => {
  UserController.findUserData()
    .catch(err => console.log(err))
    .then(data => {
      u.assertNotTrue(data === undefined)
      console.log(data)
    })
})(playjs('testFindUserData'));

((u = playjs()) => {
  UserController.addUser(`
  {
    "Name": "eric",
    "Email": "###",
    "Password": "machiggg7213"
}
    `)
    .catch(err => console.log(err))
    .then(data => {
      u.assertNotTrue(data === undefined)
      console.log(data)
    })
})(playjs('testAddUser'));

((u = playjs()) => {
  UserController.updateUserAuth(`{
    "Id": "e2c18694-a181-42f8-8860-9209b9e5a40c",
    "Auth": "9S77BV"
  }`)
    .catch(err => console.log(err))
    .then(data => {
      u.assertNotTrue(data === undefined)
      console.log(data)
    })
})(playjs('testUpdateUserAuth'));

((u = playjs()) => {
  UserController.updateUserInfo(`
      {
    "Name":"Markii",
    "Bio":"Hi! I'm Mark."
  }`).catch(err => console.log(err))
    .then(data => {
      u.assertNotTrue(data === undefined)
      console.log(data)
    })
})(playjs('testUpdateUserInfo'));

((u = playjs()) => {
  UserController.checkPassword(`{
      "Password": "000"
    }`)
    .catch(err => console.log(err))
    .then(data => {
      u.assertNotTrue(data === undefined)
      console.log(data)
    })
})(playjs('testCheckPassword'));

((u = playjs()) => {
  UserController.updatePassword(`{
  "Password": "123"
}`).catch(err => console.log(err))
    .then(data => {
      u.assertNotTrue(data === undefined)
      console.log(data)
    })
})(playjs('testUpdatePassword'))
