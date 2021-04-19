<template>
  <div style="">
    <div style="margin: 5px 0">
      <!-- 搜索 -->

      <el-form :inline="true" :model="SearchForm" size="small" class="demo-form-inline" @submit.native.prevent>
        <el-form-item label="">
          <el-input v-model="SearchForm.name" placeholder="名称" @keyup.enter.native="getData"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="getData">查询</el-button>
          <el-button type="primary" @click="newRow">新增</el-button>
          <el-button style="float: right" @click="getData()">刷新</el-button>
        </el-form-item>

      </el-form>
    </div>
    <div style="margin-top: 10px">
      <el-table border :data="data_list" size="mini" v-loading="loading">
<!--        <el-table-column type="selection" width="45" align="center"></el-table-column>-->
        <el-table-column label="ID" prop="id" width="60" align=""></el-table-column>
        <el-table-column label="名称" prop="name"></el-table-column>
        <el-table-column label="描述" prop="description"></el-table-column>
        <el-table-column label="更新者" prop="by_update" width="100"></el-table-column>
<!--        <el-table-column label="创建时间" prop="ctime"></el-table-column>-->
        <el-table-column label="更新时间" prop="mtime" width="140"></el-table-column>
        <el-table-column label="操作" align="center" width="100">
          <template slot-scope="scope">
            <el-button type="text" style="color: #e6a23c" size="mini" @click="showEdit(scope.row)">编辑</el-button>
            <el-button type="text" style="color: #f56c6c" size="mini" @click="confirmDelRows(scope.row)">删除</el-button>
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

    <el-dialog :title="title" :visible.sync="NewOrUpdate" width="80%">
      <el-form label-width="80px" :model="RowData" size="small" ref="RowData">
        <el-form-item label="名称" prop="name" :rules="[{required:true,message:'请输入名称', trigger: 'blur'}]">
          <el-input v-model="RowData.name" :disabled="RowData.id>0"></el-input>
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="RowData.description" type="textarea"></el-input>
        </el-form-item>
        <el-form-item label="权限">
          <permission_select :select-value.sync="RowData.permissions" :single_selection="false"
                             :transfer="false" :collapse_tags="false"></permission_select>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
    <el-button @click="NewOrUpdate = false">取 消</el-button>
    <el-button :plain="true" type="warning" @click="createOrUpdate('PUT', false)" v-if="RowData.id" :loading="action_loading">更新</el-button>
    <el-button :plain="true" type="warning" @click="createOrUpdate('POST', true)" v-if="!RowData.id" :loading="action_loading">保存并继续</el-button>
    <el-button :plain="true" type="primary" @click="createOrUpdate('POST', false)" v-if="!RowData.id" :loading="action_loading">保 存</el-button>
  </span>
    </el-dialog>
  </div>

</template>

<script>
  import pagination from '@/components/Pagination/pagination'
  import permission_select from '@/components/sys/permission_select'
    export default {
        name: "role",
        data(){
          const generateData = _ => {
            const data = [];
            for (let i = 1; i <= 15; i++) {
              data.push({
                key: i,
                label: `备选项 ${ i }`,
                disabled: i % 4 === 0
              });
            }
            return data;
          };
          return {
            SearchForm: {},
            RowData: {name: "", description: "",id: "", permissions: []},
            NewOrUpdate: false,
            action_loading: false,
            again_loading: false,
            title: "新增",

            data_list: [],
            total: 0,
            page: 1,
            limit: 10,
            loading: false,
            url: this.$store.state.urls.role_url,
            permission_url: this.$store.state.urls.permission_url,
            permission_list: [],
          }
        },
        methods: {
          getPermissionData: function(){
            let data = {"pageSize": 0, "pageNumber": this.page}
            this.$store.dispatch("common/Query", {url: this.permission_url,data: data}).then((response) => {
              // console.log('aaa:',response.data);
              let data_list = response.data.data.list
              this.permission_list = []
              for (let i = 0; i < data_list.length; ++i) {
                this.permission_list.push({
                  key: data_list[i].id,
                  label: `${ data_list[i].id }`,
                  disabled: false
                });
              }
            }).catch((response) => {})
          },

          getData: function(){
            let data = {"pageSize": this.limit, "pageNumber": this.page, name: this.SearchForm.name}
            this.loading = true
            this.$store.dispatch("common/Query", {url: this.url,data: data}).then((response) => {
              // console.log('aaa:',response.data);
              this.data_list = response.data.data.list
              this.total = response.data.data.total
              this.loading = false
            }).catch((response) => {
              this.loading = false

            })
          },

          createOrUpdate: function (method, again) {
            this.$refs.RowData.validate(valid =>{
              if(valid){
                let params = JSON.parse(JSON.stringify(this.RowData))
                if(method==='POST'){delete params.id}
                this.action_loading = true
                this.$store.dispatch("common/CreateUpdate",{url: this.url,"method": method, "data": params}).then((response) => {
                  // console.log('createOrUpdate:',response.data);
                  if(response.data.code===200){
                    this.$message.success(response.data.message)
                    if(again===false){
                      this.NewOrUpdate = false
                    }
                  this.getData()
                  }else {this.$message.error(response.data.message)}
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
            this.RowData.permissions = []
            for(let i=0;row.permissions.length>i;++i){this.RowData.permissions.push(row.permissions[i].id)}
          },

          newRow: function() {
            this.NewOrUpdate = true
            this.title = '添加'
            this.RowData.id = ""
            this.RowData.name = ""
            this.RowData.description = ""
          },

          confirmDelRows: function(row) {
            this.$confirm('确认删除角色：'+row.name, '提示', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'warning'
            }).then(() => {
              this.action_loading = true
              this.$store.dispatch("common/Delete",{url: this.url,"data": {rows: [row.id]}}).then((response) => {
                // console.log('createOrUpdate:',response.data);
                if(response.data.code===200){
                  this.$message.success(response.data.message)
                  this.getData()
                }else {this.$message.error(response.data.message)}
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

          filterMethod: function () {

          }
        },

        mounted () {
          this.getData()
        },
        components: {
          pagination: pagination,
          permission_select,
        }
    }
</script>

<style scoped>

</style>
