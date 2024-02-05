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
 * @param {string} [input='']
 * @param {string} [confirm='']
 * @return {boolean}
 */
const isPasswordNotConfirmed = (input = '', confirm = '') => {
  return input !== confirm || input === ''
}

/**
 * @param {string} [psw='']
 * @return {boolean}
 */
const isPasswordNotValid = (psw = '') => {
  const checkStr = '1234567890QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm'
  const isIncludesUpperLowerNum = () => {
    for (const c of checkStr) {
      if (!psw.includes(c)) {
        return false
      }
    }
    return true
  }

  return psw.length < 8 || isIncludesUpperLowerNum()
}

const psw = document.querySelector('#password')
psw.oninput = () => {
  psw.classList.toggle('is-invalid', isPasswordNotValid(psw.value))
}

const confirmedPsw = document.querySelector('#confirmpw')
confirmedPsw.oninput = () => {
  confirmedPsw.classList.toggle('is-invalid', isPasswordNotConfirmed(psw.value, confirmedPsw.value))
}

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

document.querySelector('#registerForm').addEventListener('submit', function (e) {
  e.preventDefault()

  const formData = new FormData(e.target)
  const userData = Object.fromEntries(formData)

  if (
    isPasswordNotConfirmed(userData.Password, userData.ConfirmPw) ||
  isEmailNotVaild(userData.Email) ||
  isPasswordNotValid(userData.Password)
  ) {
    alert('Please complete the register form')
  } else {
    addUserAndSendVerificationCode(JSON.stringify(userData))
  }
})

Bootstrap5ShowPassword(document.querySelector('#showPsw'), '#password')
Bootstrap5ShowPassword(document.querySelector('#showConfirmPsw'), '#confirmpw')
