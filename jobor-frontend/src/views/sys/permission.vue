<template>
  <div>
    <div style="">
      <!-- 搜索 -->
      <el-form :inline="true" :model="searchForm" size="small" class="demo-form-inline" @submit.native.prevent>
        <el-form-item label="">
          <el-input v-model="searchForm.name" placeholder="名称" @keyup.enter.native="getData"></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="searchForm.path" placeholder="路径" @keyup.enter.native="getData"></el-input>
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="searchForm.method" placeholder="请求方法, 例：GET,POST,PUT,DELETE,PATCH" @keyup.enter.native="getData"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="getData">查询</el-button>
          <el-button type="primary" @click="newRow">新增</el-button>
          <el-button style="float: right" @click="getData">刷新</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div style="">
      <el-table border :data="data_list" size="small" v-loading="loading">
        <el-table-column type="selection" width="45" align="center"></el-table-column>
        <el-table-column label="ID" prop="id" width="60"></el-table-column>
        <el-table-column label="名称" prop="name" width=""></el-table-column>
        <el-table-column label="method" prop="method" width="100"></el-table-column>
        <el-table-column label="path" prop="path" show-overflow-tooltip></el-table-column>
        <!--            <el-table-column label="创建时间" prop="ctime" width="140"></el-table-column>-->
        <el-table-column label="更新时间" prop="mtime" width="140"></el-table-column>
        <el-table-column label="操作" align="center" width="130">
          <template slot-scope="scope">
            <edit_button @click="showEdit(scope.row)"></edit_button>
            <delete_button @click="confirmDelRows(scope.row, scope.row.name)"></delete_button>
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

    <el-dialog :title="title" :visible.sync="newOrUpdate" :close-on-click-modal="false">
      <el-form label-width="120px" :model="rowData" size="small" ref="rowData">
        <el-form-item label="名称" prop="name" :rules="[{required:true,message:'请输入名称', trigger: 'blur'},
        // {pattern:/^[a-z]+$/,message:'只能输入小写英文字母', trigger: 'blur'}
        ]">
          <el-input v-model="rowData.name"></el-input>
        </el-form-item>
        <el-form-item label="method" prop="method" :rules="[{required:true,message:'请输入', trigger: 'blur'}]">
          <el-input v-model="rowData.method"></el-input>
        </el-form-item>
        <el-form-item label="Path" prop="path" :rules="[{required:true,message:'请输入', trigger: 'blur'}]">
          <el-input v-model="rowData.path"></el-input>
        </el-form-item>
      </el-form>
      <span slot="footer" class="dialog-footer">
    <el-button @click="newOrUpdate = false">取 消</el-button>
    <el-button :plain="true" type="warning" @click="createOrUpdate('PUT', false, rowData.id)" v-if="rowData.id" :loading="action_loading">更新</el-button>
    <el-button :plain="true" type="warning" @click="createOrUpdate('POST', true)" v-if="!rowData.id" :loading="action_loading">保存并继续</el-button>
    <el-button :plain="true" type="primary" @click="createOrUpdate('POST',false)" v-if="!rowData.id" :loading="action_loading">保 存</el-button>
  </span>
    </el-dialog>


  </div>

</template>

<script>

  import pagination from '@/components/Pagination/pagination'
  import delete_button from '@/components/crud/delete_button'
  import edit_button from '@/components/crud/edit_button'
  import common_mixin from '@/components/crud/common_mixin'

  export default {
    name: "list",
    mixins: [common_mixin],
    data(){return{
      permissionVisible: false,
      rowData: {name: "", id: "",method:"", path:""},
      url: this.$store.state.urls.permission_url,
      searchForm: {name: "", method: "", path: ""},
    }},
    methods: {
      showEdit: function(row) {
        this.newOrUpdate = true
        this.title = '编辑'
        this.rowData.id = row.id
        this.rowData.name = row.name
        this.rowData.method = row.method
        this.rowData.path = row.path
        // this.rowData.description = row.description
      },

      newRow: function() {
        this.newOrUpdate = true
        this.title = '添加'
        this.rowData.id = ""
        this.rowData.name = ""
        this.rowData.description = ""
      },


    },
    mounted () {
      this.getData()
    },
    components: {
      pagination: pagination,
      delete_button,
      edit_button,
    },
  }
</script>

<style scoped>

</style>
