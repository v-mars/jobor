import Vue from 'vue'
import Vuex from 'vuex'
import getters from './getters'
import app from './modules/app'
import settings from './modules/settings'
import user from './modules/user'
import permission from './modules/permission'
import urls from './modules/urls'
import apiRest from './modules/common'

Vue.use(Vuex)

const store = new Vuex.Store({
  modules: {
    app,
    settings,
    permission,
    urls,
    apiRest,
    user
  },
  getters
})

export default store
