import JwtController from '../../js/controller/jwt_controller.js'
import Bootstrap5ShowPassword from '../../js/util/bootstrap5_show_password.js'
import GotoVerifyPage from '../verification/go_to_verify_page.js'
import EmailController from '../../js/controller/email_controller.js'

document.querySelector('#loginForm').addEventListener('submit', function (e) {
  e.preventDefault()

  const formData = new FormData(e.target)
  const userData = Object.fromEntries(formData)
  const userStr = JSON.stringify(userData)

  JwtController().login(userStr)
    .catch(err => {
      if (err.Id !== null) {
        const authUser = confirm('This email hasn\'t been verified yet, do you want to verify it now ? ')
        if (authUser) {
          const unauthUser = { Id: err.Id, Email: userData.Email }
          const uStr = JSON.stringify(unauthUser)
          EmailController().sendVerificationMail(uStr)
            .catch(err => console.log(err.Response))
            .then(res => { if (res !== undefined) GotoVerifyPage(uStr) })
        }
      } else {
        console.log(err)
      }
    })
    .then(data => {
      if (data !== undefined) location.reload()
    })
})

const showPsw = document.querySelector('#showPsw')
showPsw.onclick = () => Bootstrap5ShowPassword(showPsw, '#password')
