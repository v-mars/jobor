
const Prefix = 'jobor'

/* get */
export function getToken() {
  return localStorage.getItem(`${Prefix}_token`)
}

export function getRefreshToken() {
  return localStorage.getItem(`${Prefix}_refresh_token`)
}

export function getTokenType() {
  return localStorage.getItem(`${Prefix}_token_type`)
}

export function getTokenExpires() {
  return localStorage.getItem(`${Prefix}_expires_at`)
}

/* set */
export function setToken(token) {
  return localStorage.setItem(`${Prefix}_token`, token)
}

export function setRefreshToken(token) {
  return localStorage.setItem(`${Prefix}_refresh_token`, token)
}

export function setTokenType(token) {
  return localStorage.setItem(`${Prefix}_token_type`, token)
}

export function setTokenExpires(token) {
  return localStorage.setItem(`${Prefix}_expires_at`, token)
}

/* delete */
export function removeToken() {
  return localStorage.removeItem(`${Prefix}_token`)
}

export function removeRefreshToken() {
  return localStorage.removeItem(`${Prefix}_refresh_token`)
}

export function removeTokenType() {
  return localStorage.removeItem(`${Prefix}_token_type`)
}

export function removeExpires() {
  return localStorage.removeItem(`${Prefix}_expires_at`)
}


