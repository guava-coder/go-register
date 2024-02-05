import UserController from '../../js/controller/user_controller.js'
import Bootstrap5ShowPassword from '../../js/util/bootstrap5_show_password.js'
import Bootstrap5Validation from '../../js/util/bootstrap5_vaildation.js'
import EmailController from '../../js/controller/email_controller.js'
import GotoVerifyPage from '../verification/go_to_verify_page.js'

setTimeout(() => RegisterForm(), 150)

function RegisterForm () {
  const addUserAndSendVerificationCode = (body = '') => {
    UserController().addUser(body)
      .catch(err => alert(err.Response))
      .then(data => {
        if (data !== undefined) {
          const submit = document.querySelector('#submit')

          submit.innerHTML = `Verifying email address...
        <i class="bi bi-hourglass-split"></i>`

          EmailController().sendVerificationMail(JSON.stringify(data.User))
            .catch(err => {
              alert(err.Response)
              submit.innerHTML = 'Register'
            })
            .then(res => { if (res !== undefined) GotoVerifyPage(data.User) })
        }
      })
  }

  const va = Bootstrap5Validation()

  const psw = document.querySelector('#password')
  va.passwordValidate(psw)

  const email = document.querySelector('#email')
  va.emailValidate(email)

  const confirmedPw = document.querySelector('#confirmpw')
  va.passwordNotConfirmed(confirmedPw, psw)

  document.querySelector('#registerForm').addEventListener('submit', function (e) {
    e.preventDefault()

    const formData = new FormData(e.target)
    const userData = Object.fromEntries(formData)

    if (va.isFormDataHasInvaild(userData)) {
      alert('Please complete the register form')
    } else {
      addUserAndSendVerificationCode(JSON.stringify(userData))
    }
  })

  Bootstrap5ShowPassword(document.querySelector('#showPsw'), '#password')
  Bootstrap5ShowPassword(document.querySelector('#showConfirmPsw'), '#confirmpw')
}
