import FormRenderer from './form_renderer.js'
import UnitTest from './unit_test.js'

function FormRendererTest () {
  const renderer = FormRenderer('#testForm')
  return {
    testRender: (u = UnitTest()) => {
      const name = 'name'
      const email = 'email'
      renderer.render({ Name: name, Email: email })
      u.assertNotTrue(document.querySelector('#Name') === null)
      u.assertNotTrue(document.querySelector('#Email') === null)
    }
  }
}

FormRendererTest().testRender(UnitTest('testRender'))
