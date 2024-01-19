import JwtController from '../../js/controller/jwt_controller.js'

Login()

function Login () {
  const form = document.querySelector('#loginForm')
  form.addEventListener('submit', function (e) {
    e.preventDefault()
    const formData = new FormData(e.target)
    const userData = Object.fromEntries(formData)

    JwtController().login(JSON.stringify(userData)).then(() => location.reload())
  })
}
