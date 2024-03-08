import CookieGenerator from '../../../js/cookie/cookie_generator.js'
import playjs from '../../../dependencies/playjs/playjs.js'

((u = playjs()) => {
  CookieGenerator().set('CookieGenerator unit test')
  u.assertNotTrue(CookieGenerator().get() === undefined)
})(playjs('testSet'));

((u = playjs()) => {
  u.assertNotTrue(CookieGenerator().get() === undefined)
  console.log(CookieGenerator().get())
})(playjs('testGet'));

((u = playjs()) => {
  u.assertNotTrue(CookieGenerator().getAll() === undefined)
  console.log(CookieGenerator().getAll())
})(playjs('testGetAll'));

((u = playjs()) => {
  CookieGenerator().delete()
  u.assertTrue(CookieGenerator().get() === undefined)
})(playjs('testDelete'))
