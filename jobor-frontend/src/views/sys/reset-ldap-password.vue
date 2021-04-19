<template>
    <div style="width: 400px;margin: auto;vertical-align:middle;top: 35%;position: absolute;left: 50%;transform: translate(-50%,-50%);">

      <el-card class="box-card" style="width: 100%;clear: both" size="mini">
        <div slot="header" class="clearfix" style="">
          <span style="font-size: 14px">重置密码</span>
        </div>
        <el-form label-position="right" label-width="80px" size="mini" :rules="restRules" :model="form">
          <el-form-item label="用户名" prop="username">
            <el-input v-model="form.username" placeholder="用户名" style="width: auto"></el-input>@oyohotels.cn
          </el-form-item>
          <el-form-item label="重置方式" prop="reset_type" :rules="[{required:true,message:'.', trigger: 'blur'}]">
            <el-select v-model="form.reset_type" placeholder="重置方式">
              <el-option label="Email" value="email"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="账户状态" :hidden="rsetShow">
            <el-tag :type="account_type">{{form.account_disabled}}</el-tag>
          </el-form-item>
          <el-form-item label="新密码" :hidden="rsetShow" prop="password" :rules="[
          {required:true,message:'请输入新密码.', trigger: 'blur'},
           { pattern: /^(?![a-zA-Z]+$)(?![A-Z0-9]+$)(?![A-Z\W_]+$)(?![a-z0-9]+$)(?![a-z\W_]+$)(?![0-9\W_]+$)[a-zA-Z0-9\W_]{8,30}$/,
           message: '大、小写字母，数字和特殊符号，长度为 8 - 30位' }
          ]">
            <el-input v-model="form.password" type="password"></el-input>
          </el-form-item>
          <el-form-item label="确认密码" :hidden="rsetShow" prop="repassword" :rules="[{required:true,message:'请输入确认密码.', trigger: 'blur'}]">
            <el-input v-model="form.repassword" type="password"></el-input>
          </el-form-item>
          <el-form-item label="Token" :hidden="rsetShow" prop="token" :rules="[{required:true,message:'请输入token(请查询邮箱信息).', trigger: 'blur'}]">
            <el-input v-model="form.token"></el-input>
          </el-form-item>
          <el-form-item >
            <el-button type="primary" @click="SubmitUserInfo" :hidden="submitHidden">提交</el-button>
            <el-button type="warning" :hidden="rsetShow" @click="ResetPassword" style="margin: 0">重置</el-button>
            <el-button @click="goBack" :hidden="rsetShow" style="margin: 0;" type="primary" plain>返回</el-button>
          </el-form-item>
        </el-form>
      </el-card>
      <div style="font-size: 13px;color: red;font-weight: bold">* 密码必须包含字母大小写、数字加字符，最少8位，不能包含用户名</div>
      <div style="font-size: 13px;color: red;font-weight: bold">* 如果账号已禁用,请联系IT管理员</div>
      <div style="font-size: 13px;color: red;font-weight: bold">* Token 通过邮件获取</div>
    </div>
</template>

<script>
  export default {
    name: 'reset-ldap-password',
    data () {
      return {
        restRules: {
          username: [{required:true,message:'请输入AD用户名.', trigger: 'blur'}],
        },
        submitHidden: false,
        rsetShow: true,
        account_type: 'success',
        form: {
          username: '',
          reset_type: 'email',
          account_disabled: '',
          password: '',
          repassword: '',
          token: '',
        },
        url: this.GLOBAL.service_tag+'/api/password-reset-json/'
      }
    },
    methods: {
      SubmitUserInfo: function() {
        if(!this.form.username){this.$message.warning('请输入必填项.'); return}
        this.axios({
          method: 'GET',
          url: this.url,
          params: {'username': this.form.username},
        }).then((response)=>{
          // console.log(response.data);
          if(response.data.status){
            // console.log(response.data.data);
            const user_info = response.data.data
            if(user_info.account_disabled){this.form.account_disabled='已禁用';this.account_type='danger'}else{this.form.account_disabled='未禁用';this.account_type='success'}
            this.rsetShow=false
            this.submitHidden=true
            this.$message.success({message: response.data.message})
          }else {this.$message.error({message: 'SubmitUserInfo Error: '+ response.data.error})}
        }).catch((response) => {
          this.$message.error({message: 'SubmitUserInfo Error: '+ response});
        });
      },
      ResetPassword: function() {
        if(!this.form.username || !this.form.password || !this.form.repassword || !this.form.token){this.$message.warning('请输入必填项.'); return}
        this.axios({
          method: 'POST',
          url: this.url,
          data: this.qs.stringify({'username': this.form.username, 'password': this.form.password, 'repassword': this.form.repassword, 'token': this.form.token}),
        }).then((response)=>{
          // console.log(response.data);
          if(response.data.status){
            // console.log(response.data.data);
            // const user_info = response.data.data
            this.$message.success({message: response.data.message})
          }else {this.$message.error({message: 'ResetPassword Error: '+ response.data.error})}
        }).catch((response) => {
          this.$message.error({message: 'ResetPassword Error: '+ response});
        });
      },
      // 返回上一页
      goBack: function() {
        window.history.back();
      },
    },
  }
</script>

<style scoped>

</style>
