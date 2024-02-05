import Bootstrap5ShowPassword from './bootstrap5_show_password.js'
import Bootstrap5Validation from './bootstrap5_vaildation.js'

setTimeout(() => {
  Bootstrap5ShowPassword(document.querySelector('#showPsw'), '#password')
  Bootstrap5ShowPassword(document.querySelector('#showConfirmPsw'), '#confirmpw')

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
      console.log(userData)
    }
  })
}, 150)
