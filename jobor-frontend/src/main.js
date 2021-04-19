import Vue from 'vue'

import 'normalize.css/normalize.css' // A modern alternative to CSS resets

import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
// import locale from 'element-ui/lib/locale/lang/en' // lang i18n
import iView from 'iview';
import 'iview/dist/styles/iview.css';

import '@/styles/index.scss' // global css

import App from './App'
import store from './store'
import router from './router'

import '@/icons' // icon
import '@/permission' // permission control
import common from './utils/common.js'

import axios from 'axios'
import VueAxios from 'vue-axios'
/**
 * If you don't want to use mock-server
 * you want to use MockJs for mock api
 * you can execute: mockXHR()
 *
 * Currently MockJs will be used in the production environment,
 * please remove it before going online ! ! !
 */

// set ElementUI lang to EN
// Vue.use(ElementUI, { locale })
// 如果想要中文版 element-ui，按如下方式声明
Vue.use(ElementUI)

Vue.use(VueAxios,axios);
Vue.use(iView);

Vue.config.productionTip = false;
Vue.prototype.common = common

new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
})
