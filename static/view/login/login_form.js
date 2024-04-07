import GotoVerifyPage from '../verification/go_to_verify_page.js'
import * as js from '../../js/index.js'

(() => {
  document.querySelector('#loginForm').addEventListener('submit', function (e) {
    e.preventDefault()

    const formData = new FormData(e.target)
    const userData = Object.fromEntries(formData)

    const handleUnauthorized = (err) => {
      if (userData.Email !== '') {
        const authUser = confirm('This email hasn\'t been verified yet, do you want to verify it now ? ')
        if (authUser) {
          const unauthUser = { Id: err.Id, Email: userData.Email }
          const uStr = JSON.stringify(unauthUser)
          js.EmailController.sendVerificationMail(uStr)
            .catch(err => console.log(err.Response))
            .then(res => { if (res !== undefined) GotoVerifyPage(uStr) })
        }
      }
    }

    js.JwtController.login(JSON.stringify(userData))
      .catch(err => {
        console.log(err)
        if (err.Id !== undefined) {
          handleUnauthorized(err)
        }
      })
  })
})()
