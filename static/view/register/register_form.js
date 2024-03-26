import UserController from '../../js/controller/user_controller.js'
import EmailController from '../../js/controller/email_controller.js'
import GotoVerifyPage from '../verification/go_to_verify_page.js'

(() => {
  const addUserAndSendVerificationCode = (body = '') => {
    UserController().addUser(body)
      .catch(err => alert(err.Response))
      .then(data => {
        if (data !== undefined) {
          const submit = document.querySelector('#submit')

          submit.innerHTML = `Verifying email address...
        <i class="bi bi-hourglass-split"></i>`

          const userStr = JSON.stringify(data.User)

          EmailController().sendVerificationMail(userStr)
            .catch(err => {
              alert(err.Response)
              submit.innerHTML = 'Register'
            })
            .then(res => { if (res !== undefined) GotoVerifyPage(userStr) })
        }
      })
  }

  document.querySelector('#registerForm').addEventListener('submit', function (e) {
    e.preventDefault()

    const formData = new FormData(e.target)

    let flag = true
    formData.forEach(i => {
      if (i === '') {
        flag = false
      }
    })

    if (flag) {
      if (document.querySelector('.is-invalid') === null) {
        const userData = Object.fromEntries(formData)
        delete userData.ConfirmPw
        addUserAndSendVerificationCode(JSON.stringify(userData))
      } else {
        document.querySelector('#confirmpw').value = ''
        alert('form has invalid input')
      }
    } else {
      document.querySelector('#confirmpw').value = ''
      alert('form not complete')
    }
  })
})()
