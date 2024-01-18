/**
 *Return functions that handle http error during
 *ajax or fetch
 * @return {{
 * OK: () => { },
 * BadRequest: () => { },
 * Unauthorized: () => { }
 * }}
 * OK() 200
 * BadRequest() 400
 * Unauthorized() 401
 */
export default function HttpStatusHandler () {
  return {
    OK: () => { },
    BadRequest: () => { },
    Unauthorized: () => { }
  }
}
