<template>
    <span>

      <el-select :value="selectValue" filterable placeholder="请选择"
                 :clearable="clearable" style="" :multiple="!single_selection" :collapse-tags="collapse_tags"
                 :allow-create="allowCreate"  @change="handleChange">
        <el-option :label="item.name" :value="item.id" v-for="(item,index) in data_list"
                   :key="index"></el-option>
      </el-select>
    </span>
</template>

<script>
    export default {
      name: "role_select",
      props:{
        single_selection:{default: true},
        clearable:{default: true},
        collapse_tags:{default: true},
        allowCreate:{default: false},
        selectValue:{required:false},
      },
      data(){
        return{
          url: this.$store.state.urls.role_url,
          data_list: [],
          total: 0,
        }
      },

      watch: {
      },

      methods: {
        getData: function(){
          let data = {}
          this.loading = true
          this.$store.dispatch("common/Query", {url: this.url+"s",data: data}).then((response) => {
            this.data_list = response.data.data.list
            this.total = response.data.data.total
          }).catch((response) => {})
        },
        handleChange(v){
          this.$emit('update:selectValue', v)
        },
      },
      mounted () {
        this.getData()
      },


    }
</script>

<style scoped>

</style>
