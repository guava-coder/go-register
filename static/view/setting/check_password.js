import * as js from '../../js/index.js'

(() => {
  document.querySelector('#checkPassword').addEventListener('submit', (e) => {
    e.preventDefault()

    const formData = new FormData(e.target)

    let flag = true
    formData.forEach(i => {
      if (i === '') {
        flag = false
      }
    })
    if (flag) {
      const userData = Object.fromEntries(formData)

      js.UserController.checkPassword(JSON.stringify(userData))
        .catch(err => alert(err.Response))
        .then(data => {
          if (data !== undefined) {
            const hxPage = '#setPassword'
            htmx.trigger(hxPage, 'loadPage')
            htmx.process(document.querySelector('#app'))
          }
        })
    }
  })
})()
