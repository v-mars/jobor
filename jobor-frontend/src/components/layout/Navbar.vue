<template>
  <div class="navbar">

    <hamburger
      v-if="showHamburger"
      style="padding: 0 0;"
      :is-active="sidebar.opened"
      class="hamburger-container"
      @toggleClick="toggleSideBar"
    />
    <div v-if="showLogo" style="line-height: 50px;position: relative;float: left">
      <router-link class="sidebar-logo-link" to="">
        <img src="@/assets/logo2-2.png" class="sidebar-logo">
      </router-link>
    </div>

    <div v-if="showTitle" class="title-container" @click="ShowDrawer=true">
      <el-link :underline="false" style="color: #FFFFFF;">
        <!--        <svg-icon icon-class="component" style="width: 14px;height: 14px;" />-->
        <i class="el-icon-menu" />
        <span style="margin-left: 5px">{{ title }}</span>
      </el-link>
    </div>

    <!--    <breadcrumb class="breadcrumb-container" />-->

    <div class="right-menu">
      <el-dropdown class="avatar-container" trigger="click">
        <div class="avatar-wrapper">
          <img src="@/assets/avatar.gif" class="user-avatar">
          <span class="user-avatar" style="margin-left: 5px">{{ $store.state.user.username }}</span>
          <i class="el-icon-caret-bottom" />
        </div>
        <el-dropdown-menu slot="dropdown" class="user-dropdown">
          <router-link to="/">
            <el-dropdown-item>
              首页
            </el-dropdown-item>
          </router-link>
          <el-dropdown-item divided @click.native="logout">
            <span style="display:block;">退出</span>
          </el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </div>

    <portal :show-drawer.sync="ShowDrawer" />
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import Breadcrumb from '@/components/Breadcrumb/index'
import Hamburger from '@/components/Hamburger/index'
import portal from '@/components/layout/portal'

export default {
  components: {
    // eslint-disable-next-line vue/no-unused-components
    Breadcrumb,
    Hamburger,
    portal
  },
  props: {
    // FontColor:{default: "color: #FFFFFF;"},
    // MarginRight:{default: "0"},
    // eslint-disable-next-line vue/require-prop-types
    title: { default: 'Title' },
    // eslint-disable-next-line vue/require-prop-types
    showTitle: { default: true },
    // eslint-disable-next-line vue/require-prop-types
    showHamburger: { default: false },
    // eslint-disable-next-line vue/require-prop-types
    showLogo: { default: true }
  },
  data() {
    return {
      ShowDrawer: false
    }
  },
  computed: {
    ...mapGetters([
      'sidebar',
      // 'avatar',
      'name'
    ])
  },
  methods: {
    toggleSideBar() {
      this.$store.dispatch('app/toggleSideBar')
    },
    async logout() {
      await this.$store.dispatch('user/logout')
      this.$router.push(`/login?redirect=${this.$route.fullPath}`)
    },
    ToWorkspace() {
      this.$router.push('/portal/index')
    }

  }
}
</script>

<style lang="scss" scoped>
.navbar {
  height: 50px;
  overflow: hidden;
  position: relative;
  background: #fff;
  box-shadow: 0 1px 4px rgba(0,21,41,.08);

  .hamburger-container {
    line-height: 46px;
    height: 100%;
    float: left;
    cursor: pointer;
    transition: background .3s;
    -webkit-tap-highlight-color:transparent;

    &:hover {
      background: rgba(0, 0, 0, .025)
    }
  }

  .title-container{
    float: left;
    line-height: 50px;
    height: 100%;
    padding: 0 30px;
  }

  .breadcrumb-container {
    float: left;
  }

  .right-menu {
    float: right;
    height: 100%;
    line-height: 50px;

    &:focus {
      outline: none;
    }

    .right-menu-item {
      display: inline-block;
      padding: 0 8px;
      height: 100%;
      font-size: 18px;
      color: #5a5e66;
      vertical-align: text-bottom;

      &.hover-effect {
        cursor: pointer;
        transition: background .3s;

        &:hover {
          background: rgba(0, 0, 0, .025)
        }
      }
    }

    .avatar-container {
      margin-right: 0;
      /*margin-right: 20px;*/

      .avatar-wrapper {
        /*margin-top: 5px;*/
        position: relative;
        line-height: 46px;
        .user-avatar {
          cursor: pointer;
          width: 40px;
          height: 40px;
          border-radius: 10px;
          vertical-align: middle;
        }

        .el-icon-caret-bottom {
          cursor: pointer;
          /*position: absolute;*/
          //right: -20px;
          //top: 25px;
          font-size: 12px;
          position: relative;
          vertical-align: middle;
        }
      }
    }
  }
  .font-color{
    color: #FFFFFF;
  }
  .avatar-container{
    color: #FFFFFF;
  }
  .sidebar-logo {
    width: 32px;
    height: 32px;
    vertical-align: middle;
    /*margin-right: 12px;*/
  }

}
</style>
