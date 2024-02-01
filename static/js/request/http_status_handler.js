/* eslint-disable comma-dangle */
/* eslint-disable no-trailing-spaces */
/**
 *Return functions that handle http error during
 *ajax or fetch
 * @return {{
 * OK: () => { },
 * BadRequest: () => { },
 * Unauthorized: () => { },
 * Forbidden: () => {},
 * }}
 * OK() 200, 
 * BadRequest() 400, 
 * Unauthorized() 401, 
 * Forbidden() 403,
 */
export default function HttpStatusHandler () {
  return {
    OK: () => { },
    BadRequest: () => { },
    Unauthorized: () => { },
    Forbidden: () => {},
  }
}
