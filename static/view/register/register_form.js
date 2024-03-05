import UserController from '../../js/controller/user_controller.js'
import * as Bootstrap5Validation from '../../widget/inputs/bootstrap5_validation.js'
import EmailController from '../../js/controller/email_controller.js'
import GotoVerifyPage from '../verification/go_to_verify_page.js'

export default function RegisterForm () {
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

  document.querySelector('#registerForm').addEventListener('submit', function (e) {
    e.preventDefault()

    const formData = new FormData(e.target)
    const userData = Object.fromEntries(formData)

    if (Bootstrap5Validation().isFormDataHasInvaild(userData)) {
      alert('Please complete the register form')
    } else {
      addUserAndSendVerificationCode(JSON.stringify(userData))
    }
  })
}
