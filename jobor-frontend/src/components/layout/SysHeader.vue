<template>
  <div :class="{'fixed-header':fixedHeader}">
    <el-header height="50px" style="" class="sys-header">
      <div v-if="device==='mobile'&&sidebar.opened" class="drawer-bg" @click="handleClickOutside" />
      <navbar :title="title" style="width: 100%;background-color: #0f6aa8" />
    </el-header>
  </div>
</template>

<script>
  import Navbar from '@/components/layout/Navbar'
  import { mapGetters } from 'vuex'
    export default {
      name: "sys-header",
      props: {
        title: {default: "default-title"},
        HeaderColor: {default: "#0f6aa8"},
      },
      data() {
        return {
          divStyle: {
            "--color": "red"
          },
          widthVar: "100px"
        };
      },
      computed: {
        ...mapGetters([
          // "sidebar",
          // "avatar",
          // "name",
          // "permission_routers",
          // "routers",
        ]),

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
        },

      },
      methods: {
        handleClickOutside() {
          this.$store.dispatch('app/closeSideBar', { withoutAnimation: false })
        }
      },
      components: {
        Navbar
      },
    }
</script>

<style lang="scss" scoped>

  .mobile .fixed-header {
    width: 100%;
  }
  #app .sidebar-container .svg-icon {
    margin-right: 16px;
  }

  .fixed-header {
    position: fixed;
    top: 0;
    right: 0;
    left: 0;
    z-index: 9;
    //width: calc(100% - #{$sideBarWidth});
    /*transition: width 0.28s;*/
  }
  .sys-header {
    vertical-align: middle;display:flex;flex-direction:row;align-items:center;
    background-color: #0f6aa8;
    /*background-color: var(--color);*/
    /*background-color: red;*/
  }
</style>
