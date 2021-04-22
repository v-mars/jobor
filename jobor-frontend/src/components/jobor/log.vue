<template>
  <div>
    <div style="margin: 5px 0">
      <!-- 搜索 -->
      <el-form :inline="true" :model="searchForm" size="small" class="demo-form-inline" @submit.native.prevent>
        <el-form-item label="">
          <el-input v-model="searchForm.name" placeholder="名称" @keyup.enter.native="getData"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="getData">查询</el-button>
          <el-button style="float: right" @click="getData">刷新</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div style="margin-top: 10px">
      <el-table border :data="data_list" size="small" v-loading="loading">
<!--        <el-table-column type="selection" width="45" align="center"></el-table-column>-->
        <el-table-column label="" type="expand">
          <template slot-scope="props">
            <div>
              <label>执行内容:</label>
              <pre style="white-space: pre-wrap;word-wrap: break-word;">{{ props.row.data.data }}</pre>
            </div>
            <div>
              <label>返回结果:</label>
              <pre style="white-space: pre-wrap;word-wrap: break-word;">{{ props.row.resp }}</pre>
            </div>
            <div>
              <label>err_code:</label>
              <div><span v-if="props.row.result!=='success'">{{ props.row.err_code }}</span></div>
            </div>
            <div>
              <label>错误信息:</label>
              <pre style="white-space: pre-wrap;word-wrap: break-word;">{{ props.row.err_msg }}</pre>
            </div>

          </template>
        </el-table-column>
        <el-table-column label="ID" prop="id" width="60"></el-table-column>
        <el-table-column label="名称" prop="name" width="150"></el-table-column>
        <el-table-column label="类型" prop="lang" width="80"></el-table-column>
        <el-table-column label="触发方式" prop="trigger_method" width="80">
          <template slot-scope="scope">
            <span v-if="scope.row.trigger_method==='auto'">自动</span>
            <span v-else-if="scope.row.trigger_method==='manual'">手动</span>
          </template>
        </el-table-column>
        <el-table-column label="表达式" prop="expr" width=""></el-table-column>
        <el-table-column label="耗时" prop="cost_time" width="100">
          <template slot-scope="scope">{{scope.row.cost_time}}s</template>
        </el-table-column>
        <el-table-column label="结果" prop="" width="100" align="">
          <template slot-scope="scope">
            <div style="vertical-align: middle;display:flex;flex-direction:row;align-items:center;" v-if="scope.row.result==='success'">
              <Badge status="success" text="成功" />
            </div>
            <div v-else-if="scope.row.result==='failed'">
              <Badge status="error" text="失败" />
            </div>
            <div v-else>
              <Badge status="default" :text="scope.row.result" />
            </div>
          </template>
        </el-table-column>
        <el-table-column label="更新时间" prop="mtime" width="140"></el-table-column>
        <el-table-column label="操作" align="center" width="120">
          <template slot-scope="scope">
            <detail_button title="详细"></detail_button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 5px">
        <div style="display: inline"></div>
        <div class="block" style="float: right;display: inline">
          <pagination :total="total" :page.sync="page" :limit.sync="limit" @pagination="getData()"></pagination>
        </div>
      </div>
    </div>


  </div>

</template>

<script>
  import pagination from '@/components/Pagination/pagination'
  import delete_button from '@/components/crud/delete_button'
  import edit_button from '@/components/crud/edit_button'
  import green_button from '@/components/crud/green_button'
  import detail_button from '@/components/crud/detail_button'
  import user_select from '@/components/sys/user_select'
  import common_mixin from '@/components/crud/common_mixin'
    export default {
      name: "jobor_task_log",
      mixins: [common_mixin],
      data(){return{

        url: this.$store.state.urls.jobor_log_url,
      }},
      methods: {


        changeStatus: function (row,status) {
          let statusName="<span style='color: #00c752'>运行</span>"
          if(status==="stop"){statusName="停止"}
          this.$confirm('确认任务名为['+row.name+']执行'+statusName+'动作吗？', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }).then(() => {
            let api_url = this.url+"/"+row.id+"/"+status
            this.$store.dispatch("common/CreateUpdate",{url: api_url,"method": "PUT", "data": {}}).then((response) => {
              this.$message.success(response.data.message)
              this.getData()
              // this.action_loading = false
            }).catch((response) => {
              // this.action_loading = false
            })
          }).catch(() => {this.$message({type: 'info', message: '已取消'})});


        }

      },
      mounted () {
        this.getData()
      },
      components: {
        pagination: pagination,
        delete_button,
        edit_button,
        user_select,
        green_button,detail_button
      },
    }
</script>

<style scoped>

</style>
