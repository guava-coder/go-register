/**
 *
 *
 * @export
 * @return {{
 * passwordNotConfirmed: (formInput = Element,input = Element) => {},
 * passwordValidate: (formInput = Element) => {},
 * emailValidate: (formInput = Element) => {},
 * isFormDataHasInvaild: (data={})=>boolean
 * }}
 */
export default function Bootstrap5Validation () {
  return {
    passwordNotConfirmed: (formInput = Element, input = Element) => {
      formInput.addEventListener('input', (e) => e.target.classList.toggle(
        'is-invalid',
        isPasswordNotConfirmed(input.value, e.target.value)
      ))
    },
    passwordValidate: (formInput = Element) => {
      formInput.addEventListener('input', (e) => {
        e.target.classList.toggle(
          'is-invalid',
          isPasswordInvalid(e.target.value)
        )
      })
    },
    emailValidate: (formInput = Element) => {
      formInput.addEventListener('input', (e) => e.target.classList.toggle(
        'is-invalid',
        isEmailInvaild(e.target.value)
      ))
    },
    isFormDataHasInvaild: (data = {}) => {
      return isPasswordNotConfirmed(data.Password, data.ConfirmPw) ||
  isEmailInvaild(data.Email) ||
  isPasswordInvalid(data.Password)
    }
  }
}

/**
 * @param {string} [input='']
 * @param {string} [confirm='']
 * @return {boolean}
 */
const isPasswordNotConfirmed = (input = '', confirm = '') => {
  return input !== confirm || input === ''
}

/**
 * @param {string} [psw='']
 * @return {boolean}
 */
const isPasswordInvalid = (psw = '') => {
  const checkStr = '1234567890QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm'
  const isIncludesUpperLowerNum = () => {
    for (const c of checkStr) {
      if (!psw.includes(c)) {
        return false
      }
    }
    return true
  }
  return psw.length < 8 || isIncludesUpperLowerNum()
}

/**
 * @param {string} [value='']
 * @return {boolean}
 */
const isEmailInvaild = (value = '') => {
  return !value.includes('@')
}
