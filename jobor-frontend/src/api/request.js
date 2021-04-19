import axios from 'axios'
import {Message, MessageBox} from "element-ui";
import store from '@/store'
import router from '../router'
import {getToken, getTokenExpires, getTokenType} from '@/utils/user_token'

//异步请求前在header里加入token
axios.interceptors.request.use(
  async config => {
    // console.log("config.header.url:",config.url, localStorage.getItem('token'))
    if (config.url === '/api/login' || config.url === '/api/v1/login' || config.url === '/api/v1/refresh-token') {  //如果是登录和注册操作，则不需要携带header里面的token
    }
    else {
      let token = getToken()
      let token_type = getTokenType()
      let expire_at = getTokenExpires() || 0
      let curr_timestamp = Date.parse(new Date()) / 1000;
      // expire_at + 3600 refresh_token 过期时间
      // console.log("timestamp:",curr_timestamp < (expire_at + 3600), expire_at, curr_timestamp, Number(expire_at) - curr_timestamp )
      if (curr_timestamp < (Number(expire_at) + 3600) && (Number(expire_at) - curr_timestamp) <= 600) {
        await store.dispatch('user/refreshToken').then((response) => {
          config.headers.Authorization = response.data.data.token_type + ' ' + response.data.data.token;
        }).catch((err) => {
          Message({message: String(err) || 'Request Refresh Token Api Error', type: 'error', duration: 5 * 1000})
          return Promise.reject("刷新token失败:" + err);
        })
      }
      else {
        if (token && token_type) {
          config.headers.Authorization = token_type + ' ' + token;
        } else {
          return Promise.reject("没有token");
        }
      }

    }
    return config
  },
  error => {
    // console.log('axios.interceptor request.error:', error)
    return Promise.reject(error);
  });

/**/
// http response 服务器响应拦截器，
// 这里拦截401错误，并重新跳入登页重新获取token
axios.interceptors.response.use(
  response => {
    // console.log('axios.interceptors.response:', response)
    const res = response.data

    // if the custom code is not 20000, it is judged as an error.
    if (res.code !== 200 && ["arraybuffer", "blob"].indexOf(response.config.responseType)===-1) {
      // Message({message: res.message || 'Error', type: 'error', duration: 5 * 1000});
      let Start600x=/^600\d+$/
      // 999: Illegal token; 50012: Other clients logged in; 50014: Token expired;
      if (res.code === 999 || res.code === 50012 || res.code === 50014) {
        // to re-login
        // console.log("redirect:", router.currentRoute.fullPath)
        MessageBox.confirm('你已经退出, 你可以取消留在这个页面或重新登陆', '确认登出', {
          confirmButtonText: '重新登陆',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          store.dispatch('user/resetToken').then(() => {
            // location.reload()
            // console.log("redirect1:", router.currentRoute.fullPath)
            router.replace({
              path: '/login',
              query: {redirect: router.currentRoute.fullPath}//登录成功后跳入浏览的当前页面
            })
          })
        })
      }
      else if (Start600x.test(res.code)){
        console.log("Warning提示:",res.message)
      }
      else {
        Message({ message: res.message || 'Error', type: 'error' });
      }
      return Promise.reject(new Error(res.message || 'Error'))
    } else {
      // console.log("redirect:", router.currentRoute.fullPath)
      return response
    }
  }, error => {
    // console.log('axios.interceptor .error:', error)
    Message({ message: error || 'Error', type: 'error' });
    return Promise.reject(error)
  }
);


export function Get(url, params) {
  // console.log("params:", params)
  return axios({
    url: url,
    method: 'GET',
    params: params,
  })
}

export function Other(url,method, data) {
  return axios({
    url: url,
    method: method,
    data: data
  })
}

export function Request(url,method, data) {
  if(method.toLowerCase()==='get'){
    return axios({
      url: url,
      method: method,
      params: data,
    })
  }else {
    return axios({
      url: url,
      method: method,
      data: data
    })
  }

}

export function Post(url, data) {
  return axios({
    url: url,
    method: 'post',
    data: data
  })
}

export function Put(url, data) {
  return axios({
    url: url,
    method: 'put',
    data: data
  })
}

export function Delete(url, data) {
  return axios({
    url: url,
    method: 'delete',
    data: data
  })
}
