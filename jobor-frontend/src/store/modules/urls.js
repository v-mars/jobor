
const state = {
  login_url: '/api/login',
  refresh_token_url: '/api/v1/refresh-token',
  user_info_url: '/api/v1/user-info',
  logout_url: '/api/logout',

  // sys
  user_url: '/api/v1/sys/user',
  password_url: '/api/v1/sys/user-set-password',
  user_group_url: '/api/v1/sys/usergroup',
  role_url: '/api/v1/sys/role',
  permission_url: '/api/v1/sys/permission',
  portal_url: '/api/v1/sys/portal',
  favor_portal_url: '/api/v1/sys/portal/favor',
  ldap_url: '/api/v1/sys/ldap',
  email_url: '/api/v1/sys/email',
  gitlab_url: '/api/v1/sys/gitlab',
  jenkins_url: '/api/v1/sys/jenkins',
  imageRepository_url: '/api/v1/sys/image-repository',
  property_url: '/api/v1/sys/property',
  sys_state_url: '/api/v1/sys/state',

  // jobor
  jobor_task_url: '/api/v1/jobor/task',
  jobor_log_url: '/api/v1/jobor/log'
}

const mutations = {}
const actions = {}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
