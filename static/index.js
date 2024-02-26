const saveTemp = () => {
  const sceneTemp = document.querySelector('#sceneTemp')
  const view = document.querySelector('div[hx-get]').getAttribute('hx-get')
  sceneTemp.innerHTML = view
}
saveTemp()
