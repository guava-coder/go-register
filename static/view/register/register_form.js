import UserController from '../../js/controller/user_controller.js'
import Bootstrap5ShowPassword from '../../js/util/bootstrap5_show_password.js'
import EmailController from '../../js/controller/email_controller.js'
import GotoVerifyPage from '../verification/go_to_verify_page.js'

const addUserAndSendVerificationCode = (body = '') => {
  UserController().addUser(body)
    .catch(err => alert(err.Response))
    .then(data => {
      if (data !== undefined) {
        document.querySelector('#submit').innerHTML = `Verifying email address...
        <i class="bi bi-hourglass-split"></i>`
        EmailController().sendVerificationMail(JSON.stringify(data.User))
          .catch(err => alert(err.Response))
          .then(res => { if (res !== undefined) GotoVerifyPage(data.User.Id) })
      }
    })
}

document.querySelector('#registerForm').addEventListener('submit', function (e) {
  e.preventDefault()
  const formData = new FormData(e.target)
  const userData = Object.fromEntries(formData)
  if (userData.Password === userData.ConfirmPw && userData.Password !== '') {
    addUserAndSendVerificationCode(JSON.stringify(userData))
  } else {
    alert('Confirm Password incorrect')
  }
})

const showPsw = document.querySelector('#showPsw')
showPsw.onclick = () => Bootstrap5ShowPassword(showPsw, '#password')

const showConfirmPsw = document.querySelector('#showConfirmPsw')
showConfirmPsw.onclick = () => Bootstrap5ShowPassword(showConfirmPsw, '#confirmpw')
