<template>
    <el-button type="text" :size="size" style="color: #909399" @click="handleClick">{{title}}</el-button>
</template>

<script>
    export default {
      name: "delete",
      props:{
        size: {default: "mini"},
        title: {default: "更多"},
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
