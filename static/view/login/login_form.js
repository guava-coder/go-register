import FormRenderer from '../../js/util/form_renderer.js'
import JwtController from '../../js/controller/jwt_controller.js'

const form = document.querySelector('#loginForm')
const temp = form.innerHTML
const inputs = FormRenderer().getForm({ Email: 'Email Address', Password: 'Password' })
form.innerHTML = inputs + temp

form.addEventListener('submit', function (e) {
  e.preventDefault()
  const formData = new FormData(e.target)
  const userData = Object.fromEntries(formData)

  JwtController().login(JSON.stringify(userData)).then(() => location.reload())
})
