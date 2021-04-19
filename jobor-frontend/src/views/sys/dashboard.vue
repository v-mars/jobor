<template>
    <div style="margin: 20px">
      <el-row :gutter="20" style="margin-bottom: 20px">
        <!--
        <el-col :span="6">
          <el-card style="text-align: center; background-color: rgb(213,118,38)">
            <div style="font-weight: bolder;opacity:0.7;">当前创建版本中</div>
            <div style="font-weight: bold;font-size: 18px">{{current_app_version_count}}</div>
          </el-card>
        </el-col>
        <el-col :span="6">
          <el-card style="text-align: center; background-color: rgb(211,74,58)">
            <div style="font-weight: bolder;opacity:0.7;">当前部署中</div>
            <div style="font-weight: bold;font-size: 18px">{{current_deployment_step_count}}</div>
          </el-card>
        </el-col>
        -->

        <el-col :span="12">
          <el-card style="text-align: center; background-color: rgb(40,157,69)">
            <div style="font-weight: bolder;opacity:0.7;">当前在线用户</div>
            <div style="font-weight: bold;font-size: 18px">{{online_users_count}}</div>
          </el-card>
        </el-col>
        <el-col :span="12">
          <el-card style="text-align: center; background-color: #15a9f0">
            <div style="font-weight: bolder;opacity:0.7;">用户总计</div>
            <div >
              <el-link :underline="true" style="font-weight: bold;font-size: 18px;color: #1f2d3d">{{total_users_count}}</el-link>
            </div>
          </el-card>
        </el-col>
      </el-row>

      <el-row :gutter="20" style="margin-bottom: 20px">
        <el-col :span="24">
          <div style="width: 100%;line-height:38px;height:40px;background:#F5f6FA;border: 1px solid #e1e6eb;border-left: 4px solid #6d7781;border-bottom: none">
            <span style="margin-left: 16px;font-size: 14px;color: #333333">当前在线用户</span>
            <el-input v-model="search" placeholder="请输入搜索内容" style="width: 300px;" size="mini"></el-input>
            <el-button size="mini" type="" @click="manyForceToquit" style="margin: 4px; float: right">批量强制踢人</el-button>
          </div>
          <el-table :data="tables.slice((page-1)*limit,page*limit)" v-loading="loading" ref="table">
            <el-table-column type="selection" width="45" align="center"></el-table-column>
            <el-table-column label="姓名" prop="name" align="" width=""></el-table-column>
            <el-table-column label="用户名" prop="user" align="" width=""></el-table-column>
            <el-table-column label="用户类型" prop="user_type" align="" width="120"></el-table-column>
            <el-table-column label="登陆IP" prop="ip" align="" width="150"></el-table-column>
            <el-table-column label="登陆时间" prop="login_time" align="" width="150"></el-table-column>
            <el-table-column label="sessionId" prop="sessionid" align="" width=""></el-table-column>
            <el-table-column label="操作" prop="" align="center" width="120">
              <template slot-scope="scope">
                <el-button size="mini" type="danger" @click="ForceToQuit([{'sessionid': scope.row.sessionid}])">强制退出</el-button>
              </template>
            </el-table-column>
          </el-table>

          <div class="block">
            <pagination :total="total" :page.sync="page" :limit.sync="limit"></pagination>
          </div>
        </el-col>
        <!--
        <el-col :span="12">
          <div style="width: 100%;line-height:38px;height:40px;background:#F5f6FA;border: 1px solid #e1e6eb;border-left: 4px solid #6d7781;border-bottom: none">
            <span style="margin-left: 16px;font-size: 14px;color: #333333">当前部署中</span>
          </div>
          <el-table :data="current_deployment_steps">
            <el-table-column label="应用" prop="deploy_record__app_version__appCode" align="" :show-overflow-tooltip="true"></el-table-column>
            <el-table-column label="版本" prop="deploy_record__app_version__app_version" :show-overflow-tooltip="true"></el-table-column>
            <el-table-column label="环境" prop="deploy_record__env" width="50" align="center"></el-table-column>
            <el-table-column label="创建者" prop="deploy_record__creator" width="120" align="center"></el-table-column>
            <el-table-column label="创建时间" prop="ctime" align="" :show-overflow-tooltip="true"></el-table-column>
          </el-table>
        </el-col>
        -->
      </el-row>
    </div>
</template>

<script>
  import pagination from '@/components/Pagination/pagination'
  export default {
    name: 'dashboard',
    data(){return{
      search: '',
      total_users_count: 0,
      online_users_count: 0,
      online_users: [],
      limit: 10,
      page: 1,
      total: 0,

      url: this.GLOBAL.service_tag+'/api/system-setting/dashboard/',
      loading: false,
    }},

    created: function(){},
    watch: {
      // 检测表格数据过滤变化，自动跳到第一页
      tables () {
        this.page = 1
        this.total = this.tables.length
      }
    },

    methods: {
      // 获取；
      GetInformation: function() {
        const parms = {}
        this.loading = true
        this.axios({
          method: 'GET',
          url: this.url,
          params: parms,
        }).then((response)=>{
          // console.log('aaa:',response.data);
          if(response.data.status){
            this.total_users_count = response.data.data.total_users
            this.online_users = response.data.data.online_users
            this.online_users_count=this.online_users.length
            this.total = this.online_users_count
          }else {this.$message.error(response.data.error)}
          this.loading = false
        }).catch((response) => {
          this.loading = false

        });
      },

      // 强制退出用户
      ForceToQuit: function(tgt_list) {
        this.$confirm('将强制退出目标, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(() => {
          this.axios({
            method: 'DELETE',
            url: this.url,
            data: tgt_list,
          }).then((response)=>{
            // console.log('aaa:',response.data);
            if(response.data.status){
              this.GetInformation()
              this.$message.success(response.data.message)
            }else {this.$message.error(response.data.error)}
          }).catch((response) => {

          });
          // this.$message({ type: 'success', message: '删除成功!' });
        }).catch(() => {this.$message({ type: 'info', message: '已取消' });});

      },

      // 批量强制退出
      manyForceToquit: function() {
        const _selectData = this.$refs.table.selection
        // console.log('_selectData:', _selectData)
        // return
        if(_selectData.length>0){
          this.ForceToQuit(_selectData)
        }else {
          this.$message.error('请选择目标')
        }
      },

    },

    mounted () {
      this.GetInformation()
      // this.editPermission()

    },
    computed: {
      // 前端过滤
      tables () {
        // console.log("time_range:", this.time_range)
        const search = this.search
        if (search) {
          return this.online_users.filter(dataNews => {
            return Object.keys(dataNews).some(key => {
              return String(dataNews[key]).toLowerCase().indexOf(search) > -1
            })
          })
        }
        return this.online_users
      },

    },
    components: {
      pagination: pagination,
    }

  }
</script>

<style scoped>

</style>
