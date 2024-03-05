import * as validation from '../../js/util/validation.js'
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
    /**
     * Listens for input on formInput and toggles 'is-invalid' class based on password confirmation.
     *
     * @param {Element} formInput - the form input element
     * @param {Element} input - the input element
     * @return {void}
     */
    passwordNotConfirmed: (formInput = Element, input = Element) => {
      formInput.addEventListener('input', (e) => e.target.classList.toggle(
        'is-invalid',
        validation.isPasswordNotConfirmed(input.value, e.target.value)
      ))
    },
    /**
     * Validates the password input in the form.
     *
     * @param {Element} formInput - the input element in the form
     * @return {void} no return value
     */
    passwordValidate: (formInput = Element) => {
      formInput.addEventListener('input', (e) => {
        e.target.classList.toggle(
          'is-invalid',
          validation.isPasswordInvalid(e.target.value)
        )
      })
    },
    /**
     * Validates the email input in the form.
     *
     * @param {Element} formInput - the input element to validate
     * @return {void}
     */
    emailValidate: (formInput = Element) => {
      formInput.addEventListener('input', (e) => e.target.classList.toggle(
        'is-invalid',
        validation.isEmailInvalid(e.target.value)
      ))
    },
    /**
     * Check if the provided form data has any invalid values.
     *
     * @param {object} data - the form data to be validated
     * @return {boolean} true if the form data has invalid values, false otherwise
     */
    isFormDataHasInvaild: (data = {}) => {
      return validation.isPasswordNotConfirmed(data.Password, data.ConfirmPw) ||
  validation.isEmailInvalid(data.Email) ||
  validation.isPasswordInvalid(data.Password)
    }
  }
}
