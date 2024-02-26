const sceneTemp = document.querySelector('#sceneTemp')
const returnBtn = document.querySelector('#returnBtn')
const app = document.querySelector('#app')
try {
  returnBtn.setAttribute('hx-get', sceneTemp.innerHTML)
  htmx.process(returnBtn)
  htmx.process(app)
} catch (err) {
  console.log(err)
}
