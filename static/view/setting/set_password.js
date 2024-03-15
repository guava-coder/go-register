import UserController from '../../js/controller/user_controller.js'

(() => {
  document.querySelector('#passwordForm').addEventListener('submit', (e) => {
    e.preventDefault()
    const formData = new FormData(e.target)

    let flag = true
    formData.forEach(i => {
      if (i === '') {
        flag = false
      }
    })
    if (flag && document.querySelector('.is-invalid') === null) {
      const userData = Object.fromEntries(formData)

      delete userData.ConfirmPw

      UserController().updatePassword(JSON.stringify(userData))
        .catch(err => {
          alert(err.Response)
          document.querySelector('#confirmpw').value = ''
        })
        .then(data => {
          if (data !== undefined) {
            const conf = confirm(data.Response)
            if (conf) {
              location.reload()
            } else {
              location.reload()
            }
          }
        })
    } else {
      document.querySelector('#confirmpw').value = ''
      alert('form has invalid input')
    }
  })
})()
