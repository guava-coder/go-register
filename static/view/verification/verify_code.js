import UserController from '../../js/controller/user_controller.js'
import EmailController from '../../js/controller/email_controller.js'

const updateUserAuth = (body = '') => {
  UserController().updateUserAuth(body)
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

const userObj = JSON.parse(document.querySelector('#userdata').innerHTML)
const input = { Id: userObj.Id, Auth: document.querySelector('#auth').value }

document.querySelector('#submit').onclick = () => updateUserAuth(JSON.stringify(input))

const resendMail = document.querySelector('#resend')

resendMail.onclick = () => EmailController().sendVerificationMail(JSON.stringify(userObj))
  .catch(err => console.log(err.Response))
  .then(res => alert(res.Response))
