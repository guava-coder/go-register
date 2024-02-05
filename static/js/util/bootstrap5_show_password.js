/**
 * handle bootstrap 5 form password visibility.
 * example of show password button:
 * <span><button type="button" id="showPsw"
 * class="btn bi bi-eye-slash-fill"
 * style="font-size: 2rem;"></button></span>
 *
 * @param {showPsw=Element}
 * @param {string} [input='']
 */
export default function Bootstrap5ShowPassword (showPswBtn = Element, inputSelector = '') {
  showPswBtn.addEventListener('click', (e) => handlePasswordVisibility(e.target, inputSelector))
}

function handlePasswordVisibility (showPsw = Element, input = '') {
  const password = document.querySelector(input)
  const classes = showPsw.classList
  const ishow = 'bi-eye-fill'
  const ihide = 'bi-eye-slash-fill'
  const showPassword = () => {
    classes.remove(ihide)
    classes.add(ishow)
    password.type = 'text'
  }
  const hidePassword = () => {
    classes.remove(ishow)
    classes.add(ihide)
    password.type = 'password'
  }

  if (classes.contains(ishow)) {
    hidePassword()
  } else if (classes.contains(ihide)) {
    showPassword()
  }
}
