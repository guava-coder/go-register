import UserController from '../../js/controller/user_controller.js'

/**
 *
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
    Userdata().set(JSON.stringify(user))
    const inputs = BasicInfoInputs()
    for (const k of Object.keys(inputs)) {
      inputs[k].value = user[k]
    }
  })

const updateZone = document.querySelector('#updateZone')
const editMode = document.querySelector('#editMode')
editMode.onclick = () => {
  const inputs = BasicInfoInputs()
  for (const k of Object.keys(inputs)) {
    inputs[k].disabled = false
  }
  updateZone.hidden = false
  editMode.hidden = true
}

const cancelBtn = document.querySelector('#cancel')
cancelBtn.onclick = () => {
  const inputs = BasicInfoInputs()
  const user = Userdata().get()
  for (const k of Object.keys(inputs)) {
    inputs[k].value = user[k]
    inputs[k].disabled = true
  }
  updateZone.hidden = true
  editMode.hidden = false
}
