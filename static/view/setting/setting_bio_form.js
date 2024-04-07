import * as js from '../../js/index.js'

(() => document.querySelector('#userform').addEventListener('submit', function (e) {
  e.preventDefault()

  const formData = new FormData(e.target)
  const userData = Object.fromEntries(formData)
  try {
    const dataStr = JSON.stringify(userData)
    js.UserController.updateUserInfo(dataStr)
      .catch(err => alert(err.Response))
      .then(data => {
        if (data !== undefined) {
          alert(data.Response)
        }
      })
  } catch (error) {
    alert('Server error. Error', error)
  }
}))()
