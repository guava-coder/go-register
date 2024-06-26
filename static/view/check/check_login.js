import * as js from '../../js/index.js'

export default function CheckLogin () {
  const fetchHTMLPage = (url = '', handleData = () => {}) => {
    const content = document.querySelector('#content')
    fetch(url)
      .then(res => { return res.text() })
      .catch(err => console.log(err))
      .then(data => {
        content.innerHTML = data
        htmx.process(content)
        handleData()
      })
  }

  js.UserController.findUserData()
    .catch(err => console.log(err))
    .then(user => {
      if (user === undefined) {
        fetchHTMLPage('/static/view/check/default_page.html')
      } else {
        fetchHTMLPage('/static/view/check/login_success.html', () => {
          document.querySelector('#welcome').innerHTML = `<h2>Welcome ${user.User.Name}!</h2>`
        })
      }
    })
}
