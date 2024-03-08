import Bootstrap5ShowPassword from './bootstrap5_show_password.js'

(() => {
  try {
    const show = document.querySelector('#showPsw')
    if (show) { Bootstrap5ShowPassword(show, '#password') }
  } catch (err) {
    console.log(err)
  }
})()
