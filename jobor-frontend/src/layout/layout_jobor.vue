<template>
  <el-container>

    <sys-header :title="title" />

    <div class="aside-main">
      <el-container style="width: 100%;height: 100%">
        <el-aside width="" style="position: relative;height: 100%;overflow-x: hidden;overflow-y: scroll;">
          <el-menu
            class="el-menu-vertical-demo"
            style="height: 100%"
            :unique-opened="false"
            :collapse-transition="false"
            :default-active="activeMenu"
            :collapse="false"
          >
<!--            <sidebar-item v-for="route in permission_routers" :key="route.path" :item="route" :base-path="route.path" />-->
<!--            <sidebar-item v-for="route in permission_routers" :key="route.path" :item="route" :base-path="route.path" />-->
            <sidebar-item v-for="route in routers.routers" :key="route.path" :item="route" :base-path="route.path" />
          </el-menu>
        </el-aside>
        <el-main style="width: 100%;padding: 10px">
          <app-main />
        </el-main>
      </el-container>
    </div>
  </el-container>

</template>

<script>
import Hamburger from '@/components/Hamburger'
import Breadcrumb from '@/components/Breadcrumb'
import { mapGetters } from 'vuex'
import { Sidebar, AppMain } from './components'
import Navbar from '@/components/layout/Navbar'
import SidebarItem from '@/layout/components/Sidebar/SidebarItem'
import variables from '@/styles/variables.scss'
import ResizeMixin from '@/layout/mixin/ResizeHandler'
import Watermark from '@/utils/watermark' /** 水印**/
import SysHeader from '@/components/layout/SysHeader'
export default {
  name: 'LayoutJobor',
  mixins: [ResizeMixin],
  data() {
    return {
      // isCollapse: true
      title: 'Jobor定时任务'
    }
  },

  computed: {
    ...mapGetters([
      'sidebar',
      // "avatar",
      'name',
      // 'permission_routers',
      'routers'
    ]),

    // sidebar
    routes() {
      return this.$router.options.routes
    },
    activeMenu() {
      const route = this.$route
      const { meta, path } = route
      // console.log(path, meta.activeMenu)
      // if set path, the sidebar will highlight the path you set
      if (meta.activeMenu) {
        return meta.activeMenu
      }
      return path
    },
    showLogo() {
      return this.$store.state.settings.sidebarLogo
    },
    variables() {
      return variables
    },
    isCollapse() {
      return !this.sidebar.opened
    },

    // layout index
    sidebar() {
      return this.$store.state.app.sidebar
    },
    device() {
      return this.$store.state.app.device
    },
    fixedHeader() {
      return this.$store.state.settings.fixedHeader
    },
    classObj() {
      return {
        hideSidebar: !this.sidebar.opened,
        openSidebar: this.sidebar.opened,
        withoutAnimation: this.sidebar.withoutAnimation,
        mobile: this.device === 'mobile'
      }
    }

  },

  mounted: function() {
    const username = this.$store.state.user.username
    const date = new Date()
    Watermark.set(username + ' ' + date.toLocaleDateString())
  },
  methods: {
    handleOpen(key, keyPath) {
      console.log(key, keyPath)
    },
    handleClose(key, keyPath) {
      console.log(key, keyPath)
    },

    handleClickOutside() {
      this.$store.dispatch('app/closeSideBar', { withoutAnimation: false })
    }
  },
  // eslint-disable-next-line vue/order-in-components
  components: {
    // eslint-disable-next-line vue/no-unused-components
    Hamburger: Hamburger,
    // eslint-disable-next-line vue/no-unused-components
    Breadcrumb: Breadcrumb,
    // eslint-disable-next-line vue/no-unused-components
    Navbar,
    // eslint-disable-next-line vue/no-unused-components
    Sidebar,
    AppMain,
    SidebarItem,
    SysHeader
  }

}
</script>

<style scoped>

</style>
<!--lang="scss" scoped-->
<style lang="scss" scoped>
  .el-menu-vertical-demo:not(.el-menu--collapse) {
    width: 200px;
    /*min-height: 400px;*/
  }
  .breadcrumb-container {
    float: left;
  }

  /*@import "~@/styles/mixin.scss";*/
  /*@import "~@/styles/variables.scss";*/

 /* .app-wrapper {
    @include clearfix;
    position: relative;
    height: 100%;
    width: 100%;
    &.mobile.openSidebar{
      position: fixed;
      top: 0;
    }
  }*/
  .drawer-bg {
    background: #000;
    opacity: 0.3;
    width: 100%;
    top: 0;
    height: 100%;
    position: absolute;
    z-index: 999;
  }

  .mobile .fixed-header {
    width: 100%;
  }

  .fixed-header {
    position: fixed;
    top: 0;
    right: 0;
    left: 0;
    z-index: 9;
    //width: calc(100% - #{$sideBarWidth});
    transition: width 0.28s;
  }

  .aside-main{
    display: flex;
    position: relative;
    margin-top: 50px;
    overflow: auto;
    height: calc(100vh - 63px);
    width: 100%;
    margin-left: 0;
  }
</style>
