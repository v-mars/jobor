import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'
import layout_jobor from '@/layout/layout_jobor'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */


export const sys_routers = {
  path: '/sys',
  component: layout_jobor,
  redirect: '/sys/user',
  name: '系统管理',
  meta: { title: '系统管理', icon: '', roles: ['devops', "system"] },
  children: [
    {
      path: 'user',
      name: '用户',
      component: () => import('@/views/sys/user'),
      meta: { title: '用户', icon: '', roles: ['devops', "system"] }
    },
    {
      path: 'hostgroup',
      name: '用户组',
      component: () => import('@/views/sys/user-group'),
      meta: { title: '用户组', icon: '', roles: ['devops', "system"] }
    },
    {
      path: 'role',
      name: '角色',
      component: () => import('@/views/sys/role'),
      meta: { title: '角色', icon: '', roles: ['devops', "system"] }
    },
    {
      path: 'permission',
      name: '权限',
      component: () => import('@/views/sys/permission'),
      meta: { title: '权限', icon: '', roles: ['devops', "system"] }
    },
    // {
    //   path: 'portal',
    //   name: 'Portal',
    //   component: () => import('@/views/sys/portal'),
    //   meta: { title: 'Portal', icon: '', roles: ['devops', "system"] }
    // },
    // {
    //   path: 'tree',
    //   name: '树管理',
    //   component: () => import('@/views/fort/tree-node-test'),
    //   meta: { title: '树管理', icon: '', roles: ['devops', "system"] }
    // },
    {
      path: 'state',
      name: '系统状态',
      component: () => import('@/views/sys/sys_state'),
      meta: { title: '系统状态', icon: '', roles: ['devops', "system"] }
    },
    // {
    //   path: 'config',
    //   name: '系统配置',
    //   component: () => import('@/views/sys/config'),
    //   meta: { title: '系统配置', icon: '', roles: ['devops', "system"] }
    // },
  ]
}

export const jobor = [
  // {
  //   path: '/jobor/index',
  //   component: layout_jobor,
  //   redirect: '/jobor/index',
  //   name: '概览',
  //   meta: { title: '持续集成概览', icon: '', roles: ['devops', 'admin'] },
  //   children: [
  //     {
  //       path: '',
  //       name: 'joborDashboard',
  //       component: () => import('@/views/jobor/dashboard'),
  //       meta: { title: '概览', icon: '', roles: ['jobor', 'normal'] }
  //     }
  //   ]
  // },
  {
    path: '/jobor',
    component: layout_jobor,
    redirect: '/jobor/task',
    name: 'disTask',
    meta: { title: '定时任务', icon: '', roles: ['devops', 'normal'] },
    children: [
      {
        path: 'task',
        name: '任务管理',
        component: () => import('@/components/jobor/task'),
        meta: { title: '任务管理', icon: '', roles: ['jobor', 'normal'] }
      },
      {
        path: 'log',
        name: 'joborLog',
        component: () => import('@/components/jobor/log'),
        meta: { title: '执行记录', icon: '', roles: ['jobor', 'normal'] }
      },
      {
        path: 'worker',
        name: 'Worker',
        component: () => import('@/components/jobor/worker'),
        meta: { title: '节点', icon: '', roles: ['jobor','normal'] }
      },
    ]
  },
  sys_routers,
  {
    path: '/jobor/help',
    component: layout_jobor,
    redirect: '/jobor/help',
    name: 'jobor帮助',
    meta: { title: '帮助', icon: '', roles: [] },
    children: [
      {
        path: '',
        name: 'joborHelp',
        component: () => import('@/views/jobor/help'),
        meta: { title: '帮助', icon: '', roles: [] }
      }
    ]
  }
]

export const constantRoutes = [
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },

  {
    path: '/404',
    component: () => import('@/views/404'),
    hidden: true
  },

  {
    path: '/',
    component: layout_jobor,
    redirect: '/',
    children: [{
      path: '',
      name: 'Dashboard',
      component: () => import('@/views/jobor/dashboard'),
      meta: { title: '概览', icon: '' }
    },
    ]
  },

  // 404 page must be placed at the end !!!
  // { path: '*', redirect: '/404', hidden: true }
]

export let asyncRouterMap = [
  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]
// eslint-disable-next-line no-const-assign
asyncRouterMap = asyncRouterMap.concat(jobor)

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
