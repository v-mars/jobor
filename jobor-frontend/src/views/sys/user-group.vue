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
          <el-button type="primary" @click="NewRow">新增</el-button>
          <el-button style="float: right" @click="getData()">刷新</el-button>
        </el-form-item>

      </el-form>
    </div>
    <div style="margin-top: 10px">
      <el-table border :data="tableData" size="mini" v-loading="loading">
<!--        <el-table-column type="selection" width="45" align="center"></el-table-column>-->
        <el-table-column label="ID" prop="id" width="60" align=""></el-table-column>
        <el-table-column label="用户组名" prop="name"></el-table-column>
        <el-table-column label="描述" prop="description"></el-table-column>
        <el-table-column label="用户" prop="roles">
          <template slot-scope="scope">
            <el-tag size="mini" :plain="true" type="primary" style="margin: 1px 2px"
                    v-for="(item, index) in scope.row.users" :key="index"
                    v-if="scope.row.hidden && index<3 || !scope.row.hidden">{{item.nickname}}</el-tag>
            <el-tag style="margin: 0 2px;cursor: pointer" size="mini" v-if="scope.row.users.length>3 && scope.row.hidden"
                    @click="scope.row.hidden=false">…+{{scope.row.users.length-3}}</el-tag>
            <el-tag style="margin: 0 2px;cursor: pointer" size="mini" type="info"
                    v-if="scope.row.users.length>3 && !scope.row.hidden" @click="scope.row.hidden=true">隐藏</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="角色" prop="roles">
          <template slot-scope="scope">
            <el-tag size="mini" :plain="true" type="primary" style="margin: 1px 2px"
                    v-for="(item, index) in scope.row.roles" :key="index"
                    v-if="scope.row.role_hidden && index<3 || !scope.row.role_hidden">{{item.name}}</el-tag>
            <el-tag style="margin: 0 2px;cursor: pointer" size="mini" v-if="scope.row.roles.length>3 && scope.row.role_hidden"
                    @click="scope.row.role_hidden=false">…+{{scope.row.roles.length-3}}</el-tag>
            <el-tag style="margin: 0 2px;cursor: pointer" size="mini" type="info"
                    v-if="scope.row.roles.length>3 && !scope.row.role_hidden" @click="scope.row.role_hidden=true">隐藏</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="拥有者" prop="owner_name" width="100"></el-table-column>
        <el-table-column label="更新者" prop="by_update" width="100"></el-table-column>
<!--        <el-table-column label="创建时间" prop="ctime"></el-table-column>-->
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
        <el-form-item label="用户组名" prop="name" :rules="[{required:true,message:'请输入名称', trigger: 'blur'}]">
          <el-input v-model="RowData.name"></el-input>
        </el-form-item>
        <el-form-item label="描述">
          <el-input v-model="RowData.description" type="textarea"></el-input>
        </el-form-item>
        <el-form-item label="选择用户">
          <user_select :selectValue.sync="RowData.users" :single_selection="false"></user_select>
        </el-form-item>
        <el-form-item label="选择角色">
          <role_select :selectValue.sync="RowData.roles" :single_selection="false"></role_select>
        </el-form-item>
        <el-form-item label="拥有者" :hidden="!RowData.id">
          <user_select :selectValue.sync="RowData.owner_id"></user_select>
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
  import user_select from '@/components/sys/user_select'
  import role_select from '@/components/sys/role_select'
  import pagination from '@/components/Pagination/pagination'
    export default {
        name: "role",
        data(){
          return {
            roles: [],
            users: [],

            SearchForm: {},
            RowData: {name: "", description: "",users: [], roles:[],id: "",owner_id:""},
            NewOrUpdate: false,
            action_loading: false,
            again_loading: false,
            title: "新增",

            tableData: [],
            total: 0,
            page: 1,
            limit: 10,
            loading: false,
            url: this.$store.state.urls.user_group_url,
          }
        },
        methods: {

          getData: function(){
            let data = {"pageSize": this.limit, "pageNumber": this.page, name: this.SearchForm.name}
            this.loading = true
            this.$store.dispatch("common/Query", {url: this.url,data: data}).then((response) => {
              // console.log('aaa:',response.data);
              if(response.data.code===200){
                for(let i=0;i<response.data.data.list.length;++i)
                {response.data.data.list[i].hidden=true;response.data.data.list[i].role_hidden=true;}
                this.tableData = response.data.data.list
                this.total = response.data.data.total
              }else {this.$message.error(response.data.message)}
              this.loading = false
            }).catch((response) => {
              this.loading = false

            })
          },

          CreateOrUpdate: function (method, again) {
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
            this.RowData.owner_id = row.owner_id
            let roles = []
            for(let i=0;row.roles.length>i;i++){
              roles.push(row.roles[i].id)
            }
            this.RowData.roles = roles
            let users = []
            for(let i=0;row.users.length>i;i++){
              users.push(row.users[i].id)
            }
            this.RowData.users = users
          },

          NewRow: function() {
            this.NewOrUpdate = true
            this.title = '添加'
            this.RowData.id = ""
            this.RowData.name = ""
            this.RowData.description = ""
            this.RowData.owner_id = null
            this.RowData.roles = []
            this.RowData.users = []
          },

          ConfirmDelRows: function(row) {
            this.$confirm('确认删除用户组：'+row.name, '提示', {
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
        },

        mounted () {
          this.getData()
        },
        components: {
          pagination: pagination,
          user_select:user_select,
          role_select:role_select,
        }
    }
</script>

<style scoped>

</style>
