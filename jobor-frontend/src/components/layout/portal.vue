<template>

  <div>

    <el-drawer title="我是标题" :visible.sync="drawer" :with-header="false" direction="ltr"
               :modal-append-to-body="false" size="280px">
      <el-card :body-style="bodyStyle">
        <div slot="header" class="clearfix">
          <i class="el-icon-menu"></i>
          <span>系统平台</span>
          <span style="float: right;cursor: pointer;color: #00A2CA;font-weight: normal;font-size: 12px;" v-if="!opened" @click="opened=true">展开</span>
          <span style="float: right;cursor: pointer;color: #00A2CA;font-weight: normal;font-size: 12px;" v-if="opened" @click="opened=false">收起</span>
        </div>
        <template v-for="(item,index) in $store.state.settings.platforms" :index="index"
                  v-if="hasPermission($store.state.user.roles, item.roles)">
          <div class="platform-menu" @click="toTarget(item.index)" style="cursor: pointer;">
            <el-link :underline="false" style="color: #303133">{{item.name}}</el-link>
            <div style="font-size: 10px;color: #c3c4c9;">
              {{item.description}}
            </div>
          </div>

          <el-divider class="menu-divider" v-if="$store.state.settings.platforms &&
        $store.state.settings.platforms.length-1!==index"></el-divider>
        </template>
      </el-card>



    </el-drawer>

<!--    <el-drawer-->
<!--      title="我是里面的"-->
<!--      :with-header="false"-->
<!--      size="580px"-->
<!--      direction="ltr"-->
<!--      :append-to-body="false"-->
<!--      :modal-append-to-body="false"-->
<!--      :modal="false"-->
<!--      :visible.sync="opened">-->
<!--      <p>_(:зゝ∠)_</p>-->
<!--    </el-drawer>-->
  </div>


</template>

<script>
  import { mapGetters } from 'vuex'
  import { hasPermission } from '@/utils/common'
    export default {
      name: "portal",
      props: {
        ShowDrawer: {default: false}
      },
      data(){return{
        drawer: this.ShowDrawer,
        bodyStyle: {
          "margin": "0 auto",
          "overflow-y": "scroll",
          "overflow-x": "hidden",
          "width": "100%",
          "height": "calc(100vh - 50px)",
          "padding": "20px 30px"
        },
        opened: false,
        innerDrawer: false,
      }},
      computed: {
        ...mapGetters([
          // "platforms",
        ]),
      },
      watch: {
        "drawer": function (val) {
          // console.log("val:", val, "a1:", this.ShowDrawer)
          if(!this.drawer){this.opened=false}
          this.$emit('update:show-drawer', val)
        },
        "ShowDrawer": function (val) {
          this.drawer = val
        },
      },
      methods: {
        toTarget: function (url) {
          this.$router.push(url)
          this.drawer=false
          this.opened=false
        },
        hasPermission,
      },
    }
</script>

<style lang="scss" scoped>
  /deep/ .menu-divider{
    margin: 10px 0;
  }
  /deep/ .clearfix{
    font-size: 16px;
    font-weight: bold;
  }
  /deep/ .platform-menu{
    position: relative;
    box-shadow: 0 1px 2px -2px rgba(0,0,0,.1), 0 3px 6px 0 rgba(0,0,0,.01), 0 5px 12px 4px rgba(0,0,0,.05);
    border-radius: 4px;
    padding: 8px;
    background: #fff;
    width: 100%;
  }
  /*/deep/ .portal-body{*/
  /*  margin: 0 auto;*/
  /*  overflow-y: scroll;*/
  /*  overflow-x: hidden;*/
  /*  width: 100%;*/
  /*  height: calc(100vh - 50px);*/
  /*  padding: 20px 30px;*/
  /*}*/
</style>
