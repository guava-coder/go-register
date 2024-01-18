/**
 *
 *
 * @export
 * @param {string} [id='']
 * @return {{
 * set: (string) => {},
 * get: () => {},
 * getAll: () => {},
 * delete: () => {}
 * }}
 */
export default function CookieGenerator (id = '') {
  const cookies = Cookies
  return {
    set: (value = '') => cookies.set(id, value, { expires: 1 }),
    get: () => cookies.get(id),
    getAll: () => cookies.get(),
    delete: () => cookies.remove(id)
  }
}
