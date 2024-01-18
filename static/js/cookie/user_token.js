import CookieGenerator from './cookie_generator.js'

/**
 * Store User bearer token in cookie, for login
 *
 * @export
 * @return {CookieGenerator()}
 */
export default function UserToken () {
  return CookieGenerator('GRUT')
}
