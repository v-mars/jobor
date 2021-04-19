<template>
    <span>

      <el-select v-if="!transfer" :value="SelectValue" filterable placeholder="请选择"
                 :clearable="clearable" style="width: 100%" :multiple="!single_selection" :collapse-tags="collapse_tags"
                 :allow-create="allowCreate"  @change="handleChange" reserve-keyword>
        <el-option v-html="getLabel(item)" :label="item.name" :value="item.id" v-for="(item,index) in data_list"
                   :key="index"></el-option>
      </el-select>

      <el-transfer v-else
        filterable
        :filter-method="filterMethod"
        filter-placeholder="请输入权限信息"
        :value="SelectValue"
        :data="permission_list"
        :titles="['权限列表', '已有权限']"
        @change="handleChange"
      >
          </el-transfer>
    </span>
</template>

<script>
    export default {
      name: "permission_select",
      props:{
        single_selection:{default: true},
        clearable:{default: true},
        collapse_tags:{default: true},
        allowCreate:{default: false},
        SelectValue:{required:false},
        transfer:{default:false},
      },
      data(){
        return{
          url: this.$store.state.urls.permission_url,
          data_list: [],
          total: 0,
          permission_list: [],
          search_list: [],
          filterMethod: function (query,item) {
            return item.pinyin.indexOf(query) > -1;
          },
        }
      },

      watch: {
      },

      methods: {
        getData: function(){
          let data = {pageSize: 0, pageNumber: 1}
          this.loading = true
          this.$store.dispatch("common/Query", {url: this.url,data: data}).then((response) => {
            this.data_list = response.data.data.list
            this.total = response.data.data.total
            let data_list = response.data.data.list
            this.permission_list = []
            this.search_list = []
            for (let i = 0; i < data_list.length; ++i) {
              this.search_list.push(`${data_list[i].name}:${data_list[i].path}`)
              this.permission_list.push({
                key: data_list[i].id,
                label: `${ data_list[i].name }`,
                disabled: false,
                pinyin: `${data_list[i].name}:${data_list[i].path}`,
              });
            }
            // console.log("this.permission_list:", this.permission_list)
          }).catch((response) => {})
        },
        handleChange(v){
          this.$emit('update:select-value', v)
          this.$emit('change', v)
        },

        getLabel: function (row) {
          return row.name+ ' <span style="font-size: 12px;color: #8f9298;float: right">'+"["+row.method+"] "+row.path+'</span>'
        },


      },
      mounted () {
        this.getData()
      },


    }
</script>

<style scoped>

</style>
