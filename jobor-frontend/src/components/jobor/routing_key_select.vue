<template>
    <span>
      <el-select :value="SelectValue" filterable placeholder="请选择"
                 :clearable="clearable" style="" :multiple="!single_selection"
                 :allow-create="allowCreate"  @change="handleChange">
<!--        <el-option :label="item.name+'<'+item.code+'>'" :value="item.code" v-for="(item,index) in data_list"-->
<!--                   :key="index" v-if="use_code"></el-option>-->
        <el-option :label="item.routing_key" :value="item.routing_key" v-for="(item,index) in data_list"
                   :key="index"></el-option>
      </el-select>
    </span>
</template>

<script>
    export default {
      name: "routing_key_select",
      props:{
        single_selection:{default: true},
        use_code:{default: false},
        clearable:{default: true},
        allowCreate:{default: false},
        SelectValue:{required:false},
      },
      data(){
        return{
          url: this.$store.state.urls.jobor_routing_key_url,
          data_list: [],

          total: 0,
        }
      },

      watch: {

      },

      methods: {
        getData: function(){
          let data = {}
          this.$store.dispatch("common/Query", {url: this.url,data: data}).then((response) => {
            this.data_list = response.data.data.list
            this.total = response.data.data.total
          }).catch((response) => {})
        },
        handleChange(v){
          this.$emit('update:select-value', v)
        },

      },
      mounted () {
        this.getData()
      },


    }
</script>

<style scoped>

</style>
