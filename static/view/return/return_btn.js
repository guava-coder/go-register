const sceneTemp = document.querySelector('#sceneTemp')
const returnBtn = document.querySelector('#returnBtn')
try {
  returnBtn.setAttribute('hx-get', sceneTemp.innerHTML)
  htmx.process(returnBtn)
} catch (err) {
  console.log(err)
}
