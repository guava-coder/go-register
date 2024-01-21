import FormRenderer from './form_renderer.js'
import UnitTest from './unit_test.js'

function FormRendererTest () {
  const renderer = FormRenderer()
  const form = document.querySelector('#testForm')
  return {
    testGetForm: (u = UnitTest()) => {
      const name = 'name'
      const email = 'email'
      form.innerHTML += renderer.getForm({ Name: name, Email: email })
      form.innerHTML += /* html */`
      <div class="text-center">
        <button type="submit" class="btn btn-primary btn-block mb-4" id="submit">submit</button>
    </div>`
      u.assertNotTrue(document.querySelector('#Name') === null && document.querySelector('#Email') === null)
    }
  }
}

FormRendererTest().testGetForm(UnitTest('testRender'))
