import Bootstrap5ShowPassword from './bootstrap5_show_password.js'

const showPsw = document.querySelector('#showPsw')
showPsw.onclick = () => Bootstrap5ShowPassword(showPsw, '#password')
