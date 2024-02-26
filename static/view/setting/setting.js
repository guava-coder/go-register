import UserController from '../../js/controller/user_controller.js'

/**
 *
 * @return {{
 * Name:Element,
 * Email:Element,
 * Bio:Element,
 * }}
 */
const BasicInfoInputs = () => {
  return {
    Name: document.querySelector('#name'),
    Email: document.querySelector('#email'),
    Bio: document.querySelector('#bio')
  }
}

UserController().findUserData()
  .catch(err => console.log(err))
  .then(data => {
    const user = data.User
    const inputs = BasicInfoInputs()
    for (const k of Object.keys(inputs)) {
      inputs[k].value = user[k]
    }
  })

document.querySelector('#userform').addEventListener('submit', function (e) {
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
})
