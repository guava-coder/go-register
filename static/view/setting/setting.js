import UserController from '../../js/controller/user_controller.js'

(() => {
  UserController().findUserData()
    .catch(err => console.log(err))
    .then(data => {
      const user = data.User
      const inputs = {
        Name: document.querySelector('#name'),
        Email: document.querySelector('#email'),
        Bio: document.querySelector('#bio')
      }
      for (const k of Object.keys(inputs)) {
        inputs[k].value = user[k]
      }
    })
})()
