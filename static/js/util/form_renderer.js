/**
 * Bootstrap form items
 *
 * @param {string} [name=""] name of input group
 * @param {string} [id=""] input id
 * @return {string} html element string
 */
function FloatingFormHTML (name = '', id = '') {
  return /* html */`
        <div class="form-floating mb-3" id="form${id}">
            <input type="text" class="form-control" name="${id}" id="${id}" disabled>
            <label for="${id}">${name}</label>
        </div>
    `
}

/**
 * Render input items to form
 *
 * @param {string} [selector=""] selector of detail list
 * @return {{
 * render:(object)=>{}
 * }}
 */
export default function FormRenderer (selector = '') {
  const render = (details) => {
    const items = () => {
      let temp = ''
      const keys = Object.keys(details)
      for (const i in keys) {
        const k = keys[i]
        temp += FloatingFormHTML(details[k], k)
      }

      return temp
    }
    const detailDisplay = document.querySelector(selector)
    detailDisplay.innerHTML = items()
  }
  return {
    render: (details = {}) => render(details)
  }
}
