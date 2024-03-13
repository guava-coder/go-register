import EmailController from '../../../js/controller/email_controller.js'
import playjs from '../../../dependencies/playjs/playjs.js'

((u = playjs()) => {
  EmailController().verifyEmail('{"Email":"###@mail.com"}')
    .catch(err => console.log(err))
    .then(data => {
      u.assertNotTrue(data === undefined)
      console.log(data)
    })
})(playjs('testSendVerificationMail'));

((u = playjs()) => {
  EmailController().sendVerificationMail(`{
    "Id":"",
    "Email":"###@mail.com"
  }`)
    .catch(err => console.log(err))
    .then(data => {
      u.assertNotTrue(data === undefined)
      console.log(data)
    })
})(playjs('testVerifyEmail'))
