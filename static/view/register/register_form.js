import UserController from '../../js/controller/user_controller.js'
import Bootstrap5ShowPassword from '../../js/util/bootstrap5_show_password.js'
import EmailController from '../../js/controller/email_controller.js'
import GotoVerifyPage from '../verification/go_to_verify_page.js'

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

/**
 *
 *
 * @param {string} [input='']
 * @param {string} [confirm='']
 * @return {boolean}
 */
const isPasswordIncorrect = (input = '', confirm = '') => {
  return input !== confirm && input === ''
}

document.querySelector('#registerForm').addEventListener('submit', function (e) {
  e.preventDefault()

  const formData = new FormData(e.target)
  const userData = Object.fromEntries(formData)

  if (isPasswordIncorrect(userData.Password, userData.ConfirmPw)) {
    alert('Confirm Password incorrect')
  } else {
    addUserAndSendVerificationCode(JSON.stringify(userData))
  }
})

Bootstrap5ShowPassword(document.querySelector('#showPsw'), '#password')
Bootstrap5ShowPassword(document.querySelector('#showConfirmPsw'), '#confirmpw')

/**
 * @param {string} [value='']
 * @return {boolean}
 */
const isEmailNotVaild = (value = '') => {
  return !value.includes('@')
}

const email = document.querySelector('#email')
email.oninput = () => {
  email.classList.toggle('is-invalid', isEmailNotVaild(email.value))
}
