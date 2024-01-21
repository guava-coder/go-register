import UserController from '../../js/controller/user_controller.js'

CheckLogin()

function CheckLogin () {
  const fetchHTMLPage = (url = '', handleData = () => {}) => {
    const content = document.querySelector('#content')
    fetch(url)
      .then(res => { return res.text() })
      .then(data => {
        content.innerHTML = data
        htmx.process(content)
        handleData()
      })
  }

  UserController().findUserData().then(user => {
    if (user === undefined) {
      fetchHTMLPage('/static/view/check/default_page.html')
    } else {
      fetchHTMLPage('/static/view/check/login_success.html', () => {
        document.querySelector('#welcome').innerHTML = `<h2>Welcome ${user.User.Name}!</h2>`
      })
    }
  })
}
