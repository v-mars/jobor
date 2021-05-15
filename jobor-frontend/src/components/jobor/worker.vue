<template>
  <div>
    <div style="margin: 5px 0">
      <!-- 搜索 -->
      <el-form :inline="true" :model="searchForm" size="small" class="demo-form-inline" @submit.native.prevent>
        <el-form-item label="">
          <el-input v-model="searchForm.hostname" placeholder="主机名" @keyup.enter.native="getData"></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="searchForm.addr" placeholder="地址" @keyup.enter.native="getData"></el-input>
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
        <el-table-column label="ID" prop="id" width="60"></el-table-column>
        <el-table-column label="主机名称" prop="hostname" width="150"></el-table-column>
        <el-table-column label="地址" prop="addr" width="160"></el-table-column>
<!--        <el-table-column label="表达式" prop="expr" width=""></el-table-column>-->
        <el-table-column label="routingKey" prop="routing_key" width=""></el-table-column>
        <el-table-column label="权重" prop="weight" width="100"></el-table-column>
        <el-table-column label="状态" prop="" width="100" align="">
          <template slot-scope="scope">
            <div style="vertical-align: middle;display:flex;flex-direction:row;align-items:center;" v-if="scope.row.status==='running'">
              <Badge status="success" text="运行中" />
            </div>
            <div v-else-if="scope.row.status==='stop'">
              <Badge status="warning" text="停止" />
            </div>
            <div v-else-if="scope.row.status==='offline'">
              <Badge status="error" text="离线" />
            </div>
            <div v-else>
              <Badge status="default" :text="scope.row.result" />
            </div>
          </template>
        </el-table-column>
        <el-table-column label="更新时间" prop="mtime" width="140"></el-table-column>
        <el-table-column label="操作" align="center" width="120">
          <template slot-scope="scope">
            <el-popconfirm icon="el-icon-info" :title="'确认启用此节点吗？'"
                           v-if="['stop'].indexOf(scope.row.status)!==-1"
                           @onConfirm="changeWorkerStatus(scope.row, 'running')">
              <green_button title="启用" slot="reference"></green_button>
            </el-popconfirm>

            <el-popconfirm icon="el-icon-info" :title="'确认禁用此节点吗？'"
                           v-if="['running'].indexOf(scope.row.status)!==-1"
                           @onConfirm="changeWorkerStatus(scope.row, 'stop')">
              <edit_button title="禁用" slot="reference"></edit_button>
            </el-popconfirm>
<!--            <delete_button title="删除" @click="changeWorkerStatus(scope.row)"></delete_button>-->
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
      name: "jobor_task_worker",
      mixins: [common_mixin],
      data(){return{
        url: this.$store.state.urls.jobor_worker_url,
        searchForm: {addr: "",hostname:""},
      }},
      methods: {
        changeWorkerStatus: async function (row, action) {
          try {
            let data = {status: action}
            let apiUrl = `${this.url}/${row.id}`
            const response = await this.$store.dispatch("common/Request",
              {url: apiUrl, method: "PUT", data: data});
            this.$message({message: response.data.message, type: "success"})
            await this.getData()
          } catch (e) {
            this.$message({message: String(e), type: "error"})
          } finally {
          }
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
