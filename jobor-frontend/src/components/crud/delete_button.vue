<template>
    <el-button type="text" :size="size" style="color: #f56c6c" @click="handleClick">{{title}}</el-button>
</template>

<script>
    export default {
      name: "delete",
      props:{
        size: {default: "mini"},
        title: {default: "删除"},
      },
      methods:{
        handleClick(evt) {
          this.$emit('click', evt);
        }
      },
    }
</script>

<style scoped>

</style>
