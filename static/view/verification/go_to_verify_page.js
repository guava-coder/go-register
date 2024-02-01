export default function GotoVerifyPage (userStr = '') {
  const content = document.querySelector('.container')
  fetch('/static/view/verification/email_verify.html')
    .then(res => { return res.text() })
    .catch(err => console.log(err))
    .then(data => {
      content.innerHTML = data
      content.innerHTML += `<div id="userdata" hidden>${userStr}</div>`
      htmx.process(content)
    })
}
