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
            <input type="text" class="form-control" name="${id}" id="${id}">
            <label for="${id}">${name}</label>
        </div>
    `
}

/**
 * Render input items to form
 *
 * @param {string} [selector=""] selector of detail list
 * @return {{
 * getForm:(object)=>{return string}
 * }}
 */
export default function FormRenderer () {
  return {
    getForm: (details = {}) => {
      let temp = ''
      const keys = Object.keys(details)
      for (const i in keys) {
        const k = keys[i]
        temp += FloatingFormHTML(details[k], k)
      }
      return temp
    }
  }
}
