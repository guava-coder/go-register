import UserController from '../../js/controller/user_controller.js'

CheckLogin()

function CheckLogin () {
  const content = document.querySelector('#content')
  UserController().findUserData().then(user => {
    if (user === undefined) {
      fetch('/static/view/check/default_page.html')
        .then(res => { return res.text() })
        .then(data => {
          content.innerHTML = data
          htmx.process(content)
        })
    } else {
      fetch('/static/view/check/login_success.html')
        .then(res => { return res.text() })
        .then(data => {
          content.innerHTML = data
          htmx.process(content)
          document.querySelector('#welcome').innerHTML = `<h2>Welcome ${user.User.Name}!</h2>`
        })
    }
  })
}
