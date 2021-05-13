<template>
  <el-date-picker
    v-model="valueCurr"
    type="datetimerange"
    :picker-options="pickerOptions"
    range-separator="至"
    start-placeholder="开始日期"
    end-placeholder="结束日期"
    :size="size"
    :clearable="false"
    @change="valueChange"
    align="left">
  </el-date-picker>
</template>

<script>
    export default {
      name: "DatetimeRange",
      props:{
        size: {default:"mini"},
        value: {Type: Array},
      },
      data(){
        return{
          pickerOptions: {
            shortcuts: [{
              text: '最近一周',
              onClick(picker) {
                const end = new Date();
                const start = new Date();
                start.setTime(start.getTime() - 3600 * 1000 * 24 * 7);
                picker.$emit('pick', [start, end]);
              }
            }, {
              text: '最近一个月',
              onClick(picker) {
                const end = new Date();
                const start = new Date();
                start.setTime(start.getTime() - 3600 * 1000 * 24 * 30);
                picker.$emit('pick', [start, end]);
              }
            }, {
              text: '最近三个月',
              onClick(picker) {
                const end = new Date();
                const start = new Date();
                start.setTime(start.getTime() - 3600 * 1000 * 24 * 90);
                picker.$emit('pick', [start, end]);
              }
            }, {
              text: '最近6个月',
              onClick(picker) {
                const end = new Date();
                const start = new Date();
                start.setTime(start.getTime() - 3600 * 1000 * 24 * 180);
                picker.$emit('pick', [start, end]);
              }
            }, {
              text: '最近1年',
              onClick(picker) {
                const end = new Date();
                const start = new Date();
                start.setTime(start.getTime() - 3600 * 1000 * 24 * 365);
                picker.$emit('pick', [start, end]);
              }
            }]
          },
          valueCurr: [new Date((new Date()).getTime() - 3600 * 1000 * 24 * 30),new Date()],
        }
      },
      methods: {
        valueChange: function (v) {
          this.$emit("update:value",this.valueCurr)
          this.$emit('change', v)
        }
      },
      mounted () {

        this.valueChange()
      },

    }
</script>

<style scoped>

</style>
