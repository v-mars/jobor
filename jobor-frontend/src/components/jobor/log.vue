<template>
  <div>
    <div style="margin: 5px 0">
      <!-- 搜索 -->
      <el-form :inline="true" :model="searchForm" size="small" class="demo-form-inline" @submit.native.prevent>
        <el-form-item label="">
          <el-input v-model="searchForm.name" placeholder="名称" @keyup.enter.native="getData"></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-select v-model="searchForm.lang" clearable placeholder="任务类型">
            <el-option value="shell" label="Shell"></el-option>
            <el-option value="python3" label="Python3"></el-option>
            <el-option value="api" label="Api"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="">
          <el-select v-model="searchForm.trigger" clearable placeholder="触发方式">
            <el-option value="auto" label="自动"></el-option>
            <el-option value="manual" label="手动"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="searchForm.worker" placeholder="运行地址" @keyup.enter.native="getData"></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-select v-model="searchForm.status" clearable placeholder="状态">
            <el-option value="running" label="运行"></el-option>
            <el-option value="success" label="成功"></el-option>
            <el-option value="timeout" label="超时"></el-option>
            <el-option value="abort" label="终止"></el-option>
            <el-option value="wait" label="等待"></el-option>
            <el-option value="unknown" label="未知"></el-option>
          </el-select>
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
            <div class="task-log-item">
              <label>执行内容:</label>
              <pre style="white-space: pre-wrap;word-wrap: break-word;margin: 1px;">{{ props.row.data.data }}</pre>
            </div>
            <div class="task-log-item">
              <label>返回结果:</label>
              <pre style="white-space: pre-wrap;word-wrap: break-word;margin: 1px;">{{ props.row.resp }}</pre>
            </div>
            <div class="task-log-item">
              <label>错误状态码:</label> <span v-if="props.row.result!=='success'">{{ props.row.err_code }}</span>
            </div>
            <div class="task-log-item">
              <label>错误信息:</label>
              <pre style="white-space: pre-wrap;word-wrap: break-word;margin: 1px;">{{ props.row.err_msg }}</pre>
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
        <el-table-column label="worker" prop="addr" width=""></el-table-column>
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
            <div v-else-if="scope.row.result==='running'">
              <Badge status="processing" text="运行中" />
            </div>
            <div v-else-if="scope.row.result==='timeout'">
              <Badge status="error" text="超时" />
            </div>
            <div v-else-if="scope.row.result==='abort'">
              <Badge status="warning" text="终止" />
            </div>
            <div v-else>
              <Badge status="default" :text="scope.row.result" />
            </div>
          </template>
        </el-table-column>
        <el-table-column label="更新时间" prop="mtime" width="140"></el-table-column>
        <el-table-column label="操作" align="center" width="120">
          <template slot-scope="scope">
            <el-popconfirm icon="el-icon-info" :title="'确认开始手动执行任务吗？'" @onConfirm="abortTask(scope.row)">
              <delete_button title="终止" slot="reference"
                             v-if="['wait','running'].indexOf(scope.row.result)!==-1"></delete_button>
            </el-popconfirm>

            <detail_button title="详细" style="margin-left:10px"></detail_button>
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
        searchForm: {name: "",lang:"",worker:"",trigger:"",status:""},
      }},
      methods: {

        abortTask: async function (row) {
          try {
            let apiUrl = `${this.url}/${row.id}/abort`
            const response = await this.$store.dispatch("common/Request",
              {url: apiUrl,method:"POST", data: {}});
            this.$message({message:response.data.message, type: "success"})
            await this.getData()
          } catch (e) {
            this.$message({message:String(e), type: "error"})
          } finally {
          }
        },

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

<style lang="scss" scoped>
  .task-log-item{
    margin-bottom: 10px;
  }
</style>
