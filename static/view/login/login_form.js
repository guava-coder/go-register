import JwtController from '../../js/controller/jwt_controller.js'
import Bootstrap5ShowPassword from '../../js/util/bootstrap5_show_password.js'
import UserToken from '../../js/cookie/cookie_generator.js'

document.querySelector('#loginForm').addEventListener('submit', function (e) {
  e.preventDefault()

  const formData = new FormData(e.target)
  const userData = Object.fromEntries(formData)

  JwtController().login(JSON.stringify(userData))
    .catch(err => console.log(err))
    .then(data => {
      if (data !== undefined) {
        UserToken().set(data.Token)
        location.reload()
      }
    })
})

const showPsw = document.querySelector('#showPsw')
showPsw.onclick = () => Bootstrap5ShowPassword(showPsw, '#password')
