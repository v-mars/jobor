import { Get, Other} from '@/api/request'


const state = {

}

const mutations = {
  update_env_list: (state, env_list) => {
    let flag = false
  },
}
const actions = {
  Query({ commit,state }, data) {
    // console.log("data:", data)
    return new Promise((resolve, reject) => {
      Get(data.url, data.data).then(response => {
        // commit('SET_TOKEN', data.token)
        // console.log("res:", response)
        resolve(response)
      }).catch(error => {
        // console.log("Query err:", error)
        reject(error)
      })
    })
  },

  Request({ commit,state }, data) {
    return new Promise((resolve, reject) => {
      Request(data.url,data.method, data.data).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },

  CreateUpdate({ commit,state }, data) {
    return new Promise((resolve, reject) => {
      Other(data.url,data.method, data.data).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },

  Delete({ commit,state }, data) {
    return new Promise((resolve, reject) => {
      Other(data.url,"DELETE", data.data).then(response => {
        resolve(response)
      }).catch(error => {
        reject(error)
      })
    })
  },

}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
