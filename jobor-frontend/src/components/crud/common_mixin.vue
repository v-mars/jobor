<script>
    export default {
      name: "common_mixin",
      props:{
        // url:{default: "", require:false},
      },
      data(){
        return{
          searchForm: {name: ""},
          extraParam: {},
          search: "",
          data_list: [],
          total: 0,
          page: 1,
          limit: 10,
          loading: false,
          url: "",

          action_loading: false,
          title: "新增",
          newOrUpdate: false,
          rowData: {},
        }
      },

      computed: {},

      watch: {},

      methods: {

        // 跳转详细
        clickTo: function(path){
          // console.log('row, column, cell, event', row)
          // this.$router.push()
        },

        getData: async function () {
          let data = {"pageSize": this.limit, "pageNumber": this.page}
          let param = Object.assign(data, this.searchForm, this.extraParam);
          this.loading = true
          try {
            const response = await this.$store.dispatch("common/Query", {url: this.url, data: param});
            this.data_list = response.data.data.list
            this.total = response.data.data.total
          } catch (e) {
            this.$message.error(e)
          } finally {
            this.loading = false
          }
        },

        createOrUpdate: function (method, again, id_url) {
          this.$refs.rowData.validate(valid =>{
            if(valid){
              // console.log("data_json:", data_json)
              let params = JSON.parse(JSON.stringify(this.rowData))
              if(method==='POST'){delete params.id}
              let api_url = this.url
              if(id_url){api_url=api_url+"/"+id_url}
              this.action_loading = true
              this.$store.dispatch("common/CreateUpdate",{url: api_url,"method": method, "data": params}).then((response) => {
                // console.log('createOrUpdate:',response.data);
                this.$message.success(response.data.message)
                if(again===false){this.newOrUpdate = false}
                this.getData()
                this.action_loading = false
              }).catch((response) => {
                this.action_loading = false
              })

            } else {this.$message.warning('请输入完整数据')}
          })
        },

        showEdit: function(row) {
          this.newOrUpdate = true
          this.title = '编辑'
          this.rowData.id = row.id
          // this.rowData.name = row.name
        },

        newRow: function() {
          this.newOrUpdate = true
          this.title = '添加'
          this.rowData.id = ""
          // this.rowData.name = ""
        },

        confirmDelRows: function(row, name,path_id) {
          this.$confirm('确认删除数据ID为['+row.id+']：'+name+' 吗？', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
          }).then(() => {
            this.action_loading = true
            let api_url = this.url
            if(path_id){api_url = api_url+"/"+row.id}
            this.$store.dispatch("common/Delete",{url: api_url,"data": {rows: [row.id]}}).then((response) => {
              // console.log('createOrUpdate:',response.data);
              this.$message.success(response.data.message)
              this.getData()
              this.action_loading = false
            }).catch((response) => {
              this.action_loading = false
            })
          }).catch(() => {this.$message({type: 'info', message: '已取消删除'})});
        },

        confirmDelBulkRows: function(rows) {
          console.log(rows);
          if (rows&&rows.length>0){
            let row_ids= []
            for (let i=0;rows.length>i;i++){row_ids.push(rows[i].id)}
            this.$confirm('确认批量删除数据ID为['+String(row_ids)+'] 吗？', '提示', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'warning'
            }).then(() => {
              this.action_loading = true
              let api_url = this.url
              this.$store.dispatch("common/Delete",{url: api_url,"data": {rows: row_ids}}).then((response) => {
                // console.log('createOrUpdate:',response.data);
                this.$message.success(response.data.message)
                this.getData()
                this.action_loading = false
              }).catch((response) => {
                this.action_loading = false
              })
            }).catch(() => {this.$message({type: 'info', message: '已取消删除'})});
          }

        },
      },

    }
</script>

