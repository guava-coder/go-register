import * as valid from './bootstrap5_validation.js'
import Bootstrap5ShowPassword from './bootstrap5_show_password.js'

const psw = document.querySelector('#password')
valid.passwordValidate(psw)
Bootstrap5ShowPassword(document.querySelector('#showPsw'), '#password')
