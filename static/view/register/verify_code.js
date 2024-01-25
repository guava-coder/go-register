import UserController from '../../js/controller/user_controller.js'

const auth = document.querySelector('#auth')
const submit = document.querySelector('#submit')
const id = document.querySelector('#userid')

const updateUserAuth = (body = '') => {
  UserController().updateUserAuth(body)
    .catch(err => console.log(err))
    .then(data => {
      if (data !== undefined) {
        if (confirm(data.Response)) {
          location.reload()
        } else {
          location.reload()
        }
      }
    })
}

submit.onclick = () => {
  const input = `{"Id":"${id.innerHTML}", "Auth":"${auth.value}"}`
  updateUserAuth(input)
}
