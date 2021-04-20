// eslint-disable-next-line no-unused-vars
import { Get, Other, Request } from '@/api/request'
// eslint-disable-next-line no-unused-vars
import router, { resetRouter } from '@/router'
import { setRefreshToken, setToken, setTokenExpires, setTokenType,
  removeExpires, removeRefreshToken, removeToken, removeTokenType, getRefreshToken } from '@/utils/user_token'
import store from '@/store'

const getDefaultState = () => {
  return {
    password_url: '/api/v1/sys/user-set-password',
    // token: getToken(),
    name: '',
    username: '',
    nickname: '',
    avatar: '',
    token_type: '',
    token: '',
    refresh_token: '',
    expires_at: 0,
    roles: []
  }
}

const state = getDefaultState()

const mutations = {
  SET_USER_DATA: (state, data) => {
    state.name = data.nickname
    state.nickname = data.nickname
    state.username = data.username
    state.expires_at = data.expires_at
    state.roles = data.roles || []
    // console.log("SET_USER_DATA:", data, state.roles)
  },

  SET_LOGIN_DATA: (state, data) => {
    state.refresh_token = data.refresh_token
    state.token_type = data.token_type
    setRefreshToken(data.refresh_token)
    setToken(data.token)
    setTokenType(data.token_type)
    setTokenExpires(data.expires_at)
  },

  RESET_STATE: (state) => {
    Object.assign(state, getDefaultState())
    removeToken()
    removeTokenType()
    removeRefreshToken()
    removeExpires()
  },

  SET_TOKEN: (state, token) => {
    state.token = token
    setToken(token)
    // console.log("token:", token)
  },
  SET_NAME: (state, name) => {
    state.name = name
  },
  SET_AVATAR: (state, avatar) => {
    state.avatar = avatar
  }
}

const actions = {
  // user login
  login({ state, commit }, data) {
    return new Promise((resolve, reject) => {
      Other(data.url, data.method, data.data).then(response => {
        // console.log("response:", response)
        commit('SET_USER_DATA', response.data.data)
        commit('SET_TOKEN', response.data.data.token)
        commit('SET_LOGIN_DATA', response.data.data)
        const roles = response.data.data.roles
        store.dispatch('generateRoutes', {roles}).then(r => {})
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },

  // get user info
  getInfo({ commit, state }) {
    return new Promise((resolve, reject) => {
      Get(store.state.urls.user_info_url, {}).then(response => {
        // console.log("user getInfo:", response.data)
        commit('SET_USER_DATA', response.data.data)
        const roles = response.data.data.roles
        store.dispatch('generateRoutes', {roles}).then(r => {})
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },

  // user logout
  logout({ commit, state }) {
    commit('RESET_STATE')
  },

  // 刷新token
  refreshToken({ state, commit }) {
    return new Promise((resolve, reject) => {
      Other(store.state.urls.refresh_token_url, 'POST', { token: getRefreshToken() }).then(response => {
        // console.log("response:", response)
        if (response.data.code === 200 && response.data.data.token) {
          commit('SET_TOKEN', response.data.data.token)
          commit('SET_LOGIN_DATA', response.data.data)
          resolve(response)
        } else {
          resolve(response)
        }
      }).catch(error => {
        // console.log("refresh error:", error)
        reject(error)
      })
    })
  },

  // remove token
  resetToken({ commit }) {
    commit('RESET_STATE')
  },

  setPassword({ commit, state }, data) {
    return new Promise((resolve, reject) => {
      Other(data.url, 'POST', data.data).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

