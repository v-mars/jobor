<template>
    <span>

      <el-select :value="selectValue" filterable placeholder="请选择" :collapse-tags="collapse_tags"
                 :clearable="clearable" style="" :multiple="!single_selection"
                 :allow-create="allowCreate"  @change="handleChange">
        <el-option :label="item.nickname+'<'+item.username+'>'" :value="item.id" v-for="(item,index) in data_list"
                   :key="index"></el-option>
      </el-select>
    </span>
</template>

<script>
    export default {
      name: "user_select",
      props:{
        single_selection:{default: true},
        clearable:{default: true},
        allowCreate:{default: false},
        collapse_tags:{default: true},
        use_code:{default: false},
        selectValue:{required:false},
      },
      data(){
        return{
          url: this.$store.state.urls.user_url,
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
