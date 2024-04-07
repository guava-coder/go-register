import * as js from '../../js/index.js'

(() => {
  const userObj = JSON.parse(document.querySelector('#userdata').innerHTML)

  const updateUserAuth = () => {
    const input = { Id: userObj.Id, TempCode: document.querySelector('#tempcode').value }

    console.log(JSON.stringify(input))

    js.UserController.updateUserAuth(JSON.stringify(input))
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

  resendMail.onclick = () => js.EmailController.sendVerificationMail(JSON.stringify(userObj))
    .catch(err => console.log(err.Response))
    .then(res => alert(res.Response))
})()
