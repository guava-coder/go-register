import playjs from '../../../dependencies/playjs/playjs.js'
import JwtController from '../../../js/controller/jwt_controller.js'

((u = playjs()) => {
  JwtController().login('{"Email":"lisa@mail.com","Password":"123"}')
    .catch(err => console.log(err))
    .then(data => u.assertNotTrue(data === undefined))
})(playjs('testLogin'))