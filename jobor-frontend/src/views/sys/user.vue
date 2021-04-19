<template>
  <div style="">
    <div style="margin: 5px 0">
      <!-- 搜索 -->

      <el-form :inline="true" :model="SearchForm" size="small" class="demo-form-inline" @submit.native.prevent>
        <el-form-item label="">
          <el-input v-model="SearchForm.nickname" placeholder="显示名" @keyup.enter.native="getData"></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="SearchForm.username" placeholder="用户名" @keyup.enter.native="getData"></el-input>
        </el-form-item>
        <el-form-item label="">
          <role_select :selectValue.sync="SearchForm.roles" :single_selection="false"></role_select>
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
        <el-table-column label="显示名/用户名" prop="nickname" width="120">
          <template slot-scope="scope">
            <div style="font-size: 12px">{{scope.row.nickname}}</div>
            <el-link :underline="false" type="primary" style="display: block;font-size: 12px">{{scope.row.username}}</el-link>
          </template>
        </el-table-column>
<!--        <el-table-column label="用户名" prop="username" width="120"></el-table-column>-->
        <el-table-column label="用户类型" prop="user_type_id" align="center" width="80">
          <template slot-scope="scope">
            <el-tag size="mini" type="primary" :plain="true" v-if="scope.row.user_type_id==='1'">本地</el-tag>
            <el-tag size="mini" type="primary" :plain="true" v-else-if="scope.row.user_type_id==='2'">LDAP</el-tag>
            <el-tag size="mini" type="primary" :plain="true" v-else>其它</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="邮箱" prop="email"></el-table-column>
        <el-table-column label="角色" prop="roles">
          <template slot-scope="scope">
            <el-tag size="mini" :plain="true" type="primary" style="margin: 1px 2px"
                    v-for="(item, index) in scope.row.roles" :key="index"
                    v-if="scope.row.hidden && index<3 || !scope.row.hidden">{{item.name}}</el-tag>
            <el-tag style="margin: 0 2px;cursor: pointer" size="mini" v-if="scope.row.roles.length>3 && scope.row.hidden"
                    @click="scope.row.hidden=false">…+{{scope.row.roles.length-3}}</el-tag>
            <el-tag style="margin: 0 2px;cursor: pointer" size="mini" type="info"
                    v-if="scope.row.roles.length>3 && !scope.row.hidden" @click="scope.row.hidden=true">隐藏</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="状态" prop="by_update" width="80" align="center">
          <template slot-scope="scope">
            <el-switch v-model="scope.row.status" @change="ChangeStatus(scope.row)"></el-switch>
          </template>
        </el-table-column>
        <el-table-column label="更新人" prop="by_update" width="100"></el-table-column>
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
        <el-form-item label="显示名" prop="nickname" :rules="[{required:true,message:'请输入显示名', trigger: 'blur'}]">
          <el-input v-model="RowData.nickname"></el-input>
        </el-form-item>
        <el-form-item label="用户名" prop="username" :rules="[{required:true,message:'请输入用户名，如：zhansan', trigger: 'blur'}]">
          <el-input v-model="RowData.username"></el-input>
        </el-form-item>
        <el-form-item label="密码" v-if="RowData.user_type_id==='1' && RowData.id">
          <el-button @click="setPassword.id = RowData.id; dialogPassword=true">修改密码</el-button>
        </el-form-item>
        <el-form-item label="密码" prop="password" :rules="[{required:true,message:'请输入密码', trigger: 'blur'}]" v-if="!RowData.id">
          <el-input v-model="RowData.password" :type="passwordType">
            <span class="show-pwd" @click="showPwd" style="cursor: pointer" slot="append">
              <svg-icon :icon-class="passwordType === 'password' ? 'eye' : 'eye-open'" />
            </span>
          </el-input>
        </el-form-item>
        <el-form-item label="确认密码" prop="repassword" :rules="[{required:true,message:'请输入确认密码', trigger: 'blur'},
        {validator: validateCheckPass,trigger:'blur'}]" v-if="!RowData.id">
          <el-input v-model="RowData.repassword" type="password"></el-input>
        </el-form-item>
        <el-form-item label="邮箱" prop="email" :rules="[{required:true,message:'请输入Email', trigger: 'blur'},{
        pattern: /^[a-z0-9]+([._\-]*[a-z0-9])*@([a-z0-9]+[-a-z0-9]*[a-z0-9]+.){1,63}[a-z0-9]+$/,message: '请输入有效的邮箱', trigger: 'blur'
        }]">
          <el-input v-model="RowData.email"></el-input>
        </el-form-item>
        <el-form-item label="手机号" prop="phone" :rules="[{ pattern: /^(13[0-9]|14[5|7]|15[0|1|2|3|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])\d{8}$/,
         message: '手机格式不对', trigger: 'blur' }]">
          <el-input v-model="RowData.phone"></el-input>
        </el-form-item>
        <el-form-item label="用户类型" prop="user_type_id" :rules="[{required:true,message:'请输入选择用户类型', trigger: 'blur'}]">
          <el-select value="" v-model="RowData.user_type_id">
            <el-option value="1" label="本地"></el-option>
            <el-option value="2" label="LDAP"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="选择角色">
          <role_select :selectValue.sync="RowData.roles" :single_selection="false"></role_select>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
        <el-button @click="NewOrUpdate = false">取 消</el-button>
        <el-button :plain="true" type="warning" @click="CreateOrUpdate('PUT', false)" v-if="RowData.id" :loading="action_loading">更新</el-button>
        <el-button :plain="true" type="warning" @click="CreateOrUpdate('POST', true)" v-if="!RowData.id" :loading="action_loading">保存并继续</el-button>
        <el-button :plain="true" type="primary" @click="CreateOrUpdate('POST', false)" v-if="!RowData.id" :loading="action_loading">保 存</el-button>
      </span>
    </el-dialog>

    <!-- 修改密码 -->
    <el-dialog title="修改密码" :visible.sync="dialogPassword" :close-on-click-modal="false">
      <el-form size="mini" label-width="120px" :model="setPassword" ref="setPassword">
        <el-form-item hidden>
          <el-input v-model="setPassword.id"></el-input>
        </el-form-item>
        <el-form-item label="新密码" prop="password" :rules="[{required:true,message:'请输入密码',trigger:'blur'}]">
          <el-input autocomplete="off" style="width: 240px" type="password" v-model="setPassword.password"></el-input>
        </el-form-item>
        <el-form-item label="确认新密码" prop="aginpassword" :rules="[{required:true,message:'请输入确认密码', trigger: 'blur'},
        {validator: CheckPass,trigger:'blur'}]">
          <el-input autocomplete="off" style="width: 240px" type="password" v-model="setPassword.aginpassword"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
          <el-button @click="dialogPassword = false">取 消</el-button>
          <el-button :plain="true" type="primary" @click="changePassword" v-loading="password_loading">确 定</el-button>
        </span>
    </el-dialog>
  </div>

</template>

<script>
  import role_select from '@/components/sys/role_select'
  import pagination from '@/components/Pagination/pagination'
    export default {
        name: "user",
        data(){
          return {
            passwordType: "password",
            dialogPassword: false,

            setPassword: {},
            SearchForm: {nickname: "", username: "",roles:[]},
            RowData: {nickname: "", username: "", email: "", phone: "",password:"",repassword:"",user_type_id: "1",roles:[],id: ""},
            NewOrUpdate: false,
            action_loading: false,
            again_loading: false,
            password_loading: false,
            title: "新增",

            roles: [],

            tableData: [],
            total: 0,
            page: 1,
            limit: 10,
            loading: false,
            url: this.$store.state.urls.user_url,
            password_url: this.$store.state.urls.password_url,

          }
        },
        methods: {
          validateCheckPass: function(rule, value, callback){
            if (value === "") {
              callback(new Error("请再次输入密码"));
            } else if (value !== this.RowData.password) {
              callback(new Error("两次输入密码不一致!"));
            } else {
              callback();
            }
          },

          CheckPass: function(rule, value, callback){
            if (value === "") {
              callback(new Error("请再次输入密码"));
            } else if (value !== this.setPassword.password) {
              callback(new Error("两次输入密码不一致!"));
            } else {
              callback();
            }
          },

          showPwd() {
            if (this.passwordType === 'password') {
              this.passwordType = ''
            } else {
              this.passwordType = 'password'
            }
            // this.$nextTick(() => {
            //   this.$refs.password.focus()
            // })
          },

          changePassword: function(){
            this.$refs.setPassword.validate(valid =>{
              if(valid){
                let params = this.setPassword
                this.password_loading = true
                this.$store.dispatch("user/setPassword",{url: this.password_url,"data": params}).then((response) => {
                  // console.log('createOrUpdate:',response.data);
                  if(response.data.code===200){
                    this.dialogPassword=false
                    this.$message.success(response.data.message)
                  }else {this.$message.error(response.data.message)}
                  this.password_loading = false
                }).catch((response) => {
                  this.password_loading = false

                })

              } else {this.$message.warning('请输入完整数据')}
            })
          },

          getData: function(){
            let data = {"pageSize": this.limit, "pageNumber": this.page, nickname: this.SearchForm.nickname,
              username: this.SearchForm.username,"roles":this.SearchForm.roles}
            this.loading = true
            this.$store.dispatch("common/Query", {url:this.url,data:data}).then((response) => {
              // console.log('aaa:',response.data);
              for(let i=0;i<response.data.data.list.length;++i)
              {response.data.data.list[i].hidden=true}
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
                this.$store.dispatch("common/CreateUpdate",{"url":url,"method": method, "data": params}).then((response) => {
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

          ChangeStatus: function (row) {
            let params = {"id": row.id, "status": row.status}
            let url = this.url+"/"+params.id
            this.$store.dispatch("common/CreateUpdate",
              {"url":url,"method": "PUT", "data": params}).then((response) => {
              this.$message.success(response.data.message)
              this.getData()
            }).catch((response) => {})
          },

          showEdit: function(row) {
            this.NewOrUpdate = true
            this.title = '编辑'
            this.RowData.id = row.id
            this.RowData.nickname = row.nickname
            this.RowData.username = row.username
            this.RowData.user_type_id = row.user_type_id
            this.RowData.email = row.email
            this.RowData.phone = row.phone
            let roles = []
            for(let i=0;row.roles.length>i;i++){
              roles.push(row.roles[i].id)
            }
            this.RowData.roles = roles
          },

          NewRow: function() {
            this.NewOrUpdate = true
            this.title = '添加'
            this.RowData.id = ""
            this.RowData.nickname = ""
            this.RowData.username = ""
            this.RowData.user_type_id = "1"
            this.RowData.email = ""
            this.RowData.phone = ""
            this.RowData.password = ""
            this.RowData.repassword = ""
            this.RowData.roles = []
          },

          ConfirmDelRows: function(row) {
            this.$confirm('确认删除用户：'+row.nickname + "<" + row.username+'>', '提示', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'warning'
            }).then(() => {
              this.action_loading = true
              this.$store.dispatch("common/Delete",{"url":this.url,"data": {rows: [row.id]}}).then((response) => {
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
          role_select: role_select,
        }
    }
</script>

<style scoped>

</style>
