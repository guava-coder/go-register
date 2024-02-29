const sceneCurrent = document.querySelector('#sceneCurrent')
function storePrevScene () {
  const scenePrev = document.querySelector('#scenePrev')
  scenePrev.innerHTML = sceneCurrent.innerHTML
}

function setCurrentScene () {
  const view = document.querySelector('div[hx-get]').getAttribute('hx-get')
  sceneCurrent.innerHTML = view
}

setCurrentScene()
storePrevScene()

document.addEventListener('click', e => {
  const view = e.target.getAttribute('hx-get')
  if (view != null) {
    if (!view.includes('scene') || !view.includes('btn')) {
      sceneCurrent.innerHTML = view
    }
  }
})
