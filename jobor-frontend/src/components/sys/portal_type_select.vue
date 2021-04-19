<template>
    <span>
<!--      <el-select :value="selectValue" filterable placeholder="请选择"-->
<!--                 :clearable="clearable" style="" :multiple="!single_selection" :collapse-tags="collapse_tags"-->
<!--                 :allow-create="allowCreate"  @change="handleChange">-->
<!--        <el-option :label="item.type" :value="item.type" v-for="(item,index) in data_list"-->
<!--                   :key="index"></el-option>-->
<!--      </el-select>-->

      <el-select v-model="select_id" filterable placeholder="类型，单选" v-if="single_selection" :clearable="clearable"
                 style="" :allow-create="allowCreate" :collapse-tags="collapse_tags">
        <el-option :label="item.type" :value="item.type" v-for="(item,index) in data_list" :key="index"></el-option>
      </el-select>
      <el-select v-model="select_list" filterable placeholder="类型，多选" v-else multiple :clearable="clearable"
                 style="" :allow-create="allowCreate" :collapse-tags="collapse_tags">
        <el-option :label="item.type" :value="item.type" v-for="(item,index) in data_list" :key="index"></el-option>
      </el-select>
    </span>
</template>

<script>
    export default {
      name: "portal_type_select",
      props:{
        single_selection:{default: true},
        clearable:{default: true},
        allowCreate:{default: false},
        collapse_tags:{default: true},
        selectId:{type: String, required:false},
        selectList:{type: Array, required:false},
        selectValue:{required:false},
      },
      data(){
        return{
          url: this.$store.state.urls.portal_url,
          data_list: [],
          select_id: this.selectId,
          select_list: this.selectList,

          total: 0,
          page: 1,
          limit: 10,
        }
      },

      watch: {
        select_id(val) {
          this.$emit('update:selectId', val)
        },
        selectId(val){
          this.select_id = val
        },

        select_list(val) {
          this.$emit('update:selectList', val)
        },
        selectList(val){
          this.select_list = val
        },
      },

      methods: {
        getData: function(){
          let data = {}
          this.loading = true
          this.$store.dispatch("common/Query", {url: this.url+'-type',data: data}).then((response) => {
            // console.log('aaa:',response.data);
            this.data_list = response.data.data.list
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
