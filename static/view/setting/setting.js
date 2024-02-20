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

/**
 *
 * @return {{
 * get:object,
 * set:(str='')=>{}
 * }}
 */
const Userdata = () => {
  const target = document.querySelector('#userdata')
  return {
    get: () => { return JSON.parse(target.innerHTML) },
    set: (str = '') => { target.innerHTML = str }
  }
}

UserController().findUserData()
  .catch(err => console.log(err))
  .then(data => {
    const user = data.User
    Userdata().set(JSON.stringify())

    const inputs = BasicInfoInputs()
    for (const k of Object.keys(inputs)) {
      inputs[k].value = user[k]
    }
  })

document.querySelector('#userform').addEventListener('submit', function (e) {
  e.preventDefault()

  const formData = new FormData(e.target)
  const userData = Object.fromEntries(formData)
  console.log(userData)
})
