import UserController from '../../js/controller/user_controller.js'

(() => document.querySelector('#userform').addEventListener('submit', function (e) {
  e.preventDefault()

  const formData = new FormData(e.target)
  const userData = Object.fromEntries(formData)
  try {
    const dataStr = JSON.stringify(userData)
    UserController().updateUserInfo(dataStr)
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
