Login()

function Login () {
  const form = document.querySelector('#loginForm')
  form.addEventListener('submit', function (e) {
    e.preventDefault()
    const formData = new FormData(e.target)
    console.log(Object.fromEntries(formData))
  })
}
