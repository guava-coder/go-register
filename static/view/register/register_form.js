import Bootstrap5ShowPassword from '../../js/util/bootstrap5_show_password.js'

const gotoVerifyPage = () => {
  const content = document.querySelector('.container')
  fetch('/static/view/register/register_verify.html')
    .then(res => { return res.text() })
    .catch(err => console.log(err))
    .then(data => {
      content.innerHTML = data
      htmx.process(content)
    })
}

document.querySelector('#registerForm').addEventListener('submit', function (e) {
  e.preventDefault()
  const formData = new FormData(e.target)
  const userData = Object.fromEntries(formData)
  if (userData.Password === userData.ConfirmPw && userData.Password !== '') {
    console.log(userData)
    gotoVerifyPage()
  } else {
    alert('Confirm Password incorrect')
  }
})

const showPsw = document.querySelector('#showPsw')
showPsw.onclick = () => Bootstrap5ShowPassword(showPsw, '#password')

const showConfirmPsw = document.querySelector('#showConfirmPsw')
showConfirmPsw.onclick = () => Bootstrap5ShowPassword(showConfirmPsw, '#confirmpw')