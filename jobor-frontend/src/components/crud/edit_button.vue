<template>
    <el-button type="text" :size="size" style="color: #e6a23c" @click="handleClick">{{title}}</el-button>
</template>

<script>
    export default {
      name: "delete",
      props:{
        size: {default: "mini"},
        title: {default: "编辑"},
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
