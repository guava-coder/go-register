import UserController from '../../js/controller/user_controller.js'
import EmailController from '../../js/controller/email_controller.js'

(() => {
  const userObj = JSON.parse(document.querySelector('#userdata').innerHTML)

  const updateUserAuth = () => {
    const input = { Id: userObj.Id, Auth: document.querySelector('#auth').value }

    UserController().updateUserAuth(JSON.stringify(input))
      .catch(err => console.log(err))
      .then(data => {
        if (data !== undefined) {
          if (confirm(data.Response)) {
            location.reload()
          } else {
            location.reload()
          }
        }
      })
  }
  document.querySelector('#submit').onclick = updateUserAuth

  const resendMail = document.querySelector('#resend')

  resendMail.onclick = () => EmailController().sendVerificationMail(JSON.stringify(userObj))
    .catch(err => console.log(err.Response))
    .then(res => alert(res.Response))
})()
