<template>
    <el-button type="text" :size="size" style="" @click="handleClick">{{title}}</el-button>
</template>

<script>
    export default {
      name: "delete",
      props:{
        size: {default: "mini"},
        title: {default: "详细"},
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
