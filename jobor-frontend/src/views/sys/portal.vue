<template>
  <div style="margin: 10px">
    <div style="margin: 5px 0">
      <!-- 搜索 -->
      <el-form :inline="true" :model="SearchForm" size="small" class="demo-form-inline" @submit.native.prevent>
        <el-form-item label="">
          <el-input v-model="SearchForm.name" placeholder="名称" @keyup.enter.native="getData"></el-input>
        </el-form-item>
        <el-form-item label="">
          <portal_type_select ref="portalTypeSelect" :selectList.sync="select_types" :single_selection="false"></portal_type_select>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="getData">查询</el-button>
          <el-button type="primary" @click="NewRow">新增</el-button>
          <el-button style="float: right" @click="getData()">刷新</el-button>
        </el-form-item>

      </el-form>
    </div>
    <div style="margin-top: 10px">
      <el-table border :data="tableData" size="mini" v-loading="loading">
        <el-table-column type="selection" width="45" align="center"></el-table-column>
        <el-table-column label="名称" prop="name"></el-table-column>
        <el-table-column label="描述" prop="description" :show-overflow-tooltip="true"></el-table-column>
        <el-table-column label="URL" prop="url"></el-table-column>
        <el-table-column label="图标" prop="icon"></el-table-column>
        <el-table-column label="类型" prop="type"></el-table-column>
        <el-table-column label="更新者" prop="by_update" width="100"></el-table-column>
        <el-table-column label="更新时间" prop="mtime" width="140"></el-table-column>
        <el-table-column label="操作" align="center" width="100">
          <template slot-scope="scope">
            <el-button type="text" style="color: #e6a23c" size="mini" @click="showEdit(scope.row)">编辑</el-button>
            <el-button type="text" style="color: #f56c6c" size="mini" @click="ConfirmDelRows(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div style="margin-top: 5px">
        <div style="display: inline">
<!--          <el-button :plain="true" type="default" size="mini">全 选</el-button>-->
<!--          <el-button :plain="true" type="danger" size="mini">选中删除</el-button>-->
        </div>
        <div class="block" style="float: right;display: inline">
          <pagination :total="total" :page.sync="page" :limit.sync="limit" @pagination="getData()"></pagination>
        </div>
      </div>
    </div>

    <el-dialog :title="title" :visible.sync="NewOrUpdate">
      <el-form label-width="80px" :model="RowData" size="small" ref="RowData">
        <el-form-item label="名称" prop="name" :rules="[{required:true,message:'请输入名称', trigger: 'blur'}]">
          <el-input v-model="RowData.name"></el-input>
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="RowData.description" type="textarea"></el-input>
        </el-form-item>
        <el-form-item label="URL" prop="url" :rules="[
          {required:true,message:'请输入URL', trigger: 'blur'},
          {pattern:/(https?|ftp|http):\/\/[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]/,
          message:'请输入正确URL格式：https://www.example.com', trigger: 'blur'},
          ]">
          <el-input v-model="RowData.url"></el-input>
        </el-form-item>
        <el-form-item label="图标" prop="icon">
          <el-input v-model="RowData.icon"></el-input>
        </el-form-item>
        <el-form-item label="类型" prop="type">
          <portal_type_select ref="portalType" :selectId.sync="RowData.type" :allowCreate="true"></portal_type_select>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
    <el-button @click="NewOrUpdate = false">取 消</el-button>
    <el-button :plain="true" type="warning" @click="CreateOrUpdate('PUT', false)" v-if="RowData.id" :loading="action_loading">更新</el-button>
    <el-button :plain="true" type="warning" @click="CreateOrUpdate('POST', true)" v-if="!RowData.id" :loading="action_loading">保存并继续</el-button>
    <el-button :plain="true" type="primary" @click="CreateOrUpdate('POST', false)" v-if="!RowData.id" :loading="action_loading">保 存</el-button>
  </span>
    </el-dialog>
  </div>

</template>

<script>
  import pagination from '@/components/Pagination/pagination'
  import portal_type_select from '@/components/sys/portal_type_select'
    export default {
        name: "portal",
        data(){
          return {
            SearchForm: {},
            RowData: {name: "", description: "",id: "",url:"",icon:"",type:""},
            NewOrUpdate: false,
            action_loading: false,
            again_loading: false,
            title: "新增",

            select_types: [],
            type_list: [],

            tableData: [],
            total: 0,
            page: 1,
            limit: 10,
            loading: false,
            url: this.$store.state.urls.portal_url,
          }
        },
        methods: {
          getData: function(){
            let data = {"pageSize": this.limit, "pageNumber": this.page,
              name: this.SearchForm.name,types:this.select_types}
            this.loading = true
            this.$store.dispatch("common/Query", {url: this.url,data: data}).then((response) => {
              // console.log('aaa:',response.data);
              this.tableData = response.data.data.list
              this.total = response.data.data.total
              this.loading = false
            }).catch((response) => {
              this.loading = false
            })
          },

          CreateOrUpdate: function (method, again) {
            this.$refs.RowData.validate(valid =>{
              if(valid){
                let params = JSON.parse(JSON.stringify(this.RowData))
                let url = this.url
                if(method==='PUT'){url=url+"/"+params.id}
                if(method==='POST'){delete params.id}
                this.action_loading = true
                this.$store.dispatch("common/CreateUpdate",{url: url,"method": method, "data": params}).then((response) => {
                  // console.log('createOrUpdate:',response.data);
                  this.$message.success(response.data.message)
                  if(again===false){
                    this.NewOrUpdate = false
                  }
                  this.getData()
                  this.$refs.portalTypeSelect.getData()
                  this.action_loading = false
                }).catch((response) => {
                  this.action_loading = false

                })

              } else {this.$message.warning('请输入完整数据')}
            })
          },

          showEdit: function(row) {
            this.NewOrUpdate = true
            this.title = '编辑'
            this.RowData.id = row.id
            this.RowData.name = row.name
            this.RowData.description = row.description
            this.RowData.url = row.url
            this.RowData.icon = row.icon
            this.RowData.type = row.type
          },

          NewRow: function() {
            this.NewOrUpdate = true
            this.title = '添加'
            this.RowData.id = ""
            this.RowData.name = ""
            this.RowData.description = ""
            this.RowData.url = ""
            this.RowData.icon = ""
            this.RowData.type = ""
          },

          ConfirmDelRows: function(row) {
            this.$confirm('确认删除角色：'+row.name, '提示', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'warning'
            }).then(() => {
              this.action_loading = true
              this.$store.dispatch("common/Delete",{url: this.url,"data": {rows: [row.id]}}).then((response) => {
                // console.log('createOrUpdate:',response.data);
                this.$message.success(response.data.message)
                this.getData()
                this.$refs.portalTypeSelect.getData()
                this.action_loading = false
              }).catch((response) => {
                this.action_loading = false

              })
            }).catch(() => {
              this.$message({
                type: 'info',
                message: '已取消删除'
              });
            });
          },
        },

        mounted () {
          this.getData()

        },
        components: {
          pagination: pagination,
          portal_type_select: portal_type_select,
        }
    }
</script>

<style scoped>

</style>
