import JwtController from '../../js/controller/jwt_controller.js'

document.querySelector('#loginForm').addEventListener('submit', function (e) {
  e.preventDefault()
  const formData = new FormData(e.target)
  const userData = Object.fromEntries(formData)

  JwtController().login(JSON.stringify(userData))
    .then(data => { if (data !== undefined) location.reload() })
})

const showPsw = document.querySelector('#showPsw')
showPsw.onclick = () => {
  const password = document.querySelector('#password')
  const classes = showPsw.classList
  const ishow = 'bi-eye-fill'
  const ihide = 'bi-eye-slash-fill'
  const showPassword = () => {
    classes.remove(ihide)
    classes.add(ishow)
    password.type = 'text'
  }
  const hidePassword = () => {
    classes.remove(ishow)
    classes.add(ihide)
    password.type = 'password'
  }
  if (classes.contains(ishow)) {
    hidePassword()
  } else if (classes.contains(ihide)) {
    showPassword()
  }
}
